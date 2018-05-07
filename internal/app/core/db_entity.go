package core

import (
	"net/rpc"
	"log"
	"strings"
	"net"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
)

var listeningPort = "7000"

type Entity struct {
	ip 				string
	port 			string
	identifier 		string
	connector		rpc.Client
	isActive		bool
	slaves			[]Entity
}

type RPCRequest struct {
	Data []byte
}

func getSlavesIps() ([]string, error) {
	var ips []string
	var ipsJson, err = ioutil.ReadFile("databases/asd/connections.config")
	if err != nil {
		log.Fatal("Problem: ", err)
	}
	println(string(ipsJson))
	err = json.Unmarshal([]byte(ipsJson), &ips)

	return ips, err
}

func setMasterProp(ip string, port string) Entity{
	master = Entity{
		ip:         ip,
		port:       port,
		identifier: "master",
		isActive:   true,
		connector:  rpc.Client{},
		slaves:     nil,
	}
	setSlavesTo(&master)
	return master
}

func setSlavesTo(master *Entity) {
	var slavesAddresses, err = getSlavesIps()
	if err != nil {
		log.Fatal("Problem in decoding JSON Ips", err)
	}
	for _, slaveAddress := range slavesAddresses {
		slaveAddress := strings.Split(slaveAddress, ":")
		newSlave := Entity {
			ip:			slaveAddress[0],
			port:		slaveAddress[1],
			identifier:	"slave",
			isActive:	false,
			connector:	rpc.Client{},
			slaves:		nil,
		}
		master.slaves = append(master.slaves, newSlave)
	}
}

func getEntityIpAddress() (string, error) {
	var res = ""
	address, err := net.InterfaceAddrs()
	for _, a := range address {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				res = ipNet.IP.String()
			}
		}
	}
	return res, err
}

func InitEntity(entityType int) {
	switch entityType {
	case 0: // slave
		entity := new(Entity)
		rpc.Register(&entity)
		rpc.HandleHTTP()
		l, e := net.Listen("tcp", ":" + listeningPort)
		if e != nil {
			log.Fatal("listen error:", e)
		}
		go http.Serve(l, nil)
	case 1: // master
		var myIp, err = getEntityIpAddress()
		if err == nil {
			println("My IP: ", myIp)
		}
		master = setMasterProp(myIp, listeningPort)

		rpc.Register(&master)
		rpc.HandleHTTP()

		l, e := net.Listen("tcp", myIp + ":" + listeningPort)
		if e != nil {
			log.Fatal("listen error:", e)
		}

		go http.Serve(l, nil)

		for i, slave := range master.slaves {
			var rpcClient *rpc.Client
			attempts := 0
			for attempts != -1 {
				log.Printf("Try to connect (attempts %d) to %s:%s\n", attempts, slave.ip, slave.port)
				attempts++
				c := make(chan error, 1)
				go func() {
					rpcClient, err = rpc.DialHTTP("tcp", slave.ip + ":" + listeningPort)
					if err == nil {
						master.slaves[i].connector = *rpcClient
					}
					c <- err
				}()
				select {
				case err := <-c:
					if err != nil {
						log.Print("Dialing:", err)
						if attempts == 5 {
							attempts = -1
						} else {
							time.Sleep(time.Second)
						}
					} else {
						attempts = -1
					}
				case <-time.After(time.Second * 5):
					println("Timeout...")
				}
			}
		}

		for i, slave := range master.slaves {
			resp := RequestSlaveStatus(&slave)
			if resp == nil {
				resp := SendDeploy(&slave)
				if resp != nil {
					master.slaves[i].isActive = true
				}
			}
		}
	}
}
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
	Ip 				string
	Port 			string
	Identifier 		int
	Connector		rpc.Client
	IsActive		bool
	Slaves			[]Entity
}

type RequestedData struct {
	Payload			string
	File 			string
	Offset			int
	Id				int
	Bs 				[]byte
}

type Reply struct {
	Message string
	Data	[]byte
}

type RPCRequest struct {
	Data RequestedData
}

func getSlavesIps() ([]string, error) {
	var ips []string
	var ipsJson, err = ioutil.ReadFile("databases/asd/connections.config")
	if err != nil {
		log.Fatal("Problem: ", err)
	}
	err = json.Unmarshal([]byte(ipsJson), &ips)

	return ips, err
}

func setMasterProp(ip string, port string) Entity{
	master = Entity{
		Ip:         ip,
		Port:       port,
		Identifier: 0,
		IsActive:   true,
		Connector:  rpc.Client{},
		Slaves:     nil,
	}
	setSlavesTo(&master)
	return master
}

func setSlavesTo(master *Entity) {
	var slavesAddresses, err = getSlavesIps()
	if err != nil {
		log.Fatal("Problem in decoding JSON Ips", err)
	}
	for i, slaveAddress := range slavesAddresses {
		slaveAddress := strings.Split(slaveAddress, ":")
		newSlave := Entity {
			Ip:			slaveAddress[0],
			Port:		slaveAddress[1],
			Identifier:	i + 1,
			IsActive:	false,
			Connector:	rpc.Client{},
			Slaves:		nil,
		}
		master.Slaves = append(master.Slaves, newSlave)
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
		var entity Entity
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

		for i, slave := range master.Slaves {
			var rpcClient *rpc.Client
			attempts := 0
			for attempts != -1 {
				log.Printf("Try to connect (attempts %d) to %s:%s\n", attempts, slave.Ip, slave.Port)
				attempts++
				c := make(chan error, 1)
				go func() {
					rpcClient, err = rpc.DialHTTP("tcp", slave.Ip + ":" + listeningPort)
					if err == nil {
						master.Slaves[i].Connector = *rpcClient
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

		for i, slave := range master.Slaves {
			resp := RequestSlaveStatus(&slave)
			if resp == nil {
				resp := SendDeploy(&slave)
				if resp != nil {
					master.Slaves[i].IsActive = true
				}
			}
		}
	}
}
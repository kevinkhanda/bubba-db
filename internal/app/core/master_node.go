package core

import (
	"encoding/json"
	"net/rpc"
	"log"
	"time"
	"net"
	"net/http"
	"io/ioutil"
	"errors"
)

var master Entity

func requestSlaveStatus(entity *Entity) error {
	var reply string
	request := RPCRequest{"ping", nil}
	err = entity.connector.Call("Entity.SendStatus", &request, &reply)
	if err != nil {
		log.Fatal("Problems in requestSlaveStatus ", err)
	}
	if reply == "success" {
		println("Slave " + entity.ip + ":"+ entity.port + "is ready")
		for i, slave := range master.slaves {
			if slave.ip == entity.ip && slave.port == entity.port {
				master.slaves[i].isActive = true
				break
			}
		}
	} else {
		err = errors.New("incorrect response")
	}
	return err
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

func Test() {
	var myIp, err = getEntityIpAddress()
	if err == nil {
		println("My IP: ", myIp)
	}

	master = initMaster(myIp, "7000")
	initSlaves(&master)
	rpc.Register(&master)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", myIp+":7000")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	for i, slave := range master.slaves {
		// RPC call...
		var rpcClient *rpc.Client
		//var err error
		attempt := 1
		for attempt != -1 {
			log.Printf("Try to connect (attempt %d) to %s", attempt, slave.ip)
			println()
			attempt = attempt + 1
			c := make(chan error, 1)
			go func() {
				rpcClient, err = rpc.DialHTTP("tcp", slave.ip + ":7000")
				if err == nil {
					master.slaves[i].connector = *rpcClient
				}
				c <- err
			}()
			select {
			case err := <-c:
				if err != nil {
					log.Print("dialing:", err)
					time.Sleep(time.Second)
				} else {
					attempt = -1
				}
			case <-time.After(time.Second * 5):
				println("timeout...")
			}
		}
	}

	for _, slave := range master.slaves {
		requestSlaveStatus(&slave)
		var readyCount= 0
		for _, entity := range master.slaves {
			if entity.isActive {
				readyCount++
			}
		}
		if readyCount == len(master.slaves) {
		}
	}
}

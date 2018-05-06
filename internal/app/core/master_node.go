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
	var attempts = 0
	for attempts < 5 {
		err = errors.New("")
		request := RPCRequest{nil}
		err = entity.connector.Call("Entity.SendStatus", &request, &reply)
		if err != nil {
			log.Fatal("Problems in requestSlaveStatus ", err)
			err = errors.New("problems in requestSlaveStatus")
			attempts++
			continue
		}
		if reply == "success" {
			println("Slave " + entity.ip + ":" + entity.port + " is available")
			reply = ""
			request := RPCRequest{[]byte("db")}
			err = entity.connector.Call("Entity.Deploy", &request, &reply)
			if err != nil {
				log.Fatal("Problems during Deploying")
				err = errors.New("problems during Deploying")
				attempts++
				continue
			}
			if reply == "success" {
				println("Slave " + entity.ip + ":" + entity.port + "is initialized")
				for i, slave := range master.slaves {
					if slave.ip == entity.ip && slave.port == entity.port {
						master.slaves[i].isActive = true
						attempts = 5
						break
					}
				}
			}
		}
	}
	return err
}

func SendDropDatabase(entity *Entity) error {
	var reply string
	var attempts = 0
	for attempts < 5 {
		err = errors.New("")
		request := RPCRequest{nil}
		err = entity.connector.Call("Entity.SendStatus", &request, &reply)
		if err != nil {
			log.Fatal("Problems in requestSlaveStatus ", err)
			err = errors.New("problems in requestSlaveStatus")
			attempts++
			continue
		}
		if reply == "success" {
			attempts = 5
		}
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

func Init() {
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
		var rpcClient *rpc.Client
		attempt := 0
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
					log.Print("Dialing:", err)
					if attempt == 5 {
						attempt = -1
					} else {
						time.Sleep(time.Second)
					}
				} else {
					attempt = -1
				}
			case <-time.After(time.Second * 5):
				println("Timeout...")
			}
		}
	}

	for _, slave := range master.slaves {
		requestSlaveStatus(&slave)
	}
}

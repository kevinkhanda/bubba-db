package core

import (
	"encoding/json"
	"net/rpc"
	"log"
	"time"
)

var master Entity

type RPCRequest struct {
	Method 	string
	Data 	[]byte
}

func sendPing(slave *Entity)  {
	var reply string
	request := RPCRequest{"ping", nil}
	slave.connector.Call("Ping", &request, &reply)
	println("Answer from ", slave.ip, ":", slave.port, " ", reply)
}

func getSlavesIps() ([]string, error) {
	var ips []string
	var ipsJson = string("[\"10.91.40.48:5000\"]")
	err := json.Unmarshal([]byte(ipsJson), &ips)
	return ips, err
}

func Test() {
	var myIp, err = getEntityIpAddress()
	if err == nil {
		println("My IP: ", myIp)
	}

	master = initMaster(myIp, "7000")
	initSlaves(&master)

	for _, slave := range master.slaves {
		// RPC call...
		var rpcClient *rpc.Client
		//var err error
		attempt := 1
		for attempt != -1 {
			log.Printf("Try to connect (attempt %d) to %s", attempt, slave.ip)
			attempt = attempt + 1
			c := make(chan error, 1)
			go func() {
				rpcClient, err = rpc.DialHTTP("tcp", slave.ip + ":7000")
				if err == nil {
					println("GG WP")
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
		sendPing(&slave)
	}
}

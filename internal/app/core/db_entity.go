package core

import (
	"net/rpc"
	"log"
	"strings"
	"net"
)

type Entity struct {
	ip 				string
	port 			string
	identifier 		int
	connector		rpc.Client
	isActive		bool
	slaves			[]Entity
}


func initMaster(ip string, port string) Entity{
	master = Entity{
		ip:         ip,
		port:       port,
		identifier: 0,
		isActive:   true,
		connector:  rpc.Client{},
		slaves:     nil,
	}
	return master
}

func initSlaves(master *Entity){
	var slavesAddresses, err = getSlavesIps()
	if err != nil {
		log.Fatal("Problem in decoding JSON Ips", err)
	}
	initialIdentifier := 1
	for _, slaveAddress := range slavesAddresses {
		slaveAddress := strings.Split(slaveAddress, ":")
		newSlave := Entity {
			ip:			slaveAddress[0],
			port:		slaveAddress[1],
			identifier:	initialIdentifier,
			isActive:	false,
			connector:	rpc.Client{},
			slaves:		nil,
		}
		initialIdentifier++
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
package main

import (
	"fmt"
	"flag"
	// "net/http"
	// "gobackup/tempCodeRunnerFile"
)

func main() {
	var currentNodeStatus string
	flag.StringVar(&currentNodeStatus, "status", "master", "Slave port node number")

	var currentMasterNode int
	flag.IntVar(&currentMasterNode, "master-port", DEFAULT_MASTER_PORT, "Master node port number")

	var currentNodePort int
	flag.IntVar(&currentNodePort, "port", slavePort, "help message for flagname")
	// currentNodePort = flag.String("port", 8080, "current node port")

	flag.Parse()

	var node NodeConfig

	if currentNodeStatus == "master" {
		// currentNodePort = 48004
		node = NodeConfig{name: "master", port: currentNodePort}
		setNode(&masterNode, "master", currentNodePort)

		startMaster()
	} else {
		// currentNodePort = 8081

		var masterNodePort *int = &currentMasterNode

		node = NodeConfig{name: "slave", port: currentNodePort}
		setNode(&slaveNode, "slave", currentNodePort)

		var hostNode string = getMasterNode(masterNodePort)

		fmt.Println(hostNode)
		// check for and connect with master

	}

	fmt.Println(node.GetNodeName())
	fmt.Println(node.GetNodePort())

	// fmt.Println(slaveName)
	// fmt.Println(masterName)
}




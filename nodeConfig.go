package main

import (
	// "fmt"
	"strconv"
)

type NodeConfig struct {
	name string
	port int
}

// Config for TCP Connections
const (
    TCP_CONN_HOST = "localhost"
    TCP_CONN_TYPE = "tcp"
	DEFAULT_MASTER_PORT = 8080
)

func (n NodeConfig) GetNodeName() string {
	return n.name
}

func (n NodeConfig) GetNodePort() int {
	return n.port
}

func (n NodeConfig) GetNodePortStr() string {
	return strconv.Itoa(n.port)
}

func (n NodeConfig) SetNodeName(name string) {
	n.name = name
}

func (n NodeConfig) SetNodePort(port int) {
	n.port = port
}

func setNode(inputNode *NodeConfig, name string, port int) {
	inputNode.name = name
	inputNode.port = port
}
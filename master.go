package main

import (
	"fmt"
	"net"
	"os"
	// "strconv"
)

var masterName string = "master"

var masterNode NodeConfig

func startMaster() {
	fmt.Println("Starting master node")

	l, err := net.Listen(TCP_CONN_TYPE, TCP_CONN_HOST+":"+masterNode.GetNodePortStr())

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()

	fmt.Println("Listening on " + TCP_CONN_TYPE + ":" + masterNode.GetNodePortStr())

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go handleMasterRequest(conn)
	}
	// fmt.Println("Master node started")
}

func handleMasterRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
	  fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}


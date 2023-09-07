package main

import (
	"fmt"
	"net/http"
)

// masterPort := 8080
var masterPort int = 8080
var slavePort int = 8081

var HOSTNAMES[16] string = [16]string{
	"vlab00.cs.usfca.edu",
	"vlab01.cs.usfca.edu",
	"vlab02.cs.usfca.edu",
	"vlab03.cs.usfca.edu",
	"vlab04.cs.usfca.edu",
	"vlab05.cs.usfca.edu",
	"vlab06.cs.usfca.edu",
	"vlab07.cs.usfca.edu",
	"vlab08.cs.usfca.edu",
	"vlab09.cs.usfca.edu",
	"vlab10.cs.usfca.edu",
	"vlab11.cs.usfca.edu",
	"vlab12.cs.usfca.edu",
	"vlab13.cs.usfca.edu",
	"vlab14.cs.usfca.edu",
	"vlab15.cs.usfca.edu",
}

func checkMaster(host string, port *int) chan bool {
	ch := make(chan bool, 1)
	ch <- false

	if port == nil {
		port = &masterPort
	}

	fmt.Println("Checking master on port " + string(*port))

	go func() {
		_, err := http.Get("http://" + host + ":" + string(*port))
		if err != nil {
			<-ch // empty the channel
			ch <- true
			close(ch)
			return
		}
	}()
	return ch
}

func getMasterNode(port *int) string {
	// ping all hosts through goroutines
	// on port 8080 to see if they are master

	select {
		case <-checkMaster(HOSTNAMES[0], port):
			return HOSTNAMES[0]
		case <-checkMaster(HOSTNAMES[1], port):
			return HOSTNAMES[1]
		case <-checkMaster(HOSTNAMES[2], port):
			return HOSTNAMES[2]
		case <-checkMaster(HOSTNAMES[3], port):
			return HOSTNAMES[3]
		case <-checkMaster(HOSTNAMES[4], port):
			return HOSTNAMES[4]
		case <-checkMaster(HOSTNAMES[5], port):
			return HOSTNAMES[5]
		case <-checkMaster(HOSTNAMES[6], port):
			return HOSTNAMES[6]
		case <-checkMaster(HOSTNAMES[7], port):
			return HOSTNAMES[7]
		case <-checkMaster(HOSTNAMES[8], port):
			return HOSTNAMES[8]
		case <-checkMaster(HOSTNAMES[9], port):
			return HOSTNAMES[9]
		case <-checkMaster(HOSTNAMES[10], port):
			return HOSTNAMES[10]
		case <-checkMaster(HOSTNAMES[11], port):
			return HOSTNAMES[11]
		case <-checkMaster(HOSTNAMES[12], port):
			return HOSTNAMES[12]
		case <-checkMaster(HOSTNAMES[13], port):
			return HOSTNAMES[13]
		case <-checkMaster(HOSTNAMES[14], port):
			return HOSTNAMES[14]
		case <-checkMaster(HOSTNAMES[15], port):
			return HOSTNAMES[15]
		case <- checkMaster("localhost", port):
			return "localhost"
		default:
			fmt.Println("No master found")
			err := "No master found"
			panic(err)
	}

	return "vlab00.cs.usfca.edu"
	// return NodeConfig{name: "master", port: 8080}
}
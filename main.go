// https://en.wikipedia.org/wiki/Transmission_Control_Protocol
// https://www.ibm.com/docs/en/zos/2.1.0?topic=SSLTBW_2.1.0/com.ibm.zos.v2r1.halu101/constatus.htm

package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup

var channel = make(chan *TCP)

type TCP struct {
	src int
	dst int
	seq int
	ack int

	SYN bool
	RST bool
	ACK bool
	FIN bool

	data string
}

func newPackage(seq int, ack int) *TCP {
	p := TCP{}
	p.seq = seq
	p.ack = ack
	return &p
}

func client(ip int) {
	defer wait.Done()

	// Sends a Synchronize request
	seq := randomNumber()
	channel <- newPackage(seq, 0)

	// Awaits Acknowledgement

	select {
	case p := <-channel:
		// Recieved Acknowledgement
		if seq+1 != p.ack {
			fmt.Println("Connection failed!")
			return
		}

		channel <- newPackage(p.ack, p.seq+1)
	case <-time.After(3 * time.Second):
		fmt.Println("Connection Timeout")
	}

}

func server(ip int) {

	for {

		// Passive listening
		p := <-channel

		// Recieved a Synchronize request
		seq := randomNumber()
		channel <- newPackage(seq, p.seq+1)

		p = <-channel

		if seq+1 != p.ack {
			fmt.Println("Connection failed!")
			return
		}

	}
}

func listen() {

}

func connect() {

}

func close() {

}

func send() {

}

func main() {
	go client(1)
	go server(2)
	wait.Add(1)

	wait.Wait()
	fmt.Println("Program ran successfully")
}

package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup
var channel = make(chan *packet)

func client() {
	defer wait.Done()

	// Generate random sequence number
	seq := randomNumber()
	fmt.Println("Client sequence:", seq)

	// Send SYN packet
	channel <- SynPacket(seq)

	// Wait for acknowledgement
	select {
	case p := <-channel:
		// Recieved acknowledgement

		// Check if acknowledgement matches sequence number
		if seq+1 != p.ack {
			fmt.Println("Connection failed!")
			return
		}

		// Send ACK packet
		channel <- AckDataPacket(p.seq+1, "Hello World!")
	case <-time.After(3 * time.Second):
		// Timeout

		fmt.Println("Connection Timeout")
		return
	}
}

func server() {
	for {
		// Passive listening
		p := <-channel

		// Recieved a Synchronize request
		seq := randomNumber()
		fmt.Println("Server sequence:", seq)

		channel <- SynAckPacket(seq, p.seq+1)

		p = <-channel

		if seq+1 != p.ack {
			fmt.Println("Connection failed!")
			return
		}

		fmt.Println("Connection succes!")
	}
}

func main() {
	go client()
	go server()
	wait.Add(2)

	wait.Wait()
}

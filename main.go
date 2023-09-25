package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wait sync.WaitGroup
var channel = make(chan *packet)

func client() {
	defer wait.Done()

	// Generate random sequence number
	seq := randomNumber()
	fmt.Println("Client: seq=" + strconv.Itoa(seq))

	// Send SYN packet
	channel <- SynPacket(seq)

	// Wait for acknowledgement
	select {
	case p := <-channel:
		// Recieved acknowledgement
		fmt.Println("Client got: SYN+ACK seq=" + strconv.Itoa(p.seq) + " ack=" + strconv.Itoa(p.ack))

		// Check if acknowledgement matches sequence number
		if seq+1 != p.ack {
			fmt.Println("Client: Connection failed!")
			return
		}

		// Send SYN+ACK packet
		channel <- SynAckPacket(p.ack, p.seq+1)
	case <-time.After(3 * time.Second):
		// Timeout
		fmt.Println("Client: Connection Timeout")
		return
	}
}

func server() {
	defer wait.Done()

	// Generate random sequence number
	seq := randomNumber()
	fmt.Println("Server: seq=" + strconv.Itoa(seq))

	// Listens for SYN packet
	p := <-channel

	// Recieved SYN packet
	fmt.Println("Server got: SYN seq=" + strconv.Itoa(p.seq))

	// Send SYN+ACK packet
	channel <- SynAckPacket(seq, p.seq+1)

	// Wait for acknowledgement
	select {
	case p = <-channel:
		// Recieved acknowledgement
		fmt.Println("Server got: SYN+ACK seq=" + strconv.Itoa(p.seq) + " ack=" + strconv.Itoa(p.ack))

		// Check if acknowledgement matches sequence number
		if seq+1 != p.ack {
			fmt.Println("Server: Connection failed!")
			return
		}
	case <-time.After(3 * time.Second):
		// Timeout
		fmt.Println("Server: Connection Timeout")
		return
	}

	// 3-Way Handshake successful!
	fmt.Println("Connection success!")
}

func main() {
	go client()
	go server()
	wait.Add(2)

	wait.Wait()
}

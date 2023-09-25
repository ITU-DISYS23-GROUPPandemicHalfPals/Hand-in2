package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wait sync.WaitGroup
var channel = make(chan *packet)

func main() {
	go client()
	go server()
	wait.Add(2)

	wait.Wait()
}

func client() {
	defer wait.Done()

	// Generate random sequence number for the client
	seq := rand.Intn(1000)
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

	// Send data packet
	channel <- DataPacket(seq+2, "Hello, World!")
}

func server() {
	defer wait.Done()

	// Generate random sequence number for the server
	seq := rand.Intn(1000)
	fmt.Println("Server: seq=" + strconv.Itoa(seq))

	// Listen for SYN packet
	p := <-channel

	// Recieved SYN packet
	fmt.Println()
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
	fmt.Println("3-Way Handshake successful!")
	fmt.Println()

	// Listen for data packet
	p = <-channel

	// Recieved data packet
	fmt.Println("Server got: seq=" + strconv.Itoa(p.seq) + " data=" + p.data)
}

// Packet datastructure
type packet struct {
	seq int
	ack int

	data string
}

// To create SYN packet
func SynPacket(seq int) *packet {
	p := packet{}
	p.seq = seq
	return &p
}

// Function to create SYN+ACK packet
func SynAckPacket(seq int, ack int) *packet {
	p := packet{}
	p.seq = seq
	p.ack = ack
	return &p
}

// Function to create data packet
func DataPacket(seq int, data string) *packet {
	p := packet{}
	p.seq = seq
	p.data = data
	return &p
}

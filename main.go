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

	// Sends a Synchronize request
	seq := randomNumber()
	fmt.Println("Client sequence:", seq)

	channel <- newSynPacket(seq)

	// Awaits Acknowledgement
	select {
	case p := <-channel:
		// Recieved Acknowledgement
		if seq+1 != p.ack {
			fmt.Println("Connection failed!")
			return
		}

		channel <- newAckDataPacket(p.seq+1, "Hello World!")
	case <-time.After(3 * time.Second):
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

		channel <- newSynAckPacket(seq, p.seq+1)

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

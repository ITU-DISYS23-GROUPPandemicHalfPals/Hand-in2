package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup
var channel = make(chan *packet)

func client() {

	// Sends a Synchronize request
	seq := randomNumber()
	fmt.Println("Client sequence:", seq)

	channel <- newSynPacket(seq)

	// Awaits Acknowledgement
	select {
	case p := <-channel:
		// Recieved Acknowledgement
		if seq+1 != p.ack {
			fmt.Println("Client: Connection failed!")
			return
		}

		channel <- newAckDataPacket(p.seq+1, "Hello World!")
	case <-time.After(3 * time.Second):
		fmt.Println("Client: Connection Timeout")
		return
	}
}

func server() {
	defer wait.Done()

	for {

		// Passive listening
		p := <-channel

		// Recieved a Synchronize request
		seq := randomNumber()
		fmt.Println("Server sequence:", seq)

		channel <- newSynAckPacket(seq, p.seq+1)

		// Awaits final Synchronize acknowledgement acknowledgement
		select {
		case p = <-channel:
			if seq+1 != p.ack {
				fmt.Println("Server: Connection failed!")
				continue
			}
		case <-time.After(3 * time.Second):
			fmt.Println("Server: Connection Timeout")
			continue

		}

		fmt.Println("Connection succes!")
		return

	}

}

func main() {
	go client()
	go server()
	wait.Add(1)

	wait.Wait()
}

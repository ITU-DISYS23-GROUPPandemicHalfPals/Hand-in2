package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup
var channel = make(chan *packet)

func client() {
	defer wait.Done()

	seq := randomNumber()
	fmt.Println("Client sequence:", seq)

	channel <- newSynPacket(seq)

	p := <-channel

	if seq+1 != p.ack {
		fmt.Println("Connection failed!")
		return
	}

	channel <- newAckDataPacket(p.seq+1, "Hello World!")

	fmt.Println("Connection succes!")
}

func server() {
	defer wait.Done()

	seq := randomNumber()
	fmt.Println("Server sequence:", seq)

	p := <-channel
	channel <- newSynAckPacket(seq, p.seq+1)

	p = <-channel

	if seq+1 != p.ack {
		fmt.Println("Connection failed!")
		return
	}

	fmt.Println("Server recieved:", p.data)
}

func main() {
	go client()
	go server()
	wait.Add(2)

	wait.Wait()
}

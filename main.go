package main

import (
	"fmt"
	"sync"
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

	seq := randomNumber()
	channel <- newPackage(seq, 0)

	p := <-channel

	if seq+1 != p.ack {
		fmt.Println("Connection failed!")
		return
	}

	channel <- newPackage(p.ack, p.seq+1)

	for {

	}
}

func server(ip int) {
	defer wait.Done()

	seq := randomNumber()
	p := <-channel
	channel <- newPackage(seq, p.seq+1)

	p = <-channel

	if seq+1 != p.ack {
		fmt.Println("Connection failed!")
		return
	}

	for {

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
}

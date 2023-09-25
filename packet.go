package main

type packet struct {
	SYN bool
	ACK bool

	seq int
	ack int

	data string
}

func newSynPacket(seq int) *packet {
	p := packet{}
	p.SYN = true
	p.ACK = false
	p.seq = seq
	return &p
}

func newSynAckPacket(seq int, ack int) *packet {
	p := packet{}
	p.SYN = true
	p.ACK = false
	p.seq = seq
	p.ack = ack
	return &p
}

func newAckDataPacket(ack int, data string) *packet {
	p := packet{}
	p.SYN = false
	p.ACK = true
	p.ack = ack
	p.data = data
	return &p
}

package main

type packet struct {
	SYN bool
	ACK bool

	seq int
	ack int

	data string
}

func SynPacket(seq int) *packet {
	p := packet{}
	p.SYN = true
	p.ACK = false
	p.seq = seq
	return &p
}

func SynAckPacket(seq int, ack int) *packet {
	p := packet{}
	p.SYN = true
	p.ACK = false
	p.seq = seq
	p.ack = ack
	return &p
}

func AckDataPacket(ack int, data string) *packet {
	p := packet{}
	p.SYN = false
	p.ACK = true
	p.ack = ack
	p.data = data
	return &p
}

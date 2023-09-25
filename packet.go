package main

type packet struct {
	seq int
	ack int

	data string
}

func SynPacket(seq int) *packet {
	p := packet{}
	p.seq = seq
	return &p
}

func SynAckPacket(seq int, ack int) *packet {
	p := packet{}
	p.seq = seq
	p.ack = ack
	return &p
}

func DataPacket(ack int, data string) *packet {
	p := packet{}
	p.ack = ack
	p.data = data
	return &p
}

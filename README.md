(a) In our implementation we have created a datastructure that looks like the following. It is done this way such that we can simulate the 3-Way handshake with the sequence and acknowledgement numbers and transmit data as a string through the data varaible.

type packet struct {
	seq int
	ack int

	data string
}


(b) We have simulated the TCP protocol using two threads (One for server and for client) communicating through a channel. One thing that makes it unrealistic is that threads run on the same process, thus sharing CPU and data.


(c) In order to make sure that the ordering of the packets are correct, we would use the sequence number to make sure that the packets are ordered from lowest to highest. Such that it does not matter if packet #46 comes before #45.


(d) We would make a timeout timer, which waits a certain amount of time and if no messages is recieved will assume the other end stopped communication and stop as well.


(e) It assures that connection is realiable and ready to transfer data. It is made realiable by agreeing on the inital seqence numbers such that packets can be send properly.
(a) We created our own custom data-structure called Packet with variables:

     Integer syn
     Integer ack

     String Data

(b) We used Threads for both server and client simulations. One thing that makes it unrealistic is that threads run on the same process, thus sharing CPU and data. 

(c) After establishing connection, we would use the incrementally changing syn and ack values to keep track of of the correct order of messages.

(d) We would make a timeout timer, which waits a certain amount of time and if no messages is recieved will assume the other end stopped communications and stop aswell.

(e) It assures that connection is realiable and that no outside client is sending data pretending to be either part. Though not 100% secure, it heavily boosts these security factors.
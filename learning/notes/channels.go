package main

func main() {

	// creat unbuffered channels
	unbufferedChannel := make(chan int)

	// create buffered channels : bufferedChannel := make(chan int, bufferNumber)
	bufferedChannel := make(chan int, 3)

	// channel types : bi-directional(default), send-only(chan<-), receive-only (<-chan)

	// closing a channel

	// channel control flow

}
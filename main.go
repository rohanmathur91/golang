package main

import (
	"fmt"
)

func main() {
	// channel
	messages := make(chan string)

	// msg0 := <-messages
	// fmt.Println(msg0)

	// goroutine
	go func() {
		messages <- "ping1"
		messages <- "ping2"
		messages <- "ping3"
		messages <- "ping4"
		messages <- "ping5"
	}()

	// msg1 := <-messages
	// fmt.Println(msg1)

	go func() {
		receiverMsg := <-messages
		fmt.Println("receiver: ", receiverMsg)
	}()

	msg := <-messages
	fmt.Println(msg)

	close(messages)
}

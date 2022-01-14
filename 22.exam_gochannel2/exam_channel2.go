package main

import "fmt"

func main() {
	gch := make(chan string, 1)
	sendChannel(gch)
	recvChannel(gch)
}

func sendChannel(gch chan<- string) {
	gch <- "magic"
}

func recvChannel(gch <-chan string) {
	data := <-gch
	fmt.Println(data)
}

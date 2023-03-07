package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for {
		time.Sleep(1 * time.Second)
		c <- "ping"
	}
}

func pong(c chan string) {
	for {
		time.Sleep(1 * time.Second)
		c <- "pong"
	}
}

func main() {
	pingChan := make(chan string)
	pongChan := make(chan string)

	go ping(pingChan)
	go pong(pongChan)

	for {
		select {
		case msg := <-pingChan:
			fmt.Println(msg)
		case msg := <-pongChan:
			fmt.Println(msg)
		}
	}
}

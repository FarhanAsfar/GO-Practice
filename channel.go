package main

import (
	"fmt"
)

func hell(done chan bool) {
	fmt.Println("hell goroutine")
	done <- true
}

func Channel() {
	done := make(chan bool)
	go hell(done)
	<-done
	fmt.Println("channel function")
}

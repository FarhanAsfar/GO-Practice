package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello goroutine")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		// time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alpha() {
	for i := 'a'; i <= 'f'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func Goroutine() {
	go hello()
	go numbers()
	go alpha()
	time.Sleep(2301 * time.Millisecond)
	fmt.Println("go-routine function")
}

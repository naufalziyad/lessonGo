package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("this is example with single goroutines ")
	go function()
	go func() {
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()
	time.Sleep(10 * time.Second)
}

func function() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()
}

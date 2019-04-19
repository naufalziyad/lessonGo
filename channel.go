package main

import ("fmt"
"runtime"
"time")

func main(){
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloTo = func(who string){
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	var now = time.Now()


	go sayHelloTo("hai")
	go sayHelloTo("Isyana")
	go sayHelloTo("Raisa")
	go sayHelloTo("Bambang")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)
	
	var message3 = <-messages
	fmt.Println(message3)

	var message4 = <-messages
	fmt.Println(message4)

	fmt.Println(now.String())
}
package main

import (
	"fmt"
	"time"
)

func f(from string){
	for i := 0; i < 3; i++{
		fmt.Println(from, ":", i)
	}
}

func main(){
	f("direct")

	// easy to create, just go!
	go f("goroutine")

	go func(msg string){
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
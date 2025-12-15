package goroutines

import (
	"fmt"
	"time"
)

//paralelleism is a subset of concurency
//go routiens make functions run concurently

func SayHello() {
	fmt.Println("hello !!")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("hello after 100milseconds")
}

func SayHai() {
	fmt.Println("hii !!")
}

func LearnGoroutines() {
	go SayHello()
	go SayHai()

	time.Sleep(1500 * time.Millisecond)

}

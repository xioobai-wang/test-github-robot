package main

import (
	"fmt"
	"time"
)

func main() {
	test10()
}

func test10() {
	stop := make(chan bool)
	go f1(stop)
	go f2(stop)
	time.Sleep(5 * time.Second)
	stop <- true
}

func f1(stop chan bool) {
	fmt.Println("start f1")
	fmt.Println("start f1")
	fmt.Println("start f1")
	fmt.Println("start f1")
	for {

	}
	fmt.Println("end f2")
}

func f2(stop chan bool) {
	fmt.Println("start f2")
	select {
	case <-stop:
		return
	default:
		fmt.Println("xxx")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("end f2")
}

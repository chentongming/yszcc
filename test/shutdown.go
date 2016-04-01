package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	go g(1)
	time.Sleep(time.Second * 2)
}

func g(i int) {
	defer func() {
		// err := recover()
		// fmt.Println(err)
		fmt.Println(11)
	}()
	panic("error")
}

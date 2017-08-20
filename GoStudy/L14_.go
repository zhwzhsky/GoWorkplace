package main

import (
	"fmt"
	// "time"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true //存入channel中
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

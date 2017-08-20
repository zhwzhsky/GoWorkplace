package main

import (
	"fmt"
)

const (
	Byte = 1 << (10 * iota)
	KB
	MB
	GB
	TB
)

func main() {
	fmt.Println(Byte)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
}

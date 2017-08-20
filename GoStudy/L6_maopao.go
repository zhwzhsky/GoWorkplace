package main

import (
	"fmt"
)

func main() {
	a := []int{3, 4, 2, 9, 76, 7, 8, 19, 0, 122, 73, 62, 6}
	fmt.Println(a)
	l := len(a)
	temp := a[0]
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if a[i] > a[j] {
				temp = a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}

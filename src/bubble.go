package main

import (
	"fmt"
)

func bubble() {
	a := [...]int{2, 5, 3, 7, 4, 8, 1}
	fmt.Println(a)
	length := len(a)
	for index := 0; index < length; index++ {
		for j := index; j < length; j++ {
			if a[index] > a[j] {
				a[index], a[j] = a[j], a[index]
			}
		}
	}
	fmt.Println(a)
}

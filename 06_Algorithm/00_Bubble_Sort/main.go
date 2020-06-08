package main

import (
	"fmt"
)

func BubbleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}

func main() {
	a := []int{5, 6, 1, 10, 89, 43}
	fmt.Println(BubbleSort(a))
}

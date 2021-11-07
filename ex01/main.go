package main

import (
	"ex01/deepcopy"
	"fmt"
	"log"
)

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a, err := deepcopy.DeepCopy(nums, 1, 4)
	if err != nil {
		log.Fatal(err)
	}
	for i, value := range a {
		a[i] += value
	}
	fmt.Println(nums)
	fmt.Println(a)
}

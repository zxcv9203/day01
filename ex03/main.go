package main

import (
	"ex03/hackathon"
	"fmt"
)

func main() {
	a := []int{2, 1, 3, 0, 5, 4, 0, 1}
	fmt.Println(hackathon.Match(a))
	fmt.Println(hackathon.Match([]int{0, 1, 2}))
	fmt.Println(hackathon.Match([]int{2, 0, 3, 1}))
	fmt.Println(hackathon.Match([]int{0, 2, 1}))
	fmt.Println(hackathon.Match([]int{0, 0, 0, 0, 0, 0}))
}

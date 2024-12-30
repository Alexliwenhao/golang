package main

import (
	"fmt"
)

func main() {
	var nilSlice []int
	fmt.Println("nilSlice length:", len(nilSlice))
	fmt.Println("nilSlice capacity:", len(nilSlice))

	s2 := []int{9, 8, 7, 6, 5}
	fmt.Println("s2 length: ", len(s2))
	fmt.Println("s2 capacity: ", cap(s2))
}

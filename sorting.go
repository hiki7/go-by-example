package main

import (
	"fmt"
	"slices"
)

// pretty much clear, just the use of methods
func main() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("strings:", strs)

	ints := []int{5, 1, 2}
	slices.Sort(ints)
	fmt.Println("ints:", ints)

	s := slices.IsSorted(strs)
	fmt.Println("is sorted:", s)
}

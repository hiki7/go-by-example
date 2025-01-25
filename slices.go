package main

import (
	"fmt"
	"slices"
)

func main() {

	//uninitialized slice equals to nil and has length 0
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	//to create an empty slice with non-zero length, we use builtin make
	//len is the number of values in slice, whereas capacity is the size of collection in memory
	s = make([]string, 3)
	fmt.Println("emp", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	//builtin append operation which returns a slice containing one or more new values
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[:2]
	fmt.Println("sl3:", l)

	//declaring and initializing the slice in one line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}

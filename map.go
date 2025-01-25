package main

import (
	"fmt"
	"maps"
)

func main() {
	//for creating an empty map, we use builtin make
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	//when we try to access the value of non-existent key, we get zero-value
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	//the builtin delete removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	//to remove all of the key/value pairs
	clear(m)
	fmt.Println("map:", m)

	//optional second return value when getting a value indicates if the key exists in the map
	_, prs := m["k2"]
	fmt.Println("prs", prs)

	//declaring and initializing map in single line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}

package main

import "fmt"

func main() {
	//generally, range iterates over elements in different data structures
	nums := []int{2, 3, 4}
	sum := 0

	//range in slices and arrays work the same, it provides both index and value for each entry
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//range in maps iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//it can also iterate over keys only
	for k := range kvs {
		fmt.Println("key: ", k)
	}

	//in strings range iterates over Unicode code points
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

package main

import "fmt"

func main() {
	//here we declare the length of the array, if it doesn't contain the values, by default they will be 0 in the case of int, ofc
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	//this is the way of declaring and initializing the array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//we can also hvae the compiler count the number of elements in the array, so we don't need to specify the length
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//we can use "index:", so that until this index everything will be filled with 0s
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}

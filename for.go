package main

import "fmt"

func main() {
	//the most basic type of for loop with single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	//classic initil;condition;after for loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	//another way of accomplishing the basic "do smth N times" iteration is range over an integer
	for i := range 3 {
		fmt.Println("range", i)
	}
	//for without the condition will infinitly run unless you "break" out of the loop or "return" from the enclosing function
	for {
		fmt.Println("infinite loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

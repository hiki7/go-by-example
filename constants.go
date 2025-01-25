package main

import (
	"fmt"
	"math"
)

//constants can be declared anywhere just like the "var" statement
//Go does not limit the precision of constant values, which means that the calculation with constants are more precise
//constants have no type until they are used somewhere in the code
//if some function (ex. math.Sin()) requires some specific data type, we need to explicitly convert the constant to necessary data type

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

package main

import "fmt"

type rect struct {
	width, height int
}

// method with a pointer receiver(*rect)
// it's useful for avoiding the copying large structs
// it allows to modify the original struct
func (r *rect) area() int {
	return r.width * r.height
}

// method with a values receiver(rect)
// useful if we don't need to modify the struct
// useful for small strcuts where copying is cheap
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println(r.area())
	fmt.Println(r.perim())

	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}

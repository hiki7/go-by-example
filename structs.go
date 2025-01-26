package main

import "fmt"

// in Go there are no classes, because Go tries to be simple and lightweight
// but we have struct
type person struct {
	name string
	age  int
}

// this can be considered as constructor, where you should define only the name
func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	//two ways of creating a struct
	//1)
	fmt.Println(person{"Bob", 20})
	//2)
	fmt.Println(person{name: "Alice", age: 30})

	//omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	//& provides pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	//the use of constructor. This way is considered to be idiomatic
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 40}

	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rext",
		true,
	}
	fmt.Println(dog)
}

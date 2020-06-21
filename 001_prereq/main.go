package main

import "fmt"

var y int

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James"`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Good morning, James". I'm gonna kill you.`)
}

func main() {
	// x := 7
	// fmt.Printf("%T \n", x)

	// fmt.Println(y)

	// xi := []int{1, 2, 3, 4, 5}
	// fmt.Println(xi)

	// m := map[string]int{
	// 	"John": 10,
	// 	"Job":  12,
	// }
	// fmt.Println(m)

	p1 := person{
		"John",
		"Doe",
	}
	// p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	// sa1.speak()
	// sa1.person.speak()

	saySomething(p1)
	saySomething(sa1)
}

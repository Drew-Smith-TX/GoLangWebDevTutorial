package main

import "fmt"

// package level scope
var y int

type hotdog int
type person struct {
	fname string
	lname string
}

func main() {
	x := 7 //block level scope
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	y = 8
	fmt.Println(y)
	xi := []int{1, 2, 4, 7, 9, 42} ///slice of int

	fmt.Println(xi)

	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}
	fmt.Print(m)
	var t hotdog
	t = 7
	fmt.Println(t)

	p1 := person{
		"Miss",
		"MoneyPenny",
	}
	fmt.Println(p1)
}

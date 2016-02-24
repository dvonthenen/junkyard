package main

import (
	"fmt"
	"math"
)

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

type Circle struct {
	x float64
	y float64
	r float64
}

func areaCircle(c Circle) float64 {
	return math.Pi * c.r * c.r
}

//nested structs
type Person struct {
	Name string
}
type Android struct {
	person Person
	model  string
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)

	var c1 Circle
	fmt.Println(c1)

	c2 := Circle{1, 2, 3}
	fmt.Println("x", c2.x)

	fmt.Println(areaCircle(c2))

}

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func areaCircle(c Circle) float64 {
	return math.Pi * c.r * c.r
}

//instead of the above... lets create a method off Circle
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

//interface
type MyTalk interface {
	Talk()
}

//nested structs
type Person struct {
	Name string
}
type Android struct {
	person Person
	model  string
}

func (p Person) Talk() {
	fmt.Println("My name is", p.Name)
}

func MakeMeTalk(talk MyTalk) {
	talk.Talk()
}

func main() {

	var c1 Circle
	fmt.Println(c1)

	c2 := Circle{1, 2, 3}
	fmt.Println("x", c2.x)

	fmt.Println(areaCircle(c2))
	//do this instead
	fmt.Println(c2.area())

	a := new(Android)
	a.person.Name = "Data"
	a.person.Talk()

	MakeMeTalk(a.person)

}

package main

import "fmt"

func main() {
	//var x string
	//x = "Hello World"

	var x string = "Hello World"
	fmt.Println(x)

	x = x + "!!!!"
	fmt.Println(x)

	var y = "world" //you can omit "string" as it is inferred
	fmt.Println(x == y)

	x = "world"
	fmt.Println(x == y)

	z := 5 //redefines the type
	fmt.Println(z)

	const w string = "My World"
	fmt.Println(w)

}

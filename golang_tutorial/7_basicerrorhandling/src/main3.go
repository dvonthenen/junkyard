package main

import "fmt"

func main() {

	//in the code below, panic causes an immediate exit which doesnt call recover()
	//panic("PANIC")
	//str := recover()
	//fmt.Println(str)

	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("PANIC")

}

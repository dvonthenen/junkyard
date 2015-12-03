package main

import "fmt"
import "os"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	defer second() //called on leaving the function... equivalent to finally
	first()

	f, _ := os.Open("filename")
	defer f.Close()
}

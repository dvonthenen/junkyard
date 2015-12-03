package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			fmt.Println(j, "Even")
		} else {
			fmt.Println(j, "Odd")
		}
	}

	for k := 0; k < 5; k++ {
		switch k {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		default:
			fmt.Println("Unknown")
		}
	}
}

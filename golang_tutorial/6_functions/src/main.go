package main

import "fmt"

func average(xs []float64) float64 { //normal function
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func add(args ...int) int { //variable params
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	xs := []float64{98, 93, 2, 55, 159}
	fmt.Println(average(xs))

	fmt.Println(add(3, 5, 2))

	//inline function within a function in this case main()
	multi := func(args ...int) int {
		total := 1
		for _, v := range args {
			total *= v
		}
		return total
	}

	fmt.Println(multi(2, 3, 4))
}

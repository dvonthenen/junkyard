package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	//define flags
	maxp := flag.Int("max", 6, "the max value")
	//parse
	flag.Parse()
	//generate a number
	fmt.Println(rand.Intn(*maxp))

	//Args
	var args []string
	args = flag.Args()

	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

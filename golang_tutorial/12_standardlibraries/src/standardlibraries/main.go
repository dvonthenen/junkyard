package main

import (
	"container/list"
	"crypto/sha1"
	"fmt"
	"os"
	"strings"
)

func main() {
	//strings
	fmt.Println(
		//true
		strings.Contains("test", "es"),
		//2
		strings.Count("test", "t"),
		//true
		strings.HasPrefix("test", "te"),
	)

	//os
	file, err := os.Create("test.txt")
	if err != nil {
		panic("PANIC")
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic("PANIC")
	}

	fmt.Println(stat.Size())

	//list
	var x list.List
	x.PushBack(2)
	x.PushBack(1)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	//hashing
	h := sha1.New()
	h.Write([]byte("hello"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)

	//TODO
}

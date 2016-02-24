package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func client() {
	//connect
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	//send
	msg := "Hello!"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	client()
}

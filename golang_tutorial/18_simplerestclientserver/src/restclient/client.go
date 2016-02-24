package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Payload struct {
	Car   Cars
	Truck Trucks
}

type Cars map[string]int
type Trucks map[string]int

func main() {
	url := "http://127.0.0.1:9000"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var p Payload
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p.Car)
	fmt.Println(p.Truck)
}

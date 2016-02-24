package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Car   Cars
	Truck Trucks
}

type Cars map[string]int
type Trucks map[string]int

func serveRest(w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(response))
}

func main() {
	http.HandleFunc("/", serveRest)
	http.ListenAndServe(":9000", nil)
}

func getJsonResponse() ([]byte, error) {
	cars := make(map[string]int)
	cars["350z"] = 11
	cars["GT500"] = 1

	trucks := make(map[string]int)
	trucks["F-150"] = 3

	p := Payload{cars, trucks}

	return json.MarshalIndent(p, "", "  ")
}

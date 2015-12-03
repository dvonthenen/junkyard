package main

import "fmt"

func main() {

	//arrays
	var x [5]float64
	x[2] = 37
	x[3] = 23
	x[4] = 100
	fmt.Println(x)

	var total float64 = 0
	for i, value := range x {
		fmt.Println(i, value)
	}
	for _, value := range x { //foreach
		total += value
	}
	fmt.Println(total / float64(len(x)))

	arr1 := make([]int, 5, 10)
	//arr1 and arr2 are pretty much the same
	arr := []int{1, 2, 3, 4, 5}
	arr2 := arr[0:5]
	fmt.Println(arr2)

	arr3 := append(arr1, 6, 7)
	fmt.Println(arr3)

	//maps
	mymap := make(map[string]string)
	mymap["key1"] = "value1"
	mymap["key2"] = "value2"
	mymap["key3"] = "value3"
	mymap["key4"] = "value4"
	fmt.Println(mymap)

	delete(mymap, "key3")
	fmt.Println(mymap)

	fmt.Println(mymap["key3"])
	fmt.Println(mymap["key2"])

	name1, ok1 := mymap["key2"]
	fmt.Println(name1, ok1)

	name2, ok2 := mymap["key3"]
	fmt.Println(name2, ok2)

	if name3, ok3 := mymap["key2"]; ok3 {
		fmt.Println(name3, ok3)
	}

	if name4, ok4 := mymap["key3"]; ok4 { //skips because ok is false
		fmt.Println(name4, ok4)
	}
}

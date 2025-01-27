package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	mymap := map[string]any{"name": "Josh", "age": 25}

	for key, value := range mymap {
		fmt.Println(key, value)
	}

	x := func() {
		fmt.Println("this is a function")
	}
	x()

	fmt.Println(mymap)
}

package main

import "fmt"

func main() {
	var a interface{}
	a = "Hello Nipun Bunket"
	whichTime(a)

}

func whichTime(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("It's a string %s\n", i)
	case int:
		fmt.Printf("It's an int %d\n", i)
	case float64:
		fmt.Printf("It's a float64%f\n", i)
	default:
		fmt.Println("I don't know what it is")
	}
}

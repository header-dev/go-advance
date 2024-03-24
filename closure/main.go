package main

import "fmt"

func main() {

	fn := newCounter()

	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())

}

func newCounter() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

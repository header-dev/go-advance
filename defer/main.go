package main

import "fmt"

func main() {
	defer fmt.Println("Start #1")
	defer fmt.Println("Start #2")
	defer fmt.Println("Start #3")
	dontForget(3)
	fmt.Println("End")

}

func dontForget(n int) {
	defer fmt.Println(n)
	defer func() {
		fmt.Println(n)
	}()

	n += n
}

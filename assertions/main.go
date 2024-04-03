package main

import "fmt"

func main() {
	var a interface{}
	a = 12
	fmt.Printf("%T %v\n", a, a)

	if i, ok := a.(int); ok {
		fmt.Printf("%T %v\n", i, i)
	}

	var b, c Obj
	c = rectangle{w: 3, l: 4}
	b = triangle{w: 3, h: 4}

	// fmt.Println(c.Area())
	fmt.Println(b.Area())

	if v, ok := c.(rectangle); ok {
		v.l = 5
		v.w = 6
		fmt.Println(v.Area())
	}

}

type Obj interface {
	Area() float64
}

type rectangle struct {
	w, l float64
}

func (rec rectangle) Area() float64 {
	return rec.w * rec.l
}

type triangle struct {
	w, h float64
}

func (tri triangle) Area() float64 {
	return tri.h * tri.w * 0.5
}

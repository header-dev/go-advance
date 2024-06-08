package main

import (
	"fmt"
	"sync"
)

type thing struct {
	i   int
	mux sync.Mutex
}

var t = thing{}

func main() {
	go func() {
		for {
			fmt.Println(t.read())
		}
	}()

	for i := 0; i < 10; i++ {
		t.write(i)
	}
}

func (t *thing) write(n int) {
	t.mux.Lock()
	t.i = n
	t.mux.Unlock()
}

func (t *thing) read() int {
	t.mux.Lock()
	defer t.mux.Unlock()
	return t.i
}

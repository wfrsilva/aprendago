package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	novasgoroutines(2)
	wg.Wait()
}

func novasgoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		x := j
		go func(i int) {
			fmt.Println("Goroutine nro:", i+1)
			wg.Done()
		}(x)
	}
}

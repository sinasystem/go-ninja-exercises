package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	grs := 1000

	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			defer wg.Done()
			temp := counter
			runtime.Gosched()
			temp++
			counter = temp
		}()
	}
	wg.Wait()
	fmt.Println("Count:", counter)
}

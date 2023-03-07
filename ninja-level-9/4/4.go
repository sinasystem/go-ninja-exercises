package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	grs := 1000

	var wg sync.WaitGroup
	wg.Add(grs)

	var mx sync.Mutex

	for i := 0; i < grs; i++ {
		go func() {
			mx.Lock()
			defer wg.Done()
			defer mx.Unlock()
			temp := counter
			temp++
			counter = temp
		}()
	}
	wg.Wait()
	fmt.Println("Count:", counter)
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	gr := 2
	wg.Add(2)

	for i := 0; i < gr; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("Something")
		}()
	}
	wg.Wait()
}

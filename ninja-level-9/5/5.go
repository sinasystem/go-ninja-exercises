package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	grs := 1000

	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Count:", atomic.LoadInt64(&counter))

		}()
	}
	wg.Wait()
	fmt.Println("Count:", counter)

}

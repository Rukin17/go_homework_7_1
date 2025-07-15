package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Горутина %d начала работу\n", id)
	time.Sleep(5 * time.Millisecond)
	fmt.Printf("Горутина %d завершила раблоту\n", id)
}

func main() {
	var wg sync.WaitGroup
	countWorker := 3
	for i := 1; i <= countWorker; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}

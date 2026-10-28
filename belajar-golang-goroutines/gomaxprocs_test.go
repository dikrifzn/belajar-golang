package main

import (
	"fmt"
	"testing"
	"runtime"
	"sync"
	"time"
)

func TestGOMAXPROCS(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		group.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			group.Done()
		}()
	}
	totalcpus := runtime.NumCPU()
	fmt.Println("Total CPUs:", totalcpus)

	// runtime.GOMAXPROCS(4)
	totalthreads := runtime.GOMAXPROCS(0)
	fmt.Println("Total Threads:", totalthreads)

	totalgoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines:", totalgoroutines)
}
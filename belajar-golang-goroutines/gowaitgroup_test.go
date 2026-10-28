package main

import (
	"fmt"
	"testing"
	"sync"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done() // menandakan bahwa goroutine telah selesai
	group.Add(1) // menambahkan jumlah goroutine yang sedang berjalan
	fmt.Println("Hello World")
	time.Sleep(2 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}
	
	group.Wait() // menunggu semua goroutine selesai
	fmt.Println("Semua goroutine selesai")
}
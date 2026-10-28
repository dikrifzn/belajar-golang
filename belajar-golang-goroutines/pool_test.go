package main

import (
	"testing"
	"fmt"
	"sync"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Dikri")
	pool.Put("Belajar Golang")
	pool.Put("Goroutines")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(2 * time.Second)
			pool.Put(data)
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}
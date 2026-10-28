package main

import (
	"fmt"
	"testing"
	"time"
	"sync"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now(), "Timer started")

	time := <-timer.C
	fmt.Println(time, "Timer expired")
}

func TestAfter(t *testing.T) {
	after := time.After(5 * time.Second)
	fmt.Println(time.Now(), "After started")

	time := <-after
	fmt.Println(time, "After expired")
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now(), "AfterFunc executed")
		group.Done()
	})
	fmt.Println(time.Now(), "AfterFunc started")
	group.Wait()
}
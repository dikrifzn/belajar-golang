package main
import (
	"sync"
	"fmt"
	"time"
	"testing"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()

	group.Add(1)
	cond.L.Lock()
	cond.Wait()

	fmt.Println("Condition met for value:", value)

	cond.L.Unlock()
}

func TestCondition(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		fmt.Println("Signaling condition for value:", i)
	// 		cond.Signal()
	// 	}
	// }()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Broadcast()
		}
	}()

	group.Wait() 
}
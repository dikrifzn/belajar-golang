package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	
	for time := range ticker.C {
		fmt.Println(time, "Ticker ticked")
	}
}

func TestTick(t *testing.T) {
	tick := time.Tick(1 * time.Second)
	for time := range tick {
		fmt.Println(time, "Tick ticked")
	}
}
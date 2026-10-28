package main
import (
    "fmt"
    "sync"
    "testing"
    "time"
)

func TestWithoutGoroutine(t *testing.T) {
    start := time.Now()

    time.Sleep(2 * time.Second)
    fmt.Println("Pekerjaan A selesai")

    time.Sleep(1 * time.Second)
    fmt.Println("Pekerjaan B selesai")

    fmt.Println("Total:", time.Since(start))
}

func TestWithGoroutine(t *testing.T) {
    start := time.Now()

    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        time.Sleep(1 * time.Second)
        fmt.Println("Pekerjaan A selesai")
    }()

    go func() {
        defer wg.Done()
        time.Sleep(2 * time.Second)
        fmt.Println("Pekerjaan B selesai")
    }()

    wg.Wait()

    fmt.Println("Total:", time.Since(start))
}
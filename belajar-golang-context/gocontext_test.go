package main
import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println("background:", background)

	todo := context.TODO()
	fmt.Println("todo:", todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	conextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(conextE)
	fmt.Println(contextF)

	// fmt.Println(contextF.Value("f")) //ada
	// fmt.Println(contextF.Value("c")) //tidak memiliki value c
	// fmt.Println(contextF.Value("b")) //tidak memiliki value b
	// fmt.Println(contextA.Value("b")) //tidak memiliki value b
}





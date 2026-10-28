package main

import (
	"fmt"
	_ "belajar_golang_standar/belajar_fmt"
	"belajar_golang_standar/belajar_errors"
)

func main() {
	// belajar_fmt.BelajarFmtFormat()



	err := belajar_errors.GetById("")

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}


	
}
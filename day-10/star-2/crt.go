package main

import (
	"fmt"
	"log"
)

func draw(registerValues []int) {
	if len(registerValues) != 240 {
		log.Fatal("incorrect input length")
	}
	for idx, registerVal := range registerValues {
		column := idx % 40
		if column <= registerVal+1 && column >= registerVal-1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if column == 39 {
			fmt.Print("\n")
		}
	}
}

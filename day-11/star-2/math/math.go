package main

import "log"

var (
	bigNum = 969969012345
	nums   = []int{2, 17, 19, 3, 5, 13, 7, 11}
)

func main() {
	composite := 2 * 17 * 19 * 3 * 5 * 13 * 7 * 11
	log.Println("Magic number:", bigNum)
	log.Println("Composite:", composite)
	log.Println()
	log.Println("Normal results:")
	log.Println("---------------")
	for _, num := range nums {
		log.Printf("%%%d: %d", num, bigNum%num)
	}
	log.Println()
	log.Println("Composite results:")
	log.Println("------------------")
	for _, num := range nums {
		log.Printf("%%%d: %d", num, (bigNum%composite)%num)
	}
}

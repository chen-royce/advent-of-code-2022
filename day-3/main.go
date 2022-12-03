package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	filepath := "./problem/input.txt"
	input, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	var priorities int

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		currSack := scanner.Text()
		sack1 := currSack[:len(currSack)/2]
		sack2 := currSack[len(currSack)/2:]

		letterMap := make(map[uint8]struct{})

		for idx, _ := range sack1 {
			letterMap[sack1[idx]] = struct{}{}
		}

		for idx, _ := range sack2 {
			if _, ok := letterMap[sack2[idx]]; ok {
				log.Println(string(sack2[idx]))
				priorities += getPointValueOfLetter(sack2[idx])
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(priorities)
}

func getPointValueOfLetter(letter uint8) int {
	// lower case
	if letter >= 97 && letter <= 122 {
		return int(letter - 96)
	}
	// upper case
	if letter >= 65 && letter <= 90 {
		return int(letter - 38)
	}
	return -1
}

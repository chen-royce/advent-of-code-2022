package main

import (
	"bufio"
	"log"
	"os"
	"strings"
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

	var groupOf3 []string

	for scanner.Scan() {
		groupOf3 = append(groupOf3, scanner.Text())
		if len(groupOf3) >= 3 {
			sack1 := groupOf3[0]
			sack2 := groupOf3[1]
			sack3 := groupOf3[2]

			for idx, _ := range sack1 {
				if strings.Contains(sack2, string(sack1[idx])) && strings.Contains(sack3, string(sack1[idx])) {
					priorities += getPointValueOfLetter(sack1[idx])
					break
				}
			}

			// Reset group to be empty
			groupOf3 = []string{}
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

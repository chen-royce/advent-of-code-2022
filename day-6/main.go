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

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		input := scanner.Text()

		// array of encountered
		var encountered []rune

		for idx, char := range input {
			log.Println("CHAR:", char)
			log.Println("ENCOUNTERED:", encountered)

			// if we find something we've encountered, we need to reset
			// our encountered array
			if foundIdx := findInSlice(char, encountered); foundIdx != -1 {
				encountered = append(encountered[foundIdx+1:], char)
				continue
			}

			encountered = append(encountered, char)

			// if length 3, yay - we found the 4th
			if len(encountered) == 4 {
				log.Println(idx + 1)
				break
			}

			// if length >3, need to bump 1
			if len(encountered) == 5 {
				encountered = append(encountered, char)
				encountered = encountered[1:]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func findInSlice(needle rune, haystack []rune) int {
	for idx, r := range haystack {
		if r == needle {
			return idx
		}
	}
	return -1
}

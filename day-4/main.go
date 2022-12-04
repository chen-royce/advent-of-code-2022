package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filepath := "./problem/input.txt"
	input, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var numRedundant int

	for scanner.Scan() {
		assignmentStrings := strings.Split(scanner.Text(), ",")

		assignment1String := assignmentStrings[0]
		assignment1Bounds, err := getBounds(assignment1String)
		if err != nil {
			log.Fatal(err)
		}

		assignment2String := assignmentStrings[1]
		assignment2Bounds, err := getBounds(assignment2String)
		if err != nil {
			log.Fatal(err)
		}

		// Right of assignment 1 overlaps with assignment 2
		// example: 1-3,3-5
		if assignment1Bounds[1] >= assignment2Bounds[0] && assignment1Bounds[1] <= assignment2Bounds[1] {
			numRedundant++
			continue
		}

		// Left of assignment 1 overlaps with assignment 2
		// example: 3-4,1-3
		if assignment1Bounds[0] >= assignment2Bounds[0] && assignment1Bounds[0] <= assignment2Bounds[1] {
			numRedundant++
			continue
		}

		// Otherwise, check if assignment 1 contains assignment 2
		if assignment1Bounds[0] <= assignment2Bounds[0] && assignment1Bounds[1] >= assignment2Bounds[1] {
			numRedundant++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(numRedundant)
}

func getBounds(input string) ([]int, error) {
	splitStringArray := strings.Split(input, "-")
	log.Println("SPLIT", splitStringArray)
	lowerBound, err := strconv.Atoi(splitStringArray[0])
	if err != nil {
		return nil, err
	}
	upperBound, err := strconv.Atoi(splitStringArray[1])
	if err != nil {
		return nil, err
	}
	return []int{lowerBound, upperBound}, nil
}

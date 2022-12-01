package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	filepath := "./problem/input.txt"
	log.Println("Top calories:", getTopCalories(filepath))
	log.Println("Top 3 calories:", getTop3Calories(filepath))
	log.Println("Sum of top 3 calories:", sumTop3Calories(filepath))
}

func sumTop3Calories(filepath string) int {
	var acc int
	for _, count := range getTop3Calories(filepath) {
		acc += count
	}
	return acc
}

func getTopCalories(filepath string) int {
	return getTop3Calories(filepath)[2]
}

func getTop3Calories(filepath string) []int {
	input, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var currCalories int
	max3Calories := make([]int, 3)

	for scanner.Scan() {
		currString := scanner.Text()

		// Updates totals after detecting empty line
		if currString == "" {
			if currCalories > max3Calories[2] {
				max3Calories = []int{max3Calories[1], max3Calories[2], currCalories}
			} else if currCalories > max3Calories[1] {
				max3Calories = []int{max3Calories[1], currCalories, max3Calories[2]}
			} else if currCalories > max3Calories[0] {
				max3Calories = []int{currCalories, max3Calories[1], max3Calories[2]}
			}
			currCalories = 0
		} else { // Otherwise, keeps accumulating
			currNum, err := strconv.Atoi(currString)
			if err != nil {
				log.Fatal(err)
			}
			currCalories += currNum
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// One extra check for EOF, which is not terminated w/empty line
	if currCalories > max3Calories[2] {
		max3Calories = []int{max3Calories[1], max3Calories[2], currCalories}
	} else if currCalories > max3Calories[1] {
		max3Calories = []int{max3Calories[1], currCalories, max3Calories[2]}
	} else if currCalories > max3Calories[0] {
		max3Calories = []int{currCalories, max3Calories[1], max3Calories[2]}
	}

	return max3Calories
}

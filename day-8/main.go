package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filepath := "./problem/input.txt"
	// filepath := "./problem/sample.txt"
	forest, err := buildForestFromFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(forest)

	visibilityMap := findVisibleTrees(forest)
	log.Println(visibilityMap)
	log.Println(len(visibilityMap))
}

type Forest [][]int

func (f Forest) height() int {
	return len(f)
}

func (f Forest) width() int {
	return len(f[0])
}

func buildForestFromFile(filepath string) (Forest, error) {
	input, err := os.Open(filepath)
	if err != nil {
		return [][]int{}, errors.New("open file")
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var rows [][]int

	for scanner.Scan() {
		row, err := parseRow(scanner.Text())
		if err != nil {
			return [][]int{}, err
		}
		rows = append(rows, row)
	}

	if err := scanner.Err(); err != nil {
		return [][]int{}, errors.New("error scanning file text")
	}

	return rows, nil
}

func parseRow(s string) ([]int, error) {
	var ret []int
	for idx := range s {
		num, err := strconv.Atoi(s[idx : idx+1])
		if err != nil {
			return []int{}, errors.New("parsing row")
		}
		ret = append(ret, num)
	}
	return ret, nil
}

func findVisibleTrees(forest Forest) map[string]struct{} {
	visibilityMap := make(map[string]struct{})

	LookDown(forest, visibilityMap)  // look down from top of forest
	LookUp(forest, visibilityMap)    // look up from bottom of forest
	LookRight(forest, visibilityMap) // look right from left of forest
	LookLeft(forest, visibilityMap)  // look left from right of forest

	return visibilityMap
}

func LookDown(forest Forest, visibilityMap map[string]struct{}) {
	for col := 0; col < forest.width(); col++ {
		minHeight := -1
		for row := 0; row < forest.height(); row++ {
			currHeight := forest[row][col]
			if currHeight > minHeight {
				key := createTreeCoordinateKey(row, col)
				visibilityMap[key] = struct{}{}
				minHeight = currHeight
			}
		}
	}
}

func LookUp(forest Forest, visibilityMap map[string]struct{}) {
	for col := 0; col < forest.width(); col++ {
		minHeight := -1
		for row := forest.height() - 1; row >= 0; row-- {
			currHeight := forest[row][col]
			if currHeight > minHeight {
				key := createTreeCoordinateKey(row, col)
				visibilityMap[key] = struct{}{}
				minHeight = currHeight
			}
		}
	}
}

func LookLeft(forest Forest, visibilityMap map[string]struct{}) {
	for row := 0; row < forest.height(); row++ {
		minHeight := -1
		for col := forest.width() - 1; col >= 0; col-- {
			currHeight := forest[row][col]
			if forest[row][col] > minHeight {
				key := createTreeCoordinateKey(row, col)
				visibilityMap[key] = struct{}{}
				minHeight = currHeight
			}
		}
	}
}

func LookRight(forest Forest, visibilityMap map[string]struct{}) {
	for row := 0; row < forest.height(); row++ {
		minHeight := -1
		for col, currHeight := range forest[row] {
			if currHeight > minHeight {
				key := createTreeCoordinateKey(row, col)
				visibilityMap[key] = struct{}{}
				minHeight = currHeight
			}
		}
	}
}

func createTreeCoordinateKey(row, col int) string {
	rowString := strconv.Itoa(row)
	colString := strconv.Itoa(col)
	return fmt.Sprintf("%s,%s", rowString, colString)
}

package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// parse input and generate adjacency matrix
	// filepath := "./problem/input.txt"
	// filepath := "./problem/sample.txt"
	filepath := "./problem/custom.txt"
	input, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	vals := parseInput(input)
	adjList := buildAdjacencyList(vals)
	log.Println(adjList)
}

func parseInput(in []byte) [][]uint8 {
	var ret [][]uint8
	rows := strings.Split(string(in), "\n")
	for i := range rows {
		var currRow []byte
		for j := range rows[i] {
			currRow = append(currRow, rows[i][j])
		}
		ret = append(ret, currRow)
	}
	return ret
}

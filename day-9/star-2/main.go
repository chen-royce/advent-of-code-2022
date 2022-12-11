package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// filepath := "./problem/input-simple.txt"
	filepath := "../problem/input.txt"
	moves, err := ParseInput(filepath)
	if err != nil {
		log.Fatal(err)
	}
	head := newLink(10)
	for _, move := range moves {
		head.moveLink(move)
	}
	tailPastPositions := getTailPastPositions(head)
	log.Println(len(tailPastPositions))
}

type move struct {
	Direction string
	Distance  int
}

func ParseInput(filepath string) ([]move, error) {
	input, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var moves []move

	for scanner.Scan() {
		splitText := strings.Split(scanner.Text(), " ")
		direction := splitText[0]
		distance, err := strconv.Atoi(splitText[1])
		if err != nil {
			return []move{}, err
		}
		moves = append(moves, move{direction, distance})
	}

	if err := scanner.Err(); err != nil {
		return []move{}, err
	}

	return moves, nil
}

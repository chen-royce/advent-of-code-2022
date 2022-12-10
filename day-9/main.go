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
	filepath := "./problem/input.txt"
	moves, err := ParseInput(filepath)
	if err != nil {
		log.Fatal(err)
	}
	rope := NewRope(position{}, position{})
	for _, move := range moves {
		rope.MoveRope(move)
	}
	log.Println(len(rope.PastTailPositions))
}

// functionalities still to come:
// tracking where tail has been
// func to format positions into map keys

type Move struct {
	Direction string
	Distance  int
}

func ParseInput(filepath string) ([]Move, error) {
	input, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var moves []Move

	for scanner.Scan() {
		splitText := strings.Split(scanner.Text(), " ")
		direction := splitText[0]
		distance, err := strconv.Atoi(splitText[1])
		if err != nil {
			return []Move{}, err
		}
		moves = append(moves, Move{direction, distance})
	}

	if err := scanner.Err(); err != nil {
		return []Move{}, err
	}

	return moves, nil
}

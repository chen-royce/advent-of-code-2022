package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3

	LOSS = 0
	DRAW = 3
	WIN  = 6
)

var winningCombos = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var drawingCombos = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

func main() {
	filepath := "./problem/input.txt"
	input, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var score int

	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		opponentMove := moves[0]
		yourMove := moves[1]

		log.Println("yourMove, opponentMove", yourMove, opponentMove)

		// add your value
		switch yourMove {
		case "X":
			score += ROCK
			log.Println("Adding rock score")
		case "Y":
			score += PAPER
			log.Println("Adding paper score")
		case "Z":
			score += SCISSORS
			log.Println("Adding scissors score")
		}

		// check for win
		if winningCombos[opponentMove] == yourMove {
			score += 6
			log.Println("won")
		}

		// check for draw
		if drawingCombos[opponentMove] == yourMove {
			score += 3
			log.Println("draw")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(score)
}

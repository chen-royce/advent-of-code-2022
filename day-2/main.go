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

var losingCombos = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
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

		var yourThrow string

		// add your value
		switch yourMove {
		case "X":
			score += LOSS
			yourThrow = losingCombos[opponentMove]
		case "Y":
			score += DRAW
			yourThrow = drawingCombos[opponentMove]
		case "Z":
			score += WIN
			yourThrow = winningCombos[opponentMove]
		}

		// check your input
		switch yourThrow {
		case "X":
			score += ROCK
		case "Y":
			score += PAPER
		case "Z":
			score += SCISSORS
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(score)
}

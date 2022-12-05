package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	stack1 = []string{"D", "L", "J", "R", "V", "G", "F"}
	stack2 = []string{"T", "P", "M", "B", "V", "H", "J", "S"}
	stack3 = []string{"V", "H", "M", "F", "D", "G", "P", "C"}
	stack4 = []string{"M", "D", "P", "N", "G", "Q"}
	stack5 = []string{"J", "L", "H", "N", "F"}
	stack6 = []string{"N", "F", "V", "Q", "D", "G", "T", "Z"}
	stack7 = []string{"F", "D", "B", "L"}
	stack8 = []string{"M", "J", "B", "S", "V", "D", "N"}
	stack9 = []string{"G", "L", "D"}

	stacks = [][]string{
		stack1,
		stack2,
		stack3,
		stack4,
		stack5,
		stack6,
		stack7,
		stack8,
		stack9,
	}
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
		sentence := strings.Split(scanner.Text(), " ")
		if sentence[0] != "move" {
			continue
		}
		log.Println(sentence)
		action := parseSentence(sentence)

		for i := 0; i < action[0]; i++ {
			src := stacks[action[1]]
			dst := stacks[action[2]]
			popped := src[len(src)-1]

			src = src[:len(src)-1]
			dst = append(dst, popped)

			stacks[action[1]] = src
			stacks[action[2]] = dst

			log.Println("SRC", stacks[action[1]])
			log.Println("DST", stacks[action[2]])
		}
	}

	// log final answer
	for _, stack := range stacks {
		log.Print(stack[len(stack)-1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseSentence(strings []string) []int {
	numToMove, err := strconv.Atoi(strings[1])
	if err != nil {
		log.Fatal(err)
	}
	src, err := strconv.Atoi(strings[3])
	if err != nil {
		log.Fatal(err)
	}
	dst, err := strconv.Atoi(strings[5])
	if err != nil {
		log.Fatal(err)
	}

	// since stacks aren't 0 indexed
	// for example, stack 1 would be at index 0
	src--
	dst--

	return []int{numToMove, src, dst}
}

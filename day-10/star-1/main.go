package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	// filepath := "./problem/sample.txt"
	filepath := "./problem/input.txt"
	commands, _ := parseInput(filepath)
	device := newDevice()

	device.executeCommands(commands)
	device.sumSignalStrengths(20, 60, 100, 140, 180, 220)
}

func parseInput(filepath string) ([][]string, error) {
	input, err := os.Open(filepath)
	if err != nil {
		return [][]string{}, err
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var commands [][]string

	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Split(line, " ")
		commands = append(commands, command)
	}

	if err := scanner.Err(); err != nil {
		return [][]string{}, err
	}

	return commands, nil
}

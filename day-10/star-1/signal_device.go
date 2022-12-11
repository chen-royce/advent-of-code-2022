package main

import (
	"log"
	"strconv"
)

type device struct {
	cycle        int
	register     int
	cycleHistory []int
}

func newDevice() device {
	return device{
		cycle:        0,
		register:     1,
		cycleHistory: []int{1},
	}
}

func (d *device) executeCommands(commands [][]string) {
	for _, command := range commands {
		switch len(command) {
		case 1: // noop command
			d.noop()
		case 2: // addx command
			amount := command[1]
			d.addx(amount)
		}
	}
}

func (d *device) sumSignalStrengths(cycleNumbers ...int) int {
	var total int
	for _, cycleNumber := range cycleNumbers {
		// Note: The cycleNumbers provided by AOC aren't 0-indexed, so we reference cycleNumber - 1
		// when looking through our internal arrays
		log.Printf("On cycle %v, the register value was %v\n", cycleNumber, d.cycleHistory[cycleNumber-1])
		log.Printf("During that cycle, signal strength was %v\n", cycleNumber*d.cycleHistory[cycleNumber-1])
		total += cycleNumber * d.cycleHistory[cycleNumber-1]
	}
	log.Printf("The sum of these signal strengths is %v\n", total)
	return total
}

func (d *device) noop() {
	d.cycle++
	d.cycleHistory = append(d.cycleHistory, d.register)
}

func (d *device) addx(amount string) {
	d.noop()

	if signalChange, err := strconv.Atoi(amount); err != nil {
		log.Fatal(err)
	} else {
		d.register += signalChange
	}

	d.noop()
}

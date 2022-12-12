package main

import "log"

func main() {
	// ACTUAL INPUT
	monkey0 := newMonkey(
		"monkey0",
		[]int{83, 72, 93},
		func(i int) int { return i * 17 },
		func(i int) bool { return i%2 == 0 },
	)
	monkey1 := newMonkey(
		"monkey1",
		[]int{90, 55},
		func(i int) int { return i + 1 },
		func(i int) bool { return i%17 == 0 },
	)
	monkey2 := newMonkey(
		"monkey2",
		[]int{91, 78, 80, 97, 79, 88},
		func(i int) int { return i + 3 },
		func(i int) bool { return i%19 == 0 },
	)
	monkey3 := newMonkey(
		"monkey3",
		[]int{64, 80, 83, 89, 59},
		func(i int) int { return i + 5 },
		func(i int) bool { return i%3 == 0 },
	)
	monkey4 := newMonkey(
		"monkey4",
		[]int{98, 92, 99, 51},
		func(i int) int { return i * i },
		func(i int) bool { return i%5 == 0 },
	)
	monkey5 := newMonkey(
		"monkey5",
		[]int{68, 57, 95, 85, 98, 75, 98, 75},
		func(i int) int { return i + 2 },
		func(i int) bool { return i%13 == 0 },
	)
	monkey6 := newMonkey(
		"monkey6",
		[]int{74},
		func(i int) int { return i + 4 },
		func(i int) bool { return i%7 == 0 },
	)
	monkey7 := newMonkey(
		"monkey7",
		[]int{68, 64, 60, 68, 87, 80, 82},
		func(i int) int { return i * 19 },
		func(i int) bool { return i%11 == 0 },
	)

	monkey0.trueDestination = monkey1
	monkey0.falseDestination = monkey6
	monkey1.trueDestination = monkey6
	monkey1.falseDestination = monkey3
	monkey2.trueDestination = monkey7
	monkey2.falseDestination = monkey5
	monkey3.trueDestination = monkey7
	monkey3.falseDestination = monkey2
	monkey4.trueDestination = monkey0
	monkey4.falseDestination = monkey1
	monkey5.trueDestination = monkey4
	monkey5.falseDestination = monkey0
	monkey6.trueDestination = monkey3
	monkey6.falseDestination = monkey2
	monkey7.trueDestination = monkey4
	monkey7.falseDestination = monkey5

	// SAMPLE INPUT
	// monkey0 := &monkey{
	// 	name:  "monkey 0",
	// 	items: []int{79, 98},
	// 	operation: func(i int) int {
	// 		return 19 * i
	// 	},
	// 	passesTest: func(i int) bool {
	// 		return i%23 == 0
	// 	},
	// }
	// monkey1 := &monkey{
	// 	name:  "monkey 1",
	// 	items: []int{54, 65, 75, 74},
	// 	operation: func(i int) int {
	// 		return 6 + i
	// 	},
	// 	passesTest: func(i int) bool {
	// 		return i%19 == 0
	// 	},
	// }
	// monkey2 := &monkey{
	// 	name:  "monkey 2",
	// 	items: []int{79, 60, 97},
	// 	operation: func(i int) int {
	// 		return i * i
	// 	},
	// 	passesTest: func(i int) bool {
	// 		return i%13 == 0
	// 	},
	// }
	// monkey3 := &monkey{
	// 	name:  "monkey 3",
	// 	items: []int{74},
	// 	operation: func(i int) int {
	// 		return 3 + i
	// 	},
	// 	passesTest: func(i int) bool {
	// 		return i%17 == 0
	// 	},
	// }

	// monkey0.trueDestination = monkey2
	// monkey0.falseDestination = monkey3
	// monkey1.trueDestination = monkey2
	// monkey1.falseDestination = monkey0
	// monkey2.trueDestination = monkey1
	// monkey2.falseDestination = monkey3
	// monkey3.trueDestination = monkey0
	// monkey3.falseDestination = monkey1

	monkeys := []*monkey{monkey0, monkey1, monkey2, monkey3, monkey4, monkey5, monkey6, monkey7}

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			monkey.enqueueItems()
			monkey.throwItems()
		}
	}

	for _, monkey := range monkeys {
		log.Printf("%s: %v items inspected\n", monkey.name, monkey.itemsInspected)
	}
}

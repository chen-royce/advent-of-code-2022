package main

var (
	// sample input
	// composite = 23 * 19 * 13 * 17
	composite = 2 * 17 * 19 * 3 * 5 * 13 * 7 * 11
)

type monkey struct {
	name             string
	items            []int
	operation        func(int) int
	passesTest       func(int) bool
	trueItems        []int
	falseItems       []int
	trueDestination  *monkey
	falseDestination *monkey
	itemsInspected   int
}

func newMonkey(name string, items []int, operation func(int) int, passesTest func(int) bool) *monkey {
	return &monkey{
		name:       name,
		items:      items,
		operation:  operation,
		passesTest: passesTest,
	}
}

func findMonkeyBusiness(input []int) int {
	return input[0] * input[1]
}

func findTopTwoMonkeys(monkeys []*monkey) []int {
	var first, second int
	for _, monkey := range monkeys {
		if monkey.itemsInspected > first && monkey.itemsInspected > second {
			second = first
			first = monkey.itemsInspected
		} else if monkey.itemsInspected > second {
			second = monkey.itemsInspected
		}
	}
	return []int{first, second}
}

func (m *monkey) enqueueItems() {
	for _, item := range m.items {
		// run operation
		item = m.operation(item)

		// reduce stress 😌
		if item >= composite {
			item %= composite
		}

		// run test and enqueue in destination array
		if m.passesTest(item) {
			m.trueItems = append(m.trueItems, item)
		} else {
			m.falseItems = append(m.falseItems, item)
		}

		m.itemsInspected++
	}
	// clear monkey's current items
	m.items = []int{}
}

func (m *monkey) throwItems() {
	m.trueDestination.items = append(m.trueDestination.items, m.trueItems...)
	m.falseDestination.items = append(m.falseDestination.items, m.falseItems...)
	m.trueItems, m.falseItems = []int{}, []int{}
}

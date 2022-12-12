package main

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

func (m *monkey) enqueueItems() {
	for _, item := range m.items {
		// run operation
		item = m.operation(item)

		// divide by 3 and floor
		item /= 3

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

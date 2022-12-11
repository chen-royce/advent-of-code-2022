package main

import (
	"fmt"
	"math"
	"strconv"
)

type link struct {
	index            int
	parent           *link
	child            *link
	position         position
	visitedPositions map[string]struct{}
}

type position struct {
	X int
	Y int
}

func newLink(length int) *link {
	if length <= 0 {
		return nil
	}
	ret := link{
		index: 0,
	}
	currLink := &ret
	for i := 1; i < length; i++ {
		nextLink := link{
			index:  i,
			parent: currLink,
			visitedPositions: map[string]struct{}{
				visitedPositionsKey(position{}): {},
			},
		}
		currLink.child = &nextLink
		currLink = &nextLink
	}
	return &ret
}

func (l *link) moveLink(m move) *link {
	if m.Distance == 0 {
		return l
	}
	// Move head 1 unit in correct direction
	switch m.Direction {
	case "U":
		l.position.Y++
	case "D":
		l.position.Y--
	case "L":
		l.position.X--
	case "R":
		l.position.X++
	}

	l.updateChildrenPositions()

	// Decrement move distance and recursively continue
	// until the move is complete
	m.Distance--
	return l.moveLink(m)
}

func (l *link) updateChildrenPositions() {

	if l.child == nil {
		return
	}

	xDelta := l.position.X - l.child.position.X
	xDistance := math.Abs(float64(xDelta))

	yDelta := l.position.Y - l.child.position.Y
	yDistance := math.Abs(float64(yDelta))

	if xDistance <= 1 && yDistance <= 1 {
		return
	}

	// Head 2 steps away diagonally
	if xDistance > 1 && yDistance > 1 {
		l.child.position.X += int(xDelta / 2)
		l.child.position.Y += int(yDelta / 2)
	} else {
		// Head 2 steps away horizontally
		if xDistance > 1 {
			// Tail needs to catch up diagonally, since y-axis also has distance to cover
			if yDistance != 0 {
				// Move 1 horizontally in correct direction
				l.child.position.X += int(xDelta / 2)
				// Move 1 vertically in correct direction
				l.child.position.Y += yDelta
			}
			// Tail needs to catch up horizontally
			if yDistance == 0 {
				// Move 1 horizontally in correct direction
				l.child.position.X += int(xDelta / 2)
			}

		}
		// Head 2 steps away vertically
		if yDistance > 1 {
			// Tail needs to catch up diagonally, since x-axis also has distance to cover
			if xDistance != 0 {
				// Move 1 horizontally in correct direction
				l.child.position.X += xDelta
				// Move 1 vertically in correct direction
				l.child.position.Y += int(yDelta / 2)
			}
			// Tail needs to catch up vertically
			if xDistance == 0 {
				// Move 1 horizontally in correct direction
				l.child.position.Y += int(yDelta / 2)
			}
		}
	}

	if xDistance > 1 || yDistance > 1 {
		vpk := visitedPositionsKey(l.child.position)
		l.child.visitedPositions[vpk] = struct{}{}
	}

	l.child.updateChildrenPositions()
}

func getTailPastPositions(l *link) map[string]struct{} {
	currNode := l
	for currNode.child != nil {
		currNode = currNode.child
	}
	return currNode.visitedPositions
}

func visitedPositionsKey(p position) string {
	x := strconv.Itoa(p.X)
	y := strconv.Itoa(p.Y)
	return fmt.Sprintf("%s,%s", x, y)
}

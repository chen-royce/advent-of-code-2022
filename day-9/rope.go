package main

import (
	"fmt"
	"math"
	"strconv"
)

type rope struct {
	HeadPosition      position
	TailPosition      position
	PastTailPositions map[string]struct{}
}

type position struct {
	X int
	Y int
}

func NewRope(headPosition, tailPosition position) rope {
	return rope{
		HeadPosition: headPosition,
		TailPosition: tailPosition,
		PastTailPositions: map[string]struct{}{
			tailPositionKey(tailPosition): {},
		},
	}
}

func tailPositionKey(tailPosition position) string {
	x := strconv.Itoa(tailPosition.X)
	y := strconv.Itoa(tailPosition.Y)
	return fmt.Sprintf("%s,%s", x, y)
}

func (r *rope) UpdateTailPosition() {
	xDelta := r.HeadPosition.X - r.TailPosition.X
	xDistance := math.Abs(float64(xDelta))

	yDelta := r.HeadPosition.Y - r.TailPosition.Y
	yDistance := math.Abs(float64(yDelta))

	// Head 2 steps away horizontally
	if xDistance > 1 {
		// Tail needs to catch up diagonally, since y-axis also has distance to cover
		if yDistance != 0 {
			// Move 1 horizontally in correct direction
			r.TailPosition.X += int(xDelta / 2)
			// Move 1 vertically in correct direction
			r.TailPosition.Y += yDelta
		}
		// Tail needs to catch up horizontally
		if yDistance == 0 {
			// Move 1 horizontally in correct direction
			r.TailPosition.X += int(xDelta / 2)
		}
	}

	// Head 2 steps away vertically
	if yDistance > 1 {
		// Tail needs to catch up diagonally, since x-axis also has distance to cover
		if xDistance != 0 {
			// Move 1 horizontally in correct direction
			r.TailPosition.X += xDelta
			// Move 1 vertically in correct direction
			r.TailPosition.Y += int(yDelta / 2)
		}
		// Tail needs to catch up vertically
		if xDistance == 0 {
			// Move 1 horizontally in correct direction
			r.TailPosition.Y += int(yDelta / 2)
		}
	}

	if xDistance > 1 || yDistance > 1 {
		tpKey := tailPositionKey(r.TailPosition)
		r.PastTailPositions[tpKey] = struct{}{}
	}
}

func (r *rope) MoveRope(m Move) *rope {
	if m.Distance == 0 {
		return r
	}
	// Move head 1 unit in correct direction
	switch m.Direction {
	case "U":
		r.HeadPosition.Y++
	case "D":
		r.HeadPosition.Y--
	case "L":
		r.HeadPosition.X--
	case "R":
		r.HeadPosition.X++
	}

	// Update tail position to match
	r.UpdateTailPosition()

	// Decrement move distance and recursively continue
	// until the move is complete
	m.Distance--
	return r.MoveRope(m)
}

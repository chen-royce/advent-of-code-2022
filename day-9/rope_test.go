package main

import (
	"reflect"
	"testing"
)

func TestMoveRope(t *testing.T) {
	tests := []struct {
		name                  string
		rope                  rope
		moves                 []Move
		wantTailPosition      position
		wantPastTailPositions map[string]struct{}
	}{
		{
			name: "Move right",
			rope: NewRope(position{}, position{}),
			moves: []Move{
				{"R", 1},
				{"R", 1},
				{"R", 1},
			},
			wantTailPosition: position{2, 0},
			wantPastTailPositions: map[string]struct{}{
				"0,0": {},
				"1,0": {},
				"2,0": {},
			},
		},
		{
			name: "Move left",
			rope: NewRope(position{}, position{}),
			moves: []Move{
				{"L", 1},
				{"L", 1},
				{"L", 1},
			},
			wantTailPosition: position{-2, 0},
			wantPastTailPositions: map[string]struct{}{
				"0,0":  {},
				"-1,0": {},
				"-2,0": {},
			},
		},
		{
			name: "Move up",
			rope: NewRope(position{}, position{}),
			moves: []Move{
				{"U", 1},
				{"U", 1},
				{"U", 1},
			},
			wantTailPosition: position{0, 2},
			wantPastTailPositions: map[string]struct{}{
				"0,0": {},
				"0,1": {},
				"0,2": {},
			},
		},
		{
			name: "Move down",
			rope: NewRope(position{}, position{}),
			moves: []Move{
				{"D", 1},
				{"D", 1},
				{"D", 1},
			},
			wantTailPosition: position{0, -2},
			wantPastTailPositions: map[string]struct{}{
				"0,0":  {},
				"0,-1": {},
				"0,-2": {},
			},
		},
		{
			name: "Move diagonal - horizontal trigger",
			rope: NewRope(position{}, position{}),
			moves: []Move{
				{"R", 1},
				{"U", 1},
				{"R", 1},
			},
			wantTailPosition: position{1, 1},
			wantPastTailPositions: map[string]struct{}{
				"0,0": {},
				"1,1": {},
			},
		},
		{
			name: "Move diagonal - vertical trigger",
			rope: NewRope(position{}, position{}),
			moves: []Move{
				{"U", 1},
				{"R", 1},
				{"U", 1},
			},
			wantTailPosition: position{1, 1},
			wantPastTailPositions: map[string]struct{}{
				"0,0": {},
				"1,1": {},
			},
		},
	}
	for _, test := range tests {
		for _, move := range test.moves {
			test.rope.MoveRope(move)
		}
		if !reflect.DeepEqual(test.rope.TailPosition, test.wantTailPosition) {
			t.Fatalf("Tail position: Got %+v, want %+v", test.rope.TailPosition, test.wantTailPosition)
		}
		if !reflect.DeepEqual(test.rope.PastTailPositions, test.wantPastTailPositions) {
			t.Fatalf("Past tail positions: Got %+v, want %+v", test.rope.PastTailPositions, test.wantTailPosition)
		}
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestMoveLink(t *testing.T) {
	tests := []struct {
		name                  string
		head                  *link
		moves                 []move
		wantPastTailPositions map[string]struct{}
	}{
		{
			name: "Move right 2 times",
			head: newLink(3),
			moves: []move{
				{"R", 2},
			},
			wantPastTailPositions: map[string]struct{}{
				"0,0": {},
			},
		},
		{
			name: "Move right 4 times",
			head: newLink(3),
			moves: []move{
				{"R", 1},
				{"R", 1},
				{"R", 1},
				{"R", 1},
			},
			wantPastTailPositions: map[string]struct{}{
				"0,0": {},
				"1,0": {},
				"2,0": {},
			},
		},
		{
			name: "Move diagonally up/right",
			head: newLink(3),
			moves: []move{
				{"U", 1},
				{"R", 1},
				{"U", 1},
				{"R", 1},
				{"U", 1},
			},
			wantPastTailPositions: map[string]struct{}{
				"0,0": {},
				"1,1": {},
			},
		},
		{
			name: "Move diagonally down/left",
			head: newLink(3),
			moves: []move{
				{"D", 1},
				{"L", 1},
				{"D", 1},
				{"L", 1},
				{"D", 1},
			},
			wantPastTailPositions: map[string]struct{}{
				"0,0":   {},
				"-1,-1": {},
			},
		},
	}
	for _, test := range tests {
		for _, move := range test.moves {
			test.head.moveLink(move)
		}
		got := getTailPastPositions(test.head)
		if !reflect.DeepEqual(got, test.wantPastTailPositions) {
			t.Fatalf("Got %+v, want %+v", got, test.wantPastTailPositions)
		}
	}
}

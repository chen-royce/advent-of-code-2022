package main

import (
	"reflect"
	"testing"
)

func TestCheckNeighbors(t *testing.T) {

}

func TestBuildAdjacencyList(t *testing.T) {
	tests := []struct {
		input [][]uint8
		want  adjacencyList
	}{
		{
			input: [][]uint8{
				{97, 97},
				{97, 97},
			},
			want: adjacencyList{
				"0-0": map[string]int{
					"0-1": 1,
					"1-0": 1,
				},
				"0-1": map[string]int{
					"0-0": 1,
					"1-1": 1,
				},
				"1-0": map[string]int{
					"0-0": 1,
					"1-1": 1,
				},
				"1-1": map[string]int{
					"0-1": 1,
					"1-0": 1,
				},
			},
		},
		{
			input: [][]uint8{
				{100, 98},
				{98, 97},
			},
			want: adjacencyList{
				"0-0": map[string]int{
					"0-1": 1,
					"1-0": 1,
				},
				"0-1": map[string]int{
					"1-1": 1,
				},
				"1-0": map[string]int{
					"1-1": 1,
				},
				"1-1": map[string]int{
					"0-1": 1,
					"1-0": 1,
				},
			},
		},
	}
	for _, test := range tests {
		got := buildAdjacencyList(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("Got %+v, want %+v", got, test.want)
		}
	}
}

func TestGenerateElevationsMap(t *testing.T) {
	got := generateElevationsMap()
	want := map[uint8]int{
		S: 0, E: 25, 97: 0, 98: 1, 99: 2, 100: 3, 101: 4, 102: 5, 103: 6,
		104: 7, 105: 8, 106: 9, 107: 10, 108: 11, 109: 12, 110: 13, 111: 14,
		112: 15, 113: 16, 114: 17, 115: 18, 116: 19, 117: 20, 118: 21, 119: 22,
		120: 23, 121: 24, 122: 25,
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatal("Elevations map doesn't match expected input - check code")
	}
}

func TestCoordinatesToString(t *testing.T) {
	got := coordinatesToString(0, 0)
	want := "0-0"
	if got != want {
		t.Fatalf("Got %s, want %s", got, want)
	}
}

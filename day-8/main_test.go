package main

import (
	"reflect"
	"testing"
)

type test struct {
	forest Forest
	want   map[string]struct{}
}

var testForestA = Forest{
	[]int{3, 0, 3, 7, 3},
	[]int{2, 5, 5, 1, 2},
	[]int{6, 5, 3, 3, 2},
	[]int{3, 3, 5, 4, 9},
	[]int{3, 5, 3, 9, 0},
}

func TestLookDown(t *testing.T) {
	tests := []test{
		{
			forest: testForestA,
			want: map[string]struct{}{
				"0,0": struct{}{},
				"2,0": struct{}{},
				"0,1": struct{}{},
				"1,1": struct{}{},
				"0,2": struct{}{},
				"1,2": struct{}{},
				"0,3": struct{}{},
				"4,3": struct{}{},
				"0,4": struct{}{},
				"3,4": struct{}{},
			},
		},
	}
	for _, test := range tests {
		got := make(map[string]struct{})
		LookDown(test.forest, got)

		want := test.want

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("Got %+v\nWant %+v", got, want)
		}
	}
}

func TestLookUp(t *testing.T) {
	tests := []test{
		{
			forest: testForestA,
			want: map[string]struct{}{
				"4,0": struct{}{},
				"2,0": struct{}{},
				"4,1": struct{}{},
				"4,2": struct{}{},
				"3,2": struct{}{},
				"4,3": struct{}{},
				"4,4": struct{}{},
				"3,4": struct{}{},
			},
		},
	}
	for _, test := range tests {
		got := make(map[string]struct{})
		LookUp(test.forest, got)

		want := test.want

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("Got %+v\nWant %+v", got, want)
		}
	}
}

func TestLookLeft(t *testing.T) {
	tests := []test{
		{
			forest: testForestA,
			want: map[string]struct{}{
				"0,4": struct{}{},
				"0,3": struct{}{},
				"1,4": struct{}{},
				"1,2": struct{}{},
				"2,4": struct{}{},
				"2,3": struct{}{},
				"2,1": struct{}{},
				"2,0": struct{}{},
				"3,4": struct{}{},
				"4,3": struct{}{},
				"4,4": struct{}{},
			},
		},
	}
	for _, test := range tests {
		got := make(map[string]struct{})
		LookLeft(test.forest, got)

		want := test.want

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("Got %+v\nWant %+v", got, want)
		}
	}
}

func TestLookRight(t *testing.T) {
	tests := []test{
		{
			forest: testForestA,
			want: map[string]struct{}{
				"0,0": struct{}{},
				"0,3": struct{}{},
				"1,0": struct{}{},
				"1,1": struct{}{},
				"2,0": struct{}{},
				"3,0": struct{}{},
				"3,2": struct{}{},
				"3,4": struct{}{},
				"4,0": struct{}{},
				"4,1": struct{}{},
				"4,3": struct{}{},
			},
		},
	}
	for _, test := range tests {
		got := make(map[string]struct{})
		LookRight(test.forest, got)

		want := test.want

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("Got %+v\nWant %+v", got, want)
		}
	}
}

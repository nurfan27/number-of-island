package NumberOfIslands

import (
	"testing"
)

var TestCases = []struct {
	name string
	grid [][]int
	want int
}{
	{"TestCase1", [][]int{{1, 1, 0, 0, 0}, {0, 1, 0, 0, 1}, {1, 0, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 1, 1}}, 1},
	{"TestCase2", [][]int{{1, 1, 0, 0, 0}, {0, 1, 0, 0, 1}, {1, 0, 0, 1, 1}, {1, 0, 0, 0, 1}, {1, 0, 0, 1, 1}}, 2},
	{"TestCase3", [][]int{{1, 1, 0, 0, 0}, {0, 1, 0, 0, 1}, {1, 0, 0, 1, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 0, 1}}, 3},
	{"TestCase4", [][]int{{1, 1, 0, 0, 1}, {0, 1, 0, 0, 0}, {1, 0, 0, 1, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 0, 1}}, 4},
	{"TestCase5", [][]int{{1, 1, 0, 0, 0}, {0, 1, 0, 0, 1}, {1, 0, 0, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 1, 0, 1}}, 5},
}

func TestCountIslands(t *testing.T) {
	for _, tt := range TestCases {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNumberOfIsland().CountIslands(tt.grid)
			if got != tt.want {
				t.Errorf("CountIslands(%s) got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

package simulation

import (
	"testing"
	"ant-colony/types"
)

func TestSimulateAnts(t *testing.T) {
	antFarm := &types.AntFarm{
		Ants: 3,
		ValidPaths: [][]string{
			{"start", "room1", "end"},
			{"start", "room2", "end"},
		},
		Start: "start",
		End:   "end",
	}

	
	SimulateAnts(antFarm)
}

func TestPlengths(t *testing.T) {
	paths := [][]string{
		{"start", "room1", "end"},
		{"start", "room2", "room3", "end"},
		{"start", "end"},
	}

	lengths := plengths(paths)

	expected := []int{3, 4, 2}
	for i, length := range lengths {
		if length != expected[i] {
			t.Errorf("Expected length %d at index %d, got %d", expected[i], i, length)
		}
	}
}

func TestShortest(t *testing.T) {
	testCases := []struct {
		slice    []int
		index    int
		minValue int
	}{
		{[]int{3, 2, 4}, 1, 2},
		{[]int{5, 5, 5}, 0, 5},
		{[]int{1}, 0, 1},
	}

	for i, tc := range testCases {
		index, min := shortest(tc.slice)
		if index != tc.index || min != tc.minValue {
			t.Errorf("Case %d: Expected index %d and min %d, got index %d and min %d",
				i, tc.index, tc.minValue, index, min)
		}
	}
}
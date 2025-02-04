package simulation

import (
	"fmt"
	"strings"

	"ant-colony/types"
)

// Simulates the movement of ants along the shortest path
func SimulateAnts(antFarm *types.AntFarm, path []string) {
	ants := make([]int, antFarm.Ants)
	for i := range ants {
		ants[i] = 1 // Ants start at the first room in the path
	}

	for {
		moves := []string{}
		for i := range ants {
			if ants[i] < len(path) {
				moves = append(moves, fmt.Sprintf("L%d-%s", i+1, path[ants[i]]))
				ants[i]++
			}
		}
		if len(moves) == 0 {
			break
		}
		fmt.Println(strings.Join(moves, " "))
	}
}

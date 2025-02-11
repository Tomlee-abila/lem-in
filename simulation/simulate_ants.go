package simulation

import (
	"fmt"

	"ant-colony/types"
)

// Simulates the movement of ants along the shortest path
func SimulateAnts(antFarm *types.AntFarm) {
	antPaths := [][]string{}

	ants := antFarm.Ants
	pathLength := len(antFarm.ValidPaths)
	pathCount := pathLength

	for i := 0; i < ants; i++ {
		index := i
		for index >= pathLength {
			index -= pathLength
		}
		antPaths = append(antPaths, antFarm.ValidPaths[index])
	}
	change := true

	for change {
		change = false

		for i := range antPaths {
			if i < pathCount {
				if len(antPaths[i]) > 0 {
					change = true
					fmt.Printf("L%d - %s ", i+1, antPaths[i][0])
					antPaths = antPaths[1:]
				}
			}
		}
		pathCount += pathLength
		fmt.Println()
	}
}

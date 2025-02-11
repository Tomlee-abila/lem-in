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

	lPaths := plengths(antFarm.ValidPaths)

	for i := 0; i < ants; i++ {
		index, l := shortest(lPaths)		
		antPaths = append(antPaths, antFarm.ValidPaths[index][1:])
		if l == 2{
			continue
		}
		lPaths[index]++
	}
	// fmt.Println("antPaths\n",antPaths)
	fullRoom := make(map[string]bool)
	paths := make(map[int]string)
	change := true
	count := 0
	for change {
		count++
		change = false
		c := 0
		for i := range antPaths {
			if c < pathCount {
				if len(antPaths[i]) > 0 && !fullRoom[antPaths[i][0]]{
					change = true
					fmt.Printf("L%d-%s ", i+1, antPaths[i][0])
					fullRoom[paths[i]] = false
					if antPaths[i][0] != antFarm.End{
						fullRoom[antPaths[i][0]] = true
					}					
					paths[i] = antPaths[i][0]
					antPaths[i] = antPaths[i][1:]
					c++
				}
				
			}
			
		}
		pathCount += pathLength
		fmt.Println()
	}
}

func plengths(slices [][]string) []int {
	result := []int{}
	for _, p := range slices {
		result = append(result, len(p))
	}
	return result
}

func shortest(slice []int) (int, int) {
	min := -1
	var index int

	for i := range slice {
		if min == -1 {
			min = slice[i]
			index = i
		}
		if min > slice[i] {
			min = slice[i]
			index = i
		}
	}
	return index, min
}

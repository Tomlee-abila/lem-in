package pathfinding

import (
	"fmt"
	"strings"

	"ant-colony/types"
)

// Uses BFS to find the shortest path from start to end
func FindShortestPath(antFarm *types.AntFarm) error{
	// Parse input paths
	for _, t := range antFarm.Tunnels {
		r := strings.Split(t, "-")

		for _, room := range r {
			if _, exists := antFarm.Rooms[room]; !exists {
				antFarm.Rooms[room] = &types.Room{Name: room}
			}
		}

		antFarm.Rooms[r[0]].Links = appendIfNotExists(antFarm.Rooms[r[0]].Links, r[1])
		antFarm.Rooms[r[1]].Links = appendIfNotExists(antFarm.Rooms[r[1]].Links, r[0])
	}

	antFarm.FindPaths(antFarm.Start, []string{antFarm.Start})

	if len(antFarm.Paths) == 0 {
		return fmt.Errorf("ERROR: invalid data format, there are no paths available")
	}

	fmt.Println("All Paths:")
	for i, path := range antFarm.Paths {
		fmt.Println(i, ":", path)
	}

	antFarm.RemoveInvalidPaths()
	// fmt.Println("Valid Paths:", antFarm.ValidPaths)

	antFarm.FindOptimalPath()
	// fmt.Println("Optimal Paths:", antFarm.ValidPaths)
	return nil
}

// Append item to slice only if it doesn't exist
func appendIfNotExists(slice []string, item string) []string {
	for _, v := range slice {
		if v == item {
			return slice
		}
	}
	return append(slice, item)
}

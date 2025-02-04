package main

import (
	"fmt"
	"os"

	"ant-colony/parser"
	"ant-colony/pathfinding"
	"ant-colony/simulation"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <input_file>")
		return
	}

	filename := os.Args[1]
	antFarm, err := parser.ParseInput(filename)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	// Find the shortest path
	path := pathfinding.FindShortestPath(antFarm)
	if len(path) == 0 {
		fmt.Println("ERROR: No path found from start to end")
		return
	}

	// Simulate ant movement
	simulation.SimulateAnts(antFarm, path)
}

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
	antFarm, fileContent, err := parser.ParseInput(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find the shortest path
	err = pathfinding.FindShortestPath(antFarm)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("AntFarm\n",antFarm.ValidPaths)
	
	fmt.Println(fileContent+"\n")

	// Simulate ant movement
	simulation.SimulateAnts(antFarm)
}

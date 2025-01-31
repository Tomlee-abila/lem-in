package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <input_file>")
		return
	}

	filename := os.Args[1]
	antFarm, err := parseInput(filename)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	// Print the input file content
	fmt.Println(antFarm.Ants)
	for _, room := range antFarm.Rooms {
		fmt.Printf("%s %d %d\n", room.Name, room.X, room.Y)
	}
	for _, tunnel := range antFarm.Tunnels {
		fmt.Println(tunnel)
	}

	// Find the shortest path
	path := findShortestPath(antFarm)
	if len(path) == 0 {
		fmt.Println("ERROR: No path found from start to end")
		return
	}

	// Simulate ant movement
	simulateAnts(antFarm, path)
}

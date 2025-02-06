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

func parseInput(filename string) (*AntFarm, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	antFarm := &AntFarm{Rooms: make(map[string]*Room)}

	// Read number of ants
	if !scanner.Scan() {
		return nil, fmt.Errorf("invalid data format, no ants found")
	}
	ants, err := strconv.Atoi(scanner.Text())
	if err != nil || ants <= 0 {
		return nil, fmt.Errorf("invalid data format, invalid number of ants")
	}
	antFarm.Ants = ants

	// Read rooms and tunnels
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "##start") {
			if !scanner.Scan() {
				return nil, fmt.Errorf("invalid data format, no start room found")
			}
			startRoom := parseRoom(scanner.Text())
			antFarm.Start = startRoom.Name
			antFarm.Rooms[startRoom.Name] = startRoom
		} else if strings.HasPrefix(line, "##end") {
			if !scanner.Scan() {
				return nil, fmt.Errorf("invalid data format, no end room found")
			}
			endRoom := parseRoom(scanner.Text())
			antFarm.End = endRoom.Name
			antFarm.Rooms[endRoom.Name] = endRoom
		} else if strings.Contains(line, "-") {
			antFarm.Tunnels = append(antFarm.Tunnels, line)
		} else if len(strings.Fields(line)) == 3 {
			room := parseRoom(line)
			antFarm.Rooms[room.Name] = room
		}
	}

	// Validate start and end rooms
	if antFarm.Start == "" || antFarm.End == "" {
		return nil, fmt.Errorf("invalid data format, start or end room missing")
	}

	return antFarm, nil
}

func parseRoom(line string) *Room {
	parts := strings.Fields(line)
	x, _ := strconv.Atoi(parts[1])
	y, _ := strconv.Atoi(parts[2])
	return &Room{Name: parts[0], X: x, Y: y}
}

func findShortestPath(antFarm *AntFarm) []string {
	// Use BFS to find the shortest path
	visited := make(map[string]bool)
	queue := [][]string{{antFarm.Start}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		lastRoom := path[len(path)-1]

		if lastRoom == antFarm.End {
			return path
		}

		for _, tunnel := range antFarm.Tunnels {
			rooms := strings.Split(tunnel, "-")
			if rooms[0] == lastRoom && !visited[rooms[1]] {
				visited[rooms[1]] = true
				newPath := append(path, rooms[1])
				queue = append(queue, newPath)
			} else if rooms[1] == lastRoom && !visited[rooms[0]] {
				visited[rooms[0]] = true
				newPath := append(path, rooms[0])
				queue = append(queue, newPath)
			}
		}
	}

	return nil
}

func simulateAnts(antFarm *AntFarm, path []string) {
	// Initialize ants' positions at the start room
	ants := make([]int, antFarm.Ants)
	for i := range ants {
		ants[i] = 0 // All ants start at the first room in the path
	}

	// Create a map to track the occupancy of rooms
	roomOccupancy := make(map[string]int)
	for i := 0; i < antFarm.Ants; i++ {
		roomOccupancy[path[0]]++ // Start room can have multiple ants
	}

	turn := 0 // Initialize turn counter
	for {
		moves := []string{}
		// Track which rooms are occupied after this turn
		newOccupancy := make(map[string]int)

		for i := range ants {
			if ants[i] < len(path)-1 { // Ensure the ant can move
				nextRoom := path[ants[i]+1]
				if roomOccupancy[nextRoom] == 0 { // Check if the next room is empty
					moves = append(moves, fmt.Sprintf("L%d-%s", i+1, nextRoom))
					ants[i]++ // Move the ant to the next room
					newOccupancy[nextRoom]++ // Mark the next room as occupied
				}
			}
		}

		// Update room occupancy for the next turn
		for room, count := range roomOccupancy {
			newOccupancy[room] += count
		}
		roomOccupancy = newOccupancy

		if len(moves) == 0 {
			break // No more moves possible
		}

		// Print the moves for the current turn
		fmt.Printf("Turn %d: %s\n", turn+1, strings.Join(moves, " "))
		turn++ // Increment turn counter
	}
}

type Room struct {
	Name  string
	X, Y  int
	Links []string
}

type AntFarm struct {
	Ants    int
	Rooms   map[string]*Room
	Start   string
	End     string
	Tunnels []string
}
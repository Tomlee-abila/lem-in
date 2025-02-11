package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ant-colony/types"
)

// Parses the input file and returns the AntFarm structure
func ParseInput(filename string) (*types.AntFarm, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	antFarm := &types.AntFarm{Rooms: make(map[string]*types.Room)}

	// Read number of ants
	if !scanner.Scan() {
		return nil, fmt.Errorf("ERROR: invalid data format, no ants found")
	}
	ants, err := strconv.Atoi(scanner.Text())
	fileContent := ""
	fileContent += strconv.Itoa(ants)
	if err != nil || ants <= 1 {
		return nil, fmt.Errorf("ERROR: invalid data format, invalid number of ants")
	}
	antFarm.Ants = ants

	// Read rooms and tunnels
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Count(line, "#") != 1{
			fileContent += "\n"+line
		}

		if strings.HasPrefix(line, "##start") {
			if !scanner.Scan() {
				return nil, fmt.Errorf("ERROR: invalid data format, no start room found")
			}
			startRoom := ParseRoom(scanner.Text())
			fileContent += "\n"+startRoom.Name+" "+strconv.Itoa(startRoom.X)+" "+strconv.Itoa(startRoom.Y)
			antFarm.Start = startRoom.Name
			antFarm.Rooms[startRoom.Name] = startRoom
		} else if strings.HasPrefix(line, "##end") {
			if !scanner.Scan() {
				return nil, fmt.Errorf("ERROR: invalid data format, no end room found")
			}
			endRoom := ParseRoom(scanner.Text())
			fileContent += "\n"+endRoom.Name+" "+strconv.Itoa(endRoom.X)+" "+strconv.Itoa(endRoom.Y)
			antFarm.End = endRoom.Name
			antFarm.Rooms[endRoom.Name] = endRoom
		} else if strings.Contains(line, "-") {
			antFarm.Tunnels = append(antFarm.Tunnels, line)
		} else if len(strings.Fields(line)) == 3 {
			room := ParseRoom(line)
			antFarm.Rooms[room.Name] = room
		}
	}

	// Validate start and end rooms
	if antFarm.Start == "" || antFarm.End == "" {
		return nil, fmt.Errorf("ERROR: invalid data format, start or end room missing")
	}

	fmt.Println(fileContent)
	return antFarm, nil
}


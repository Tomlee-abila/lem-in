package types

// Room represents a room in the ant farm
type Room struct {
	Name  string
	X, Y  int
	Links []string
}

// AntFarm represents the entire ant farm
type AntFarm struct {
	Ants    int
	Rooms   map[string]*Room
	Start   string
	End     string
	Tunnels []string
	Paths      [][]string
	ValidPaths [][]string
}

// Recursively find all paths
func (antFarm *AntFarm) FindPaths(current string, path []string) {
	nPath := append([]string{}, path...)
	for _, room := range antFarm.Rooms[current].Links {
		if contains(path, room) {
			continue
		}

		newPath := append(nPath, room)
		if room == antFarm.End {
			antFarm.Paths = append(antFarm.Paths, newPath)
		} else {
			antFarm.FindPaths(room, newPath)
		}
	}
}

// Remove conflicting paths
func (antFarm *AntFarm) RemoveInvalidPaths() {
	if len(antFarm.Paths) == 0 {
		return
	}

	for len(antFarm.Paths) > 0 {
		// fmt.Println("length of paths", len(antFarm.Paths))
		i := antFarm.findShortestPath()
		shortest := append([]string{}, antFarm.Paths[i]...)
		antFarm.ValidPaths = append(antFarm.ValidPaths, shortest)
		// fmt.Println(shortest)

		toRemove := make(map[int]bool)
		toRemove[i] = true

		for k, path := range antFarm.Paths {
			if k == i {
				continue
			}
			for _, room := range shortest[1:] {
				if contains(path, room) && room != antFarm.End {
					// fmt.Println("path", k, path, "has", room)
					toRemove[k] = true
					break
				}
			}
		}

		// Rebuild Paths excluding removed ones
		var newPaths [][]string
		for k, path := range antFarm.Paths {
			if !toRemove[k] {
				newPaths = append(newPaths, path)
			}
		}
		toRemove = make(map[int]bool)
		antFarm.Paths = newPaths
		// fmt.Println("valid paths", antFarm.ValidPaths)
		// fmt.Println("To be removed", toRemove, "length of paths", len(antFarm.Paths), "\nPaths", antFarm.Paths)
	}
}

// Find the most balanced optimal paths
func (antFarm *AntFarm) FindOptimalPath() {
	pathLengths := make([]int, len(antFarm.ValidPaths))
	antsLeft := antFarm.Ants

	for i, path := range antFarm.ValidPaths {
		pathLengths[i] = len(path)
	}

	for antsLeft > 0 {
		index := findSmallestIndex(pathLengths)
		pathLengths[index]++
		antsLeft--
	}

	var newValidPaths [][]string
	for i := range antFarm.ValidPaths {
		if pathLengths[i] > len(antFarm.ValidPaths[i]) {
			newValidPaths = append(newValidPaths, antFarm.ValidPaths[i])
		}
	}
	antFarm.ValidPaths = newValidPaths
}

// Find index of the shortest path
func (antFarm *AntFarm) findShortestPath() int {
	minIdx, minLen := 0, len(antFarm.Paths[0])
	for i, path := range antFarm.Paths {
		if len(path) < minLen {
			minLen = len(path)
			minIdx = i
		}
	}
	return minIdx
}

// Find index of the smallest element in a slice
func findSmallestIndex(slice []int) int {
	minIdx, minVal := 0, slice[0]
	for i, val := range slice {
		if val < minVal {
			minIdx = i
			minVal = val
		}
	}
	return minIdx
}

// Check if a slice contains an item
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

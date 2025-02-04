package pathfinding

import (
	"strings"

	"ant-colony/types"
)

// Uses BFS to find the shortest path from start to end
func FindShortestPath(antFarm *types.AntFarm) []string {
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

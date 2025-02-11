package parser

import (
	"strconv"
	"strings"

	"ant-colony/types"
)

// Parses a room line and returns a Room object
func ParseRoom(line string) *types.Room {
	parts := strings.Fields(line)
	x, _ := strconv.Atoi(parts[1])
	y, _ := strconv.Atoi(parts[2])
	return &types.Room{Name: parts[0], X: x, Y: y}
}

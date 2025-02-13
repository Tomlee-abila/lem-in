package parser

import (
	"testing"
)

func TestParseRoom(t *testing.T) {
	testCases := []struct {
		input    string
		name     string
		x        int
		y        int
	}{
		{"room1 23 45", "room1", 23, 45},
		{"start 0 0", "start", 0, 0},
		{"end 100 100", "end", 100, 100},
	}

	for _, tc := range testCases {
		room := ParseRoom(tc.input)
		
		if room.Name != tc.name {
			t.Errorf("Expected room name %s, got %s", tc.name, room.Name)
		}
		
		if room.X != tc.x {
			t.Errorf("Expected X coordinate %d, got %d", tc.x, room.X)
		}
		
		if room.Y != tc.y {
			t.Errorf("Expected Y coordinate %d, got %d", tc.y, room.Y)
		}
	}
}
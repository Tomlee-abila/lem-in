package parser

import (
	"os"
	"testing"
)

func TestParseInput(t *testing.T) {
	// Create a temporary test file
	content := `3
##start
start 0 0
##end
end 10 10
room1 5 5
room2 7 7
start-room1
room1-room2
room2-end`

	tmpfile, err := os.CreateTemp("", "test*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Test parsing
	antFarm, _, err := ParseInput(tmpfile.Name())
	if err != nil {
		t.Fatalf("ParseInput failed: %v", err)
	}

	// Verify the parsed data
	if antFarm.Ants != 3 {
		t.Errorf("Expected 3 ants, got %d", antFarm.Ants)
	}

	if antFarm.Start != "start" {
		t.Errorf("Expected start room 'start', got %s", antFarm.Start)
	}

	if antFarm.End != "end" {
		t.Errorf("Expected end room 'end', got %s", antFarm.End)
	}

	if len(antFarm.Rooms) != 4 {
		t.Errorf("Expected 4 rooms, got %d", len(antFarm.Rooms))
	}

	if len(antFarm.Tunnels) != 3 {
		t.Errorf("Expected 3 tunnels, got %d", len(antFarm.Tunnels))
	}
}

func TestParseInputErrors(t *testing.T) {
	testCases := []struct {
		name    string
		content string
	}{
		{
			name: "invalid ants",
			content: `invalid
##start
start 0 0
##end
end 10 10`,
		},
		{
			name: "missing start",
			content: `3
##end
end 10 10`,
		},
		{
			name: "missing end",
			content: `3
##start
start 0 0`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tmpfile, err := os.CreateTemp("", "test*.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.Write([]byte(tc.content)); err != nil {
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			_, _, err = ParseInput(tmpfile.Name())
			if err == nil {
				t.Error("Expected error, got nil")
			}
		})
	}
}
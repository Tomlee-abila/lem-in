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
}

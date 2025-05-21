package game

// Position represents a grid position on the map.
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Tile represents a single map tile.
type Tile struct {
	ID       string   `json:"id"`
	Position Position `json:"position"`
	Image    string   `json:"image"`
}

// Unit represents a player or enemy unit on the map.
type Unit struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Health     int      `json:"health"`
	Position   Position `json:"position"`
	Conditions []string `json:"conditions"`
	Actions    int      `json:"actions"`
	IsEnemy    bool     `json:"isEnemy"`
}

// Trigger defines mission logic that fires based on round, location, or other conditions.
type Trigger struct {
	ID        string    `json:"id"`
	Round     int       `json:"round,omitempty"`
	Location  *Position `json:"location,omitempty"`
	Condition string    `json:"condition,omitempty"`
	Action    string    `json:"action"`
}

// Map represents the playable area for a mission.
type Map struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Tiles  []Tile `json:"tiles"`
}

// Mission aggregates all data for a single scenario.
type Mission struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Map      Map       `json:"map"`
	Units    []Unit    `json:"units"`
	Triggers []Trigger `json:"triggers"`
	Round    int       `json:"round"`
}

// State is the top-level game state served to the frontend.
type State struct {
	Mission Mission `json:"mission"`
}

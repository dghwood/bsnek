package models

// API Objects
// https://docs.battlesnake.com/api

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Extending the Coord model
func (c Coord) Add(d Coord) Coord {
	return Coord{c.X + d.X, c.Y + d.Y}
}

func (c Coord) Minus(d Coord) Coord {
	return Coord{c.X - d.X, c.Y - d.Y}
}

func (c Coord) Equal(d Coord) bool {
	return c.X == d.X && c.Y == d.Y
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func (c Coord) Dist(d Coord) int {
	return Abs(c.X-d.X) + Abs(c.Y-d.Y)
}

type Battlesnake struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Health         int            `json:"health"`
	Body           []Coord        `json:"body"`
	Head           Coord          `json:"head"`
	Length         int            `json:"length"`
	Latency        string         `json:"latency"`
	Shout          string         `json:"shout"`
	Customizations Customizations `json:"customizations"`
}

type Customizations struct {
	Color string `json:"color"`
	Head  string `json:"head"`
	Tail  string `json:"tail"`
}

type Board struct {
	Height  int           `json:"height"`
	Width   int           `json:"width"`
	Food    []Coord       `json:"food"`
	Hazards []Coord       `json:"hazards"`
	Snakes  []Battlesnake `json:"snakes"`
}

type GameState struct {
	Game  Game        `json:"game"`
	Turn  int         `json:"turn"`
	Board Board       `json:"board"`
	You   Battlesnake `json:"you"`
}

type Game struct {
	ID      string  `json:"id"`
	Ruleset Ruleset `json:"ruleset"`
	Map     string  `json:"map"`
	Source  string  `json:"source"`
	Timeout int     `json:"timeout"`
}

type Ruleset struct {
	Name     string          `json:"name"`
	Version  string          `json:"version"`
	Settings RulesetSettings `json:"settings"`
}

type RulesetSettings struct {
	FoodSpawnChance     int `json:"foodSpawnChance"`
	MinimumFood         int `json:"minimumFood"`
	HazardDamagePerTurn int `json:"hazardDamagePerTurn"`
}

// Response Objects
// https://docs.battlesnake.com/api

type BattlesnakeInfoResponse struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}

type BattlesnakeMoveResponse struct {
	Move  string `json:"move"`
	Shout string `json:"shout"`
}

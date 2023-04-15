package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")
	t := Terrain(TERRAIN_GRASS)
	fmt.Println(t)
}

type TreeModel struct {
	mesh  string
	bark  string
	leavs string
}

type Tree struct {
	model     *TreeModel
	position  float64
	height    float64
	thickness float64
	barkTint  string
	leafTint  string
}

type Terrain int

const (
	TERRAIN_GRASS Terrain = iota
	TERRAIN_HILL
	TERRAIN_RIVER
)

func (t Terrain) String() string {
	switch t {
	case TERRAIN_GRASS:
		return "TERRAIN_GRASS"
	case TERRAIN_HILL:
		return "TERRAIN_HILL"
	case TERRAIN_RIVER:
		return "TERRAIN_RIVER"
	default:
		return "Unknown"
	}
}

type World struct {
	tiles Terrain
}

func (w World) getMovementCost(x int) int {
	switch x {
	case int(TERRAIN_GRASS):
		return 1
	case int(TERRAIN_HILL):
		return 2
	case int(TERRAIN_RIVER):
		return 3
	default:
		return 1
	}
}

func (w World) isWater(x int) bool {
	switch x {
	case int(TERRAIN_GRASS):
		return false
	case int(TERRAIN_HILL):
		return false
	case int(TERRAIN_RIVER):
		return true
	default:
		return false
	}
}

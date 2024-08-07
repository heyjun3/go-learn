package revisited

import (
	"fmt"
	"math/rand"
)

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

type Texture int

const (
	GRASS_TEXTURE Texture = iota
	HILL_TEXTURE
	RIVER_TEXTURE
)

func (t Texture) String() string {
	switch t {
	case GRASS_TEXTURE:
		return "GRASS_TEXTURE"
	case HILL_TEXTURE:
		return "HILL_TEXTURE"
	case RIVER_TEXTURE:
		return "RIVER_TEXTURE"
	default:
		return "Unknown"
	}
}

type Terrain struct {
	moveCost int
	isWater  bool
	texture  Texture
}

func NewTerrain(moveCost int, isWater bool, texture Texture) Terrain {
	return Terrain{
		moveCost: moveCost,
		isWater:  isWater,
		texture:  texture,
	}
}

func (t Terrain) GetMoveCost() int {
	return t.moveCost
}

func (t Terrain) IsWater() bool {
	return t.isWater
}

func (t Terrain) GetTexture() Texture {
	return t.texture
}

type World struct {
	tiles        [][]*Terrain
	width        int
	height       int
	grassTerrain Terrain
	hillTerrain  Terrain
	riverTerrain Terrain
}

func NewWorld(width, height int) World {
	tiles := make([][]*Terrain, width)
	for i := 0; i < width; i++ {
		tiles[i] = make([]*Terrain, height)
	}
	w := World{
		tiles:        tiles,
		width:        width,
		height:       height,
		grassTerrain: NewTerrain(1, false, GRASS_TEXTURE),
		hillTerrain:  NewTerrain(3, false, HILL_TEXTURE),
		riverTerrain: NewTerrain(2, true, RIVER_TEXTURE),
	}
	w.generateTerrain()
	return w
}

func (w World) generateTerrain() {
	for x := 0; x < w.width; x++ {
		for y := 0; y < w.height; y++ {
			w.tiles[x][y] = &w.grassTerrain
		}
	}

	x := rand.Intn(w.width)
	for y := 0; y < w.height; y++ {
		w.tiles[x][y] = &w.riverTerrain
	}
}

func (w World) getTile(x, y int) *Terrain {
	return w.tiles[x][y]
}

// func (w World) getMovementCost(x int) int {
// 	switch x {
// 	case int(TERRAIN_GRASS):
// 		return 1
// 	case int(TERRAIN_HILL):
// 		return 2
// 	case int(TERRAIN_RIVER):
// 		return 3
// 	default:
// 		return 1
// 	}
// }

// func (w World) isWater(x int) bool {
// 	switch x {
// 	case int(TERRAIN_GRASS):
// 		return false
// 	case int(TERRAIN_HILL):
// 		return false
// 	case int(TERRAIN_RIVER):
// 		return true
// 	default:
// 		return false
// 	}
// }

type Entity struct{}

type Observer interface {
	OnNotify(entity *Entity, event string)
}

type Achivements struct {
	heroIsOnBridge bool
}

func (a Achivements) OnNotify(entity *Entity, event string) {
	switch event {
	case "EVENT_ENTITY_FELL":
		a.unclock("ACHEVEMENT_FELL_OFF_BRIDGE")
	default:
		break
	}
}

func (a Achivements) unclock(event string) {
	fmt.Println(event)
}

type Subject struct {
	observers []Observer
}

func (s *Subject) addObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s Subject) notify(entity *Entity, event string) {
	for _, ob := range s.observers {
		ob.OnNotify(entity, event)
	}
}

type Monster interface {
	clone() Monster
}

type Ghost struct {
	health int
	speed  int
}

func NewGhost(health, speed int) *Ghost {
	return &Ghost{
		health: health,
		speed:  speed,
	}
}

func (g Ghost) clone() Monster {
	return g
}

type Demon struct {
	Monster
}

type Sorcerer struct {
	Monster
}

// type Spawner interface {
// 	spawnMonster() *Monster
// }

// type GhostSpawner struct {}

// func (g GhostSpawner) spawnMonster() *Ghost {
// 	return &Ghost{}
// }

type Spawner struct {
	prototype Monster
}

func NewSpawner(prototype Monster) Spawner {
	return Spawner{prototype: prototype}
}

func (s Spawner) spawnMonster() Monster {
	return s.prototype.clone()
}

// type FileSystem struct {
// 	instance *FileSystem
// }

// func (f FileSystem) instance() {
// 	if (f.instance == nil) {
// 		f.instance = &FileSystem{}
// 	}
// }

type Bullet struct {
	x int
	y int
}

func NewBullet(x, y int) *Bullet {
	return &Bullet{
		x: x,
		y: y,
	}
}

func (b Bullet) isOnScreen() bool {
	return b.x >= 0 && b.y >= 0
}

func (b *Bullet) move() {
	b.x += 5
}

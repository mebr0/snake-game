package game

import (
	"math/rand"
	"time"

	"github.com/mebr0/snake-game/pkg/domain"
)

// BaseSnake describe methods for snake in game
type BaseSnake interface {
	Head() domain.Point
	Body() []domain.Point
	Grow(p domain.Point)
	Move(dx int, dy int) bool
}

type Game struct {
	Height int
	Width  int
	Round  int
	Score  int
	Snake  BaseSnake
	Food   *domain.Point
}

func NewGame(height int, width int) Game {
	g := Game{
		Height: height,
		Width:  width,
		Round:  0,
		Score:  0,
		Snake:  newSnake(height/2, width/2),
	}

	g.generateFood()

	return g
}

func (g *Game) generateFood() {
	found := false

	for !found {
		rand.Seed(time.Now().UnixNano())
		// calibrate x and y within bounds
		x := rand.Intn(g.Height-2) + 1
		y := rand.Intn(g.Width-2) + 1

		again := false

		// check collisions with snake
		for _, p := range g.Snake.Body() {
			if p.X == x && p.Y == y {
				again = true
				break
			}
		}

		if again {
			continue
		}

		found = true

		g.Food = &domain.Point{
			X: x,
			Y: y,
		}
	}
}

type Move int

const (
	LeftMove Move = iota
	UpMove
	RightMove
	DownMove
)

func (m Move) Delta() (int, int) {
	switch m {
	case LeftMove:
		return 0, -1
	case UpMove:
		return -1, 0
	case RightMove:
		return 0, 1
	case DownMove:
		return 1, 0
	}

	return 0, 0
}

func (g *Game) Move(direction Move) error {
	g.Round += 1

	dx, dy := direction.Delta()
	x, y := g.Snake.Head().X+dx, g.Snake.Head().Y+dy

	// check collisions with snake
	if g.Snake.Move(dx, dy) {
		return ErrCollisionWithSnake
	}

	// check collisions with bounds
	if x == 0 || x == g.Height-1 || y == 0 || y == g.Width-1 {
		return ErrCollisionWithBounds
	}

	// check collisions with food
	if x == g.Food.X && y == g.Food.Y {
		g.Snake.Grow(domain.Point{
			X: x,
			Y: y,
		})

		g.Score += 1
		g.generateFood()
	}

	return nil
}

type Entity int

const (
	Space Entity = iota
	Bound
	Snake
	SnakeHead
	Food
)

func (g *Game) GetEntity(x int, y int) Entity {
	if x == g.Food.X && y == g.Food.Y {
		return Food
	}

	if x == 0 || x == g.Height-1 || y == 0 || y == g.Width-1 {
		return Bound
	}

	if x == g.Snake.Head().X && y == g.Snake.Head().Y {
		return SnakeHead
	}

	for _, p := range g.Snake.Body() {
		if x == p.X && y == p.Y {
			return Snake
		}
	}

	return Space
}

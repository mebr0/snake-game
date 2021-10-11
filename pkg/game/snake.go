package game

import "github.com/mebr0/snake-game/pkg/domain"

type snake struct {
	body []domain.Point
}

func newSnake(x int, y int) *snake {
	return &snake{
		body: []domain.Point{
			{
				X: x,
				Y: y,
			},
		},
	}
}

func (s *snake) Head() domain.Point {
	return s.body[0]
}

func (s *snake) Body() []domain.Point {
	return s.body
}

func (s *snake) Grow(p domain.Point) {
	s.body = append(s.body, p)
}

func (s *snake) Move(dx int, dy int) bool {
	x, y := s.body[0].X+dx, s.body[0].Y+dy
	col := false

	// check collisions with snake
	for _, p := range s.body {
		if p.X == x && p.Y == y {
			col = true
			break
		}
	}

	// move snake
	for i := len(s.body) - 1; i > 0; i -= 1 {
		s.body[i] = s.body[i-1]
	}

	s.body[0] = domain.Point{
		X: x,
		Y: y,
	}

	return col
}

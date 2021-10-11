package game

import (
	"testing"

	"github.com/mebr0/snake-game/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewSnake(t *testing.T) {
	s := newSnake(0, 0)

	assert.Equal(t, 1, len(s.body))
}

func TestSnake_Grow(t *testing.T) {
	s := newSnake(0, 0)

	s.Grow(domain.Point{
		X: 0,
		Y: 1,
	})

	assert.Equal(t, 2, len(s.body))
}

func TestSnake_Move(t *testing.T) {
	tt := []struct {
		name    string
		body    []domain.Point
		dx      int
		dy      int
		result  bool
		resBody []domain.Point
	}{
		{
			name: "ok",
			body: []domain.Point{
				{
					X: 0,
					Y: 0,
				},
				{
					X: -1,
					Y: 0,
				},
			},
			dx:     1,
			dy:     0,
			result: false,
			resBody: []domain.Point{
				{
					X: 1,
					Y: 0,
				},
				{
					X: 0,
					Y: 0,
				},
			},
		},
		{
			name: "collided",
			body: []domain.Point{
				{
					X: 0,
					Y: 0,
				},
				{
					X: -1,
					Y: 0,
				},
			},
			dx:     -1,
			dy:     0,
			result: true,
			resBody: []domain.Point{
				{
					X: -1,
					Y: 0,
				},
				{
					X: 0,
					Y: 0,
				},
			},
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			s := snake{
				body: test.body,
			}

			res := s.Move(test.dx, test.dy)

			assert.Equal(t, test.result, res)
			assert.Equal(t, test.resBody, s.body)
		})
	}
}

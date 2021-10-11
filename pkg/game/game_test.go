package game

import (
	"testing"

	"github.com/mebr0/snake-game/pkg/domain"
	"github.com/mebr0/snake-game/pkg/game/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewGame(t *testing.T) {
	g := NewGame(6, 9)

	assert.Equal(t, 6, g.Height)
	assert.Equal(t, 9, g.Width)
	assert.Equal(t, 0, g.Round)
	assert.Equal(t, 0, g.Score)

	s := g.Snake

	assert.Equal(t, 1, len(s.Body()))
	assert.Equal(t, 3, s.Head().X)
	assert.Equal(t, 4, s.Head().Y)
}

func TestMove_Delta(t *testing.T) {
	tt := []struct {
		name string
		move Move
		x    int
		y    int
	}{
		{
			name: "left",
			move: LeftMove,
			x:    0,
			y:    -1,
		},
		{
			name: "up",
			move: UpMove,
			x:    -1,
			y:    0,
		},
		{
			name: "right",
			move: RightMove,
			x:    0,
			y:    1,
		},
		{
			name: "down",
			move: DownMove,
			x:    1,
			y:    0,
		},
		{
			name: "error",
			move: Move(-1),
			x:    0,
			y:    0,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			x, y := test.move.Delta()

			assert.Equal(t, test.x, x)
			assert.Equal(t, test.y, y)
		})
	}
}

func mockGame(height int, width int) (Game, *mocks.BaseSnake) {
	sMock := &mocks.BaseSnake{}

	g := Game{
		Height: height,
		Width:  width,
		Round:  0,
		Score:  0,
		Snake:  sMock,
	}

	sMock.On("Body").Return([]domain.Point{
		{
			X: 0,
			Y: 0,
		},
	})

	g.generateFood()

	return g, sMock
}

func TestGame_Move(t *testing.T) {
	type mockBehaviour func(s *mocks.BaseSnake)

	tt := []struct {
		name          string
		mockBehaviour mockBehaviour
		move          Move
		eatFood       bool
		success       bool
		err           error
	}{
		{
			name: "ok",
			mockBehaviour: func(s *mocks.BaseSnake) {
				s.On("Head").Return(domain.Point{
					X: 3,
					Y: 4,
				})
				s.On("Move", mock.Anything, mock.Anything).Return(false)
			},
			move:    UpMove,
			success: true,
		},
		{
			name: "collision with snake",
			mockBehaviour: func(s *mocks.BaseSnake) {
				s.On("Head").Return(domain.Point{
					X: 3,
					Y: 4,
				})
				s.On("Move", mock.Anything, mock.Anything).Return(true)
			},
			move:    UpMove,
			success: false,
			err:     ErrCollisionWithSnake,
		},
		{
			name: "collision with bounds",
			mockBehaviour: func(s *mocks.BaseSnake) {
				s.On("Head").Return(domain.Point{
					X: 1,
					Y: 0,
				})
				s.On("Move", mock.Anything, mock.Anything).Return(false)
			},
			move:    UpMove,
			success: false,
			err:     ErrCollisionWithBounds,
		},
		{
			name: "collision with food",
			mockBehaviour: func(s *mocks.BaseSnake) {
				s.On("Head").Return(domain.Point{
					X: 3,
					Y: 4,
				})
				s.On("Move", mock.Anything, mock.Anything).Return(false)
				s.On("Grow", mock.Anything)
			},
			move:    UpMove,
			eatFood: true,
			success: true,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			g, sMock := mockGame(6, 9)

			test.mockBehaviour(sMock)

			if test.eatFood {
				dx, dy := test.move.Delta()

				g.Food = &domain.Point{
					X: g.Snake.Head().X + dx,
					Y: g.Snake.Head().Y + dy,
				}
			}

			err := g.Move(test.move)

			if test.success {
				require.NoError(t, err)
			} else {
				require.Error(t, test.err, err)
			}
		})
	}
}

func TestGame_GetEntity(t *testing.T) {
	type mockBehaviour func(s *mocks.BaseSnake)

	tt := []struct {
		name          string
		x             int
		y             int
		mockBehaviour mockBehaviour
		food          bool
		entity        Entity
	}{
		{
			name:          "food",
			x:             3,
			y:             4,
			mockBehaviour: func(s *mocks.BaseSnake) {},
			food:          true,
			entity:        Food,
		},
		{
			name:          "bound",
			x:             0,
			y:             0,
			mockBehaviour: func(s *mocks.BaseSnake) {},
			food:          false,
			entity:        Bound,
		},
		{
			name: "head",
			x:    3,
			y:    4,
			mockBehaviour: func(s *mocks.BaseSnake) {
				s.On("Head").Return(domain.Point{
					X: 3,
					Y: 4,
				})
			},
			food:   false,
			entity: SnakeHead,
		},
		// Todo: body
		{
			name: "space",
			x:    3,
			y:    4,
			mockBehaviour: func(s *mocks.BaseSnake) {
				s.On("Head").Return(domain.Point{
					X: 2,
					Y: 4,
				})
				// not work
				s.On("Body").Return([]domain.Point{
					{
						X: 2,
						Y: 4,
					},
					{
						X: 1,
						Y: 4,
					},
				})
			},
			food:   false,
			entity: Space,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			g, sMock := mockGame(6, 9)

			test.mockBehaviour(sMock)

			if test.food {
				g.Food = &domain.Point{
					X: test.x,
					Y: test.y,
				}
			}

			entity := g.GetEntity(test.x, test.y)

			assert.Equal(t, test.entity, entity)
		})
	}
}

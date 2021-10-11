package game

import "errors"

var (
	ErrCollisionWithSnake  = errors.New("collided with snake body")
	ErrCollisionWithBounds = errors.New("collided with bounds")
)

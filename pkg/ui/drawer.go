package ui

import (
	"fmt"

	"github.com/mebr0/snake-game/pkg/game"
)

var chars = map[game.Entity]string{
	game.Space:     " ",
	game.Snake:     "o",
	game.SnakeHead: "q",
	game.Food:      "@",
	game.Bound:     "#",
}

// Draw current state of Game
func Draw(g game.Game) {
	fmt.Printf("Height: %d | Width: %d | Round: %d | Score: %d\n", g.Height, g.Width, g.Round, g.Score)

	for i := 0; i < g.Height; i += 1 {
		for j := 0; j < g.Width; j += 1 {
			e := g.GetEntity(i, j)

			fmt.Print(chars[e])
		}

		fmt.Print("\n")
	}
}

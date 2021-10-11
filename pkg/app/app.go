package app

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mebr0/snake-game/pkg/game"
	"github.com/mebr0/snake-game/pkg/ui"
)

// Run application with height and width of frame
func Run(height int, width int) {
	g := game.NewGame(height, width)

	ui.Draw(g)

	scanner := bufio.NewScanner(os.Stdin)
	alive := true

	for alive && scanner.Scan() {
		com := scanner.Text()

		dir := read(com)

		if err := g.Move(dir); err != nil {
			fmt.Println(err.Error())
			alive = false
		}

		if alive {
			ui.Draw(g)
		}
	}

	fmt.Println("game finished")
}

func read(command string) game.Move {
	switch command {
	case "l":
		return game.LeftMove
	case "u":
		return game.UpMove
	case "r":
		return game.RightMove
	case "d":
		return game.DownMove
	default:
		return game.LeftMove
	}
}

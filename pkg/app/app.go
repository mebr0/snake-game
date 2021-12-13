package app

import (
	"errors"
	"fmt"

	"github.com/mebr0/snake-game/pkg/game"
	"github.com/mebr0/snake-game/pkg/ui"
)

// Run application with height and width of frame
func Run(height int, width int) {
	g := game.NewGame(height, width)

	ui.Draw(g)

	alive := true

	var com string

	for alive {
		_, err := fmt.Scanln(&com)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		dir, err := read(com)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		if err = g.Move(dir); err != nil {
			fmt.Println(err.Error())
			alive = false
			continue
		}

		if alive {
			ui.Draw(g)
		}
	}

	fmt.Println("game finished")
}

var errInvalidCommand = errors.New("invalid command")

func read(command string) (game.Move, error) {
	switch command {
	case "l":
		return game.LeftMove, nil
	case "u":
		return game.UpMove, nil
	case "r":
		return game.RightMove, nil
	case "d":
		return game.DownMove, nil
	}

	return 0, errInvalidCommand
}

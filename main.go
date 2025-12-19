package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal()
	}
	defer screen.Fini()

	err = screen.Init()
	if err != nil {
		log.Fatal()
	}

	// Game init section
	player := NewSprite('@', 10, 10)

	coins := []*Sprite {
		NewSprite('0', 12, 4),
		NewSprite('0', 20, 3),
		NewSprite('0', 6, 10),
	}

	// Game Loop
	running := true
	for running {
		//? Draw logic
		screen.Clear()

		player.Draw(screen)

		for _, coin := range coins {
			coin.Draw(screen)
		}

		screen.Show()
		//? Update logic

		// Getting the event
		ev := screen.PollEvent()
		// Checking the event type
		switch ev := ev.(type) {
		case *tcell.EventKey:
			// Checking the event key
			switch ev.Rune() {
			//? Movement/keybinds cases
			case 'q':
				running = false
			case 'w':
				player.Y -= 1
			case 'a':
				player.X -= 1
			case 's':
				player.Y += 1
			case 'd':
				player.X += 1
			}
		}
	}
}
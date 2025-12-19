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

	// Game Loop
	running := true
	for running {
		// Draw logic
		screen.Clear()

		player.Draw(screen)

		screen.Show()
		// Update logic

		// Getting the event
		ev := screen.PollEvent()
		// Checking the event type
		switch ev := ev.(type) {
		case *tcell.EventKey:
			// Checking the event key
			switch ev.Rune() {
			case 'q':
				running = false
			}
		}
	}
}
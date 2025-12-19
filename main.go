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

	// Game Loop
	running := true
	for running {
		// Draw logic
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
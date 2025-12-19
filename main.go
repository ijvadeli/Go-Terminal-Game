package main

import (
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/gdamore/tcell/v2"
)

// Draw scores
func drawString(screen tcell.Screen, x, y int, msg string) {
	for index, char := range msg {
		screen.SetContent(x+index, y, char, nil, tcell.StyleDefault)
	}
}

// Random coin spawn + level function
func setupCoins(level int) []*Sprite {
	coins := make([]*Sprite, level+2)
	for index := range level + 2 {
		coins[index] = NewSprite(
			'0',
			rand.IntN(20),
			rand.IntN(20),
		)
	}
	return coins
}

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

	coins := setupCoins(1)

	score := 0
	level := 1

	// Game Loop
	running := true
	for running {
		// Draw logic

		screen.Clear()

		player.Draw(screen)

		for _, coin := range coins {
			coin.Draw(screen)
		}

		//? UI
		// Scoreboard
		drawString(
			screen,
			1,
			1,
			fmt.Sprintf("Score: %d", score),
		)
		// Level
		drawString(
			screen,
			1,
			2,
			fmt.Sprintf("Level: %d", level),
		)

		screen.Show()

		// Update logic

		playerMoved := false

		// Getting the event
		ev := screen.PollEvent()
		// Checking the event type
		switch ev := ev.(type) {
		case *tcell.EventKey:
			// Checking the event key
			switch ev.Rune() {
			// Movement/keybinds cases
			case 'q':
				running = false
			case 'w':
				player.Y -= 1
				playerMoved = true
			case 'a':
				player.X -= 1
				playerMoved = true
			case 's':
				player.Y += 1
				playerMoved = true
			case 'd':
				player.X += 1
				playerMoved = true
			}
		}

		// Check for coin collisions
		if playerMoved {
			coinCollectedIndex := -1
			for index, coin := range coins {
				if coin.X == player.X && coin.Y == player.Y {
					// Collect the coin
					coinCollectedIndex = index
					score++
				}
			}

			// Handle coin collision
			if coinCollectedIndex > -1 {
				// Swap target with last
				coins[coinCollectedIndex] = coins[len(coins)-1]
				// Trim off last item
				coins = coins[0: len(coins)-1]

				if len(coins) == 0 {
					level++
					coins = setupCoins(level)
				}
			}
		}
	}
}
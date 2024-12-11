package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MIN_NUMBER, MAX_NUMBER = 1, 100

type Game struct {
	attemps       int
	random_number int
	difficulty    string
}

func game(settings Game) {
	fmt.Printf("%d %d %s \n", settings.attemps, settings.random_number, settings.difficulty)
}

func main() {
	rand.New(rand.NewSource((time.Now().UnixNano())))
	var difficulty string
	randonNumber := rand.Intn(MAX_NUMBER)

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Printf("I'm thinking of a number between 1 and %d. \n", randonNumber)
	fmt.Printf("Please select the difficulty level: \n1. Easy (10 chances) \n2. Medium (5 chances) \n3. Hard (3 chances)\n")
	fmt.Scanln(&difficulty)

	var settings Game
	switch difficulty {
	case "1":
		settings = Game{
			attemps:       10,
			random_number: randonNumber,
			difficulty:    "Easy",
		}
	case "2":
		settings = Game{
			attemps:       5,
			random_number: randonNumber,
			difficulty:    "Medium",
		}
	case "3":
		settings = Game{
			attemps:       3,
			random_number: randonNumber,
			difficulty:    "Hard",
		}
	default:
		fmt.Println("Invalid difficulty level. Please select 1, 2, or 3.")
		return
	}

	game(settings)
}

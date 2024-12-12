package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const MIN_NUMBER, MAX_NUMBER = 1, 100

type Game struct {
	attemps       int
	random_number int
	difficulty    string
}

func gameLoop(settings Game) {
	fmt.Printf("\nGreat, you have selected the %s difficulty level. \nlet's start the game\n\n", settings.difficulty)

	for attemp := 1; attemp <= settings.attemps; attemp++ {
		fmt.Printf("Enter your guess: ")

		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')

		var guess int
		fmt.Scanf("%d", &guess)

		if guess == settings.random_number {
			fmt.Printf("Congratulations, you  guess the number in %d attemps\n\n", attemp)
			return
		} else if guess > settings.random_number {
			fmt.Printf("Incorrect! The number is less than %d \n\n", guess)
		} else if guess < settings.random_number {
			fmt.Printf("Incorrect! The number is great than %d \n\n", guess)
		}

	}
}

func main() {
	var settings Game
	rand.NewSource((time.Now().UnixNano()))
	randomNumber := rand.Intn(MAX_NUMBER)

	fmt.Printf("Welcome to the Number Guessing Game!\n")
	fmt.Printf("I'm thinking of a number between 1 and %d. \n\n", MAX_NUMBER)
	fmt.Printf("Please select the difficulty level: \n1. Easy (10 chances) \n2. Medium (5 chances) \n3. Hard (3 chances)\n\n")

	var difficulty int
	fmt.Printf("Enter your choice: ")
	fmt.Scanf("%d", &difficulty)

	switch difficulty {
	case 1:
		settings = Game{
			attemps:       10,
			random_number: randomNumber,
			difficulty:    "Easy",
		}
	case 2:
		settings = Game{
			attemps:       5,
			random_number: randomNumber,
			difficulty:    "Medium",
		}
	case 3:
		settings = Game{
			attemps:       3,
			random_number: randomNumber,
			difficulty:    "Hard",
		}
	default:
		fmt.Println("Invalid difficulty level. Please select 1, 2, or 3.")
		return
	}

	gameLoop(settings)
}

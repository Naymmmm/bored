package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	secretNumber := rand.Intn(100) + 1

	fmt.Println("🥱 Bored aren't you?")
	fmt.Println("👋 Welcome to Guess the Number!")
	fmt.Println("🎲 I have chosen a number between 1 and 100.")
	fmt.Println("🤔 Can you guess the number?")

	// Game loop
	for {
		fmt.Print("⌨️ Enter your guess: ")
		var guess int
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("❌ Invalid input. Please enter a valid number.")
			continue
		}

		if guess < secretNumber {
			fmt.Println("🪫❌ Too low! Try again.")
		} else if guess > secretNumber {
			fmt.Println("⚡❌ Too high! Try again.")
		} else {
			fmt.Println("✅😄 Congratulations! You guessed the number!")
			break
		}
	}

	fmt.Println("🙏 Thank you for playing Guess the Number!")
}

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	Rock     = "rock"
	Paper    = "paper"
	Scissors = "scissors"
)

var (
	choices = []string{Rock, Paper, Scissors}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("👋 Welcome to Rock, Paper, Scissors!")
	fmt.Println("🎮 Let's play!")

	// Game loop
	for {
		// Get player's choice
		playerChoice := getPlayerChoice()

		// Generate computer's choice
		computerChoice := getRandomChoice()

		fmt.Printf("Player🧑: %s\n", playerChoice)
		fmt.Printf("Computer🖥️: %s\n", computerChoice)

		// Determine the winner
		winner := determineWinner(playerChoice, computerChoice)
		if winner == "" {
			fmt.Println("👔 It's a tie!")
		} else {
			fmt.Printf("🏆 The winner is: %s\n", winner)
		}

		// Ask to play again
		playAgain := askToPlayAgain()
		if !playAgain {
			break
		}
	}

	fmt.Println("🙏 Thank you for playing Rock, Paper, Scissors!")
}

func getPlayerChoice() string {
	var choice string

	for {
		fmt.Print("⌨️ Enter your choice (rock, paper, scissors): ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("❌ Invalid input. Please try again.")
			continue
		}

		choice = normalizeChoice(choice)
		if isValidChoice(choice) {
			break
		}

		fmt.Println("❌ Invalid choice. Please try again.")
	}

	return choice
}

func normalizeChoice(choice string) string {
	return strings.ToLower(strings.TrimSpace(choice))
}

func isValidChoice(choice string) bool {
	for _, c := range choices {
		if c == choice {
			return true
		}
	}
	return false
}

func getRandomChoice() string {
	index := rand.Intn(len(choices))
	return choices[index]
}

func determineWinner(playerChoice, computerChoice string) string {
	if playerChoice == computerChoice {
		return ""
	}

	switch playerChoice {
	case Rock:
		if computerChoice == Scissors {
			return "Player"
		}
	case Paper:
		if computerChoice == Rock {
			return "Player"
		}
	case Scissors:
		if computerChoice == Paper {
			return "Player"
		}
	}

	return "Computer"
}

func askToPlayAgain() bool {
	var playAgain string

	for {
		fmt.Print("▶️ Play again? (y/n): ")
		_, err := fmt.Scanln(&playAgain)
		if err != nil {
			fmt.Println("❌ Invalid input. Please try again.")
			continue
		}

		playAgain = strings.ToLower(strings.TrimSpace(playAgain))
		if playAgain == "y" {
			return true
		} else if playAgain == "n" {
			return false
		}

		fmt.Println("❌ Invalid choice. Please try again.")
	}
}

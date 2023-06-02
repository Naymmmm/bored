package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	words = []string{
		"apple",
		"banana",
		"orange",
		"grape",
		"pineapple",
		"watermelon",
		"strawberry",
	}

	maxAttempts = 6
)

func main() {
	rand.Seed(time.Now().UnixNano())

	word := getRandomWord()
	hiddenWord := strings.Repeat("_", len(word))
	guesses := ""
	attempts := 0

	fmt.Println("🪂 Welcome to Hangman!")
	fmt.Println("🎲 Try to guess the word.")

	for {
		fmt.Println()
		fmt.Println("⚠️ Attempts remaining:", maxAttempts-attempts)
		fmt.Println("✅ Guessed letters:", guesses)
		fmt.Println("❔ Word:", hiddenWord)

		fmt.Print("⌨️ Enter a letter: ")
		reader := bufio.NewReader(os.Stdin)
		guess, _ := reader.ReadString('\n')
		guess = strings.TrimSpace(strings.ToLower(guess))

		if len(guess) != 1 {
			fmt.Println("❌ Invalid input. Please enter a single letter.")
			continue
		}

		if strings.Contains(guesses, guess) {
			fmt.Println("❌ You already guessed that letter. Try again.")
			continue
		}

		guesses += guess

		if strings.Contains(word, guess) {
			hiddenWord = revealLetters(word, hiddenWord, guess)
		} else {
			attempts++
			fmt.Println("❌ Incorrect guess.")
		}

		if hiddenWord == word {
			fmt.Println("🏆 Congratulations! You guessed the word correctly!")
			break
		}

		if attempts >= maxAttempts {
			fmt.Println("❌ Game over! You ran out of attempts.")
			fmt.Println("💭 The word was:", word)
			break
		}
	}

	fmt.Println("🙏 Thank you for playing Hangman!")
}

func getRandomWord() string {
	index := rand.Intn(len(words))
	return words[index]
}

func revealLetters(word, hiddenWord, guess string) string {
	var revealed strings.Builder

	for i := 0; i < len(word); i++ {
		if word[i] == guess[0] {
			revealed.WriteByte(word[i])
		} else {
			revealed.WriteByte(hiddenWord[i])
		}
	}

	return revealed.String()
}

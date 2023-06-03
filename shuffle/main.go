package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Print("⌨️ Enter a string to shuffle: ")
	var input string
	fmt.Scanln(&input)

	shuffledString := shuffleString(input)

	fmt.Printf("⌨️ Original String: %s\n", input)
	fmt.Printf("⌨️ Shuffled String: %s\n", shuffledString)
}

func shuffleString(input string) string {
	rand.Seed(time.Now().UnixNano())
	chars := strings.Split(input, "")
	rand.Shuffle(len(chars), func(i, j int) {
		chars[i], chars[j] = chars[j], chars[i]
	})
	return strings.Join(chars, "")
}

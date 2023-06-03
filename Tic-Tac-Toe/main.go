package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

var (
	board   [3][3]string
	player  string
	moves   int
	winner  string
	symbols = map[string]string{
		"X": "X",
		"O": "O",
	}
)

func main() {
	clearConsole()

	fmt.Println("👋 Welcome to Tic-Tac-Toe!")
	fmt.Println("❌ Player 1: X")
	fmt.Println("⭕ Player 2: O")

	initializeBoard()
	player = "X"

	for !isGameOver() {
		displayBoard()
		performMove()
		switchPlayer()
	}

	displayBoard()
	printResult()

	fmt.Println("👋 Thank you for playing Tic-Tac-Toe!")
}

func clearConsole() {
	cmd := exec.Command("clear") // Use "cls" instead of "clear" for Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func initializeBoard() {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = " "
		}
	}
}

func displayBoard() {
	clearConsole()
	fmt.Println("     1   2   3")
	fmt.Println("   -------------")
	for i := 0; i < len(board); i++ {
		fmt.Printf("%d |", i+1)
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf(" %s |", board[i][j])
		}
		fmt.Println()
		fmt.Println("   -------------")
	}
}

func performMove() {
	var row, col int

	for {
		fmt.Printf("⌨️ Player %s's turn. Enter row (1-3): ", player)
		rowInput := readInput()
		row, _ = strconv.Atoi(rowInput)

		fmt.Printf("⌨️ Player %s's turn. Enter column (1-3): ", player)
		colInput := readInput()
		col, _ = strconv.Atoi(colInput)

		if isValidMove(row, col) {
			break
		}

		fmt.Println("❌ Invalid move. Please try again.")
	}

	board[row-1][col-1] = symbols[player]
	moves++
}

func readInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func isValidMove(row, col int) bool {
	if row < 1 || row > 3 || col < 1 || col > 3 {
		return false
	}

	if board[row-1][col-1] != " " {
		return false
	}

	return true
}

func switchPlayer() {
	if player == "X" {
		player = "O"
	} else {
		player = "X"
	}
}

func isGameOver() bool {
	if moves < 5 {
		return false
	}

	if checkRows() || checkColumns() || checkDiagonals() {
		return true
	}

	if moves == 9 {
		winner = "draw"
		return true
	}

	return false
}

func checkRows() bool {
	for i := 0; i < len(board); i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != " " {
			winner = board[i][0]
			return true
		}
	}
	return false
}

func checkColumns() bool {
	for j := 0; j < len(board[0]); j++ {
		if board[0][j] == board[1][j] && board[1][j] == board[2][j] && board[0][j] != " " {
			winner = board[0][j]
			return true
		}
	}
	return false
}

func checkDiagonals() bool {
	if (board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != " ") ||
		(board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] != " ") {
		winner = board[1][1]
		return true
	}
	return false
}

func printResult() {
	if winner == "draw" {
		fmt.Println("📌 It's a draw!")
	} else {
		fmt.Printf("🥇 Player %s wins!\n", winner)
	}
}

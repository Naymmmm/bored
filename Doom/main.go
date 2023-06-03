package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	MapSize     = 10
	PlayerChar  = 'P'
	WallChar    = '#'
	ExitChar    = 'X'
	EmptyChar   = '.'
	GunChar     = 'G'
	BulletChar  = '*'
	PlayerAmmo  = 10
	MaxBullets  = 5
	BulletRange = 3
)

var (
	gameMap     [MapSize][MapSize]rune
	playerX     int
	playerY     int
	gameActive  bool
	playerAmmo  int
	bulletCount int
)

func main() {
	initGame()
	render()

	var input string
	for gameActive {
		fmt.Print("Enter a move: ")
		fmt.Scanln(&input)

		switch strings.ToLower(input) {
		case "w":
			movePlayer(playerX, playerY-1)
		case "a":
			movePlayer(playerX-1, playerY)
		case "s":
			movePlayer(playerX, playerY+1)
		case "d":
			movePlayer(playerX+1, playerY)
		case "q":
			fmt.Println("Thanks for playing!")
			return
		case "f":
			fireGun()
		default:
			fmt.Println("Invalid input. Use 'W', 'A', 'S', 'D' to move, 'F' to fire, or 'Q' to quit.")
		}

		render()
	}

	fmt.Println("Game over!")
}

func initGame() {
	gameActive = true
	playerX = 1
	playerY = 1
	playerAmmo = PlayerAmmo
	bulletCount = MaxBullets

	for i := 0; i < MapSize; i++ {
		for j := 0; j < MapSize; j++ {
			if i == 0 || i == MapSize-1 || j == 0 || j == MapSize-1 {
				gameMap[i][j] = WallChar
			} else {
				gameMap[i][j] = EmptyChar
			}
		}
	}

	gameMap[MapSize-2][MapSize-2] = ExitChar
	gameMap[playerY][playerX] = PlayerChar
	gameMap[MapSize/2][MapSize/2] = GunChar
}

func render() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for i := 0; i < MapSize; i++ {
		for j := 0; j < MapSize; j++ {
			fmt.Print(string(gameMap[i][j]), " ")
		}
		fmt.Println()
	}

	fmt.Printf("Ammo: %d\n", playerAmmo)
	fmt.Printf("Bullets: %d/%d\n", bulletCount, MaxBullets)
}

func movePlayer(x, y int) {
	if x >= 0 && x < MapSize && y >= 0 && y < MapSize {
		switch gameMap[y][x] {
		case WallChar:
			fmt.Println("You bumped into a wall!")
		case ExitChar:
			fmt.Println("Congratulations, you reached the exit!")
			gameActive = false
		case GunChar:
			gameMap[y][x] = EmptyChar
			playerAmmo += PlayerAmmo
			fmt.Println("You found a gun and gained additional ammo!")
		case EmptyChar:
			gameMap[playerY][playerX] = EmptyChar
			playerX = x
			playerY = y
			gameMap[playerY][playerX] = PlayerChar
		}
	}
}

func fireGun() {
	if bulletCount > 0 {
		dx, dy := getDirection()
		for i := 0; i < BulletRange; i++ {
			x := playerX + dx*i
			y := playerY + dy*i
			if x >= 0 && x < MapSize && y >= 0 && y < MapSize {
				switch gameMap[y][x] {
				case WallChar:
					gameMap[y][x] = EmptyChar
					fmt.Println("You destroyed a wall!")
				case PlayerChar:
					fmt.Println("You cannot shoot yourself!")
				case ExitChar:
					fmt.Println("You cannot shoot the exit!")
				case GunChar:
					fmt.Println("You missed the gun!")
				case EmptyChar:
					gameMap[y][x] = BulletChar
				}
			}
		}
		bulletCount--
		playerAmmo--
	} else {
		fmt.Println("Out of bullets!")
	}
}

func getDirection() (int, int) {
	var dx, dy int

	fmt.Print("Enter a direction (W, A, S, D): ")
	var input string
	fmt.Scanln(&input)

	switch strings.ToLower(input) {
	case "w":
		dy = -1
	case "a":
		dx = -1
	case "s":
		dy = 1
	case "d":
		dx = 1
	default:
		fmt.Println("Invalid direction. Bullet travels straight.")
	}

	return dx, dy
}

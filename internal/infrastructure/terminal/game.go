package terminal

import (
	"bufio"
	"fmt"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain/entities"
)

type GameHandler struct {
	gameConfigUseCase GameConfigUseCase
	config            *entities.GameConfig
	reader            *bufio.Reader
}

type GameConfigUseCase interface {
	CreateGameConfig(entities.Level, entities.Category) (*entities.GameConfig, error)
	Categories() []entities.Category
	Levels() []entities.Level
}

func NewGameHandler(gameConfigUseCase GameConfigUseCase) *GameHandler {
	return &GameHandler{
		gameConfigUseCase: gameConfigUseCase,
		reader:            bufio.NewReader(os.Stdin),
	}
}

func (g *GameHandler) Start() {
	g.prepareScreen()

	for {
		fmt.Println("[1] Start Game")
		fmt.Println("[2] Settings")
		fmt.Println("[3] Exit")

		choice, _ := g.readChoice()

		switch choice {
		case "1":
			g.startGame()

			fmt.Println("Press Enter to return to main menu...")
			g.reader.ReadString('\n')

			g.prepareScreen()
		case "2":
			config, err := g.setUpGameConfig()
			if err != nil {
				fmt.Fprintf(os.Stderr, "GameHandler.Start: failed to set up game config: %v", err)
				break
			}
			g.config = config

			g.prepareScreen()
		case "3":
			return
		default:
			g.prepareScreen()
			fmt.Println(InvalidInputMsg)
		}
	}
}

// Если установить случайные настройки, то в следующую игру они остаются такими же
func (g *GameHandler) startGame() {
	fmt.Println("game has started")

	// Uncomment when random settings are needed
	// if g.settings == nil {
	// 	g.settings = entities.RandomSettings()
	// }

	fmt.Printf("Level: %s, Category: %s\n", g.config.Level(),  g.config.Category())
}

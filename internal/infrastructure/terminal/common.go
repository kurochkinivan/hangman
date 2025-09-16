package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func (g *GameHandler) prepareScreen() {
	g.clearTerminal()
	g.printBanner()
}

func (g *GameHandler) printBanner() {
	fmt.Println(asciiHangman)
}

func (g *GameHandler) clearTerminal() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (g *GameHandler) readChoice() (string, error) {
	l, err := g.reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("readChoice: failed to read input: %w", err)
	}

	return strings.TrimSpace(l), nil
}

func (g *GameHandler) readInt() (int, error) {
	l, err := g.readChoice()
	if err != nil {
		return -1, fmt.Errorf("readInt: failed to read input: %w", err)
	}

	input, err := strconv.Atoi(l)
	if err != nil {
		return -1, fmt.Errorf("readInt: failed to convert input to int: %w", err)
	}

	return input, nil
}

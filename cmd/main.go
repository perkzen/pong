package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"pong/pkg/game"
)

func main() {
	err := tea.NewProgram(game.NewGame(), tea.WithAltScreen()).Start()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}

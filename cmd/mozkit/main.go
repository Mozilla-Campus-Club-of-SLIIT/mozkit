package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal"
)

func main() {
	model := internal.NewModel()
	program := tea.NewProgram(model)

	if _, err := program.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Oh no! The penguins are panicked. Moxy, do something!!! \nError: %v", err)
		os.Exit(1)
	}
}

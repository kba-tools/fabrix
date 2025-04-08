package utils

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func ShowMainMenu() {

	asciiArt := " _____     _          _      \n" +
		"|  ___|_ _| |__  _ __(_)_  __\n" +
		"| |_ / _' | '_ \\| '__| \\ \\/ /\n" +
		"|  _| (_| | |_) | |  | |>  < \n" +
		"|_|  \\__,_|_.__/|_|  |_/_/\\_\\\n"

	asciiPrint := lipgloss.NewStyle().
		Blink(true).
		Foreground(lipgloss.Color("#266ed4")).
		Bold(true).
		Render(asciiArt)

	boxStyle := lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("62")).
		Foreground(lipgloss.Color("233")).
		Align(lipgloss.Center)
	// Style for the text
	textStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#3deb34")).
		Italic(true).
		Align(lipgloss.Center)

	// Define the message
	message := `The helper tool for chaincode developers to create fabric network, 
it does all the heavy lifting for you!!! 

You will be guided throughout the process.

Let's start...`

	combinedText := asciiPrint + "\n\n" + textStyle.Render(message)

	fmt.Println(boxStyle.Render(combinedText))

}

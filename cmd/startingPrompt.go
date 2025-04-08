package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

func customTheme() *huh.Theme {
	t := huh.ThemeBase()
	var (
		normalFg = lipgloss.AdaptiveColor{Light: "235", Dark: "252"}
		indigo   = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
		cream    = lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#FFFDF5"}
		fuchsia  = lipgloss.Color("#F780E2")
		green    = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
		red      = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}
	)

	t.Focused.Base = t.Focused.Base.BorderForeground(lipgloss.Color("238"))
	t.Focused.Title = t.Focused.Title.Foreground(indigo).Bold(true)
	t.Focused.NoteTitle = t.Focused.NoteTitle.Foreground(indigo).Bold(true).MarginBottom(1)
	t.Focused.Directory = t.Focused.Directory.Foreground(indigo)
	t.Focused.Description = t.Focused.Description.Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"})
	t.Focused.ErrorIndicator = t.Focused.ErrorIndicator.Foreground(red)
	t.Focused.ErrorMessage = t.Focused.ErrorMessage.Foreground(red)
	t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(fuchsia)
	t.Focused.NextIndicator = t.Focused.NextIndicator.Foreground(fuchsia)
	t.Focused.PrevIndicator = t.Focused.PrevIndicator.Foreground(fuchsia)
	t.Focused.Option = t.Focused.Option.Foreground(normalFg)
	t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.Foreground(fuchsia)
	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(green)
	t.Focused.SelectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#02CF92", Dark: "#02A877"}).SetString("✓ ")
	t.Focused.UnselectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"}).SetString("• ")
	t.Focused.UnselectedOption = t.Focused.UnselectedOption.Foreground(normalFg)
	t.Focused.FocusedButton = t.Focused.FocusedButton.Foreground(cream).Background(fuchsia)
	t.Focused.Next = t.Focused.FocusedButton
	t.Focused.BlurredButton = t.Focused.BlurredButton.Foreground(normalFg).Background(lipgloss.AdaptiveColor{Light: "252", Dark: "237"})

	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(green)
	t.Focused.TextInput.Placeholder = t.Focused.TextInput.Placeholder.Foreground(lipgloss.AdaptiveColor{Light: "248", Dark: "238"})
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(fuchsia)

	t.Blurred = t.Focused
	t.Blurred.Base = t.Focused.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()
	return t
}

// networkCmd represents the network command
var startingPromptCmd = &cobra.Command{
	Use:     "startingPrompt",
	Aliases: []string{"sp"},
	Short:   "This command is for showing the application starting prompt",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {

		var mainChoice string

		deployCC := huh.NewForm(
			huh.NewGroup(

				huh.NewSelect[string]().
					Options(huh.NewOptions("Create new Network configuration",
						"Choose existing configuration",
						"Exit")...).
					Title("Select an Option").
					Value(&mainChoice),
			),
		).WithShowHelp(true).WithTheme(customTheme()).WithWidth(80)
		err := deployCC.Run()

		if err != nil {
			if err == huh.ErrUserAborted {
				os.Exit(130)
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}

		switch mainChoice {

		case "Create new Network configuration":
			rootCmd.SetArgs([]string{"c"})
			rootCmd.Execute()

		case "Choose existing configuration":
			rootCmd.SetArgs([]string{"l"})
			rootCmd.Execute()

		case "Exit":
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(startingPromptCmd)
}

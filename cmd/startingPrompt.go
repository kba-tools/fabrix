package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

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
					Options(huh.NewOptions("Create new Network configuration", "Choose existing configuration", "Exit")...).
					Title("Select an Option").
					// Description("What you want to do?").
					Value(&mainChoice),
			),
		).WithShowHelp(true).WithTheme(huh.ThemeCharm())

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

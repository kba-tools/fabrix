package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// networkCmd represents the network command
var domainOptionsPromptCmd = &cobra.Command{
	Use:     "domainOptionsPrompt",
	Aliases: []string{"dp"},
	Short:   "Use this command to remove all files for a network",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		domainOptions()
	},
}

func init() {
	rootCmd.AddCommand(domainOptionsPromptCmd)
}

func domainOptions() {

	networkOptions := ""
	nwOptions := huh.NewForm(
		huh.NewGroup(

			huh.NewSelect[string]().
				Options(huh.NewOptions("Start network", "Info", "Deploy chaincode", "Upgrade chaincode", "Down network", "Remove domain", "Home", "Exit")...).
				Title("Select an Option").
				// Description("What you want to do?").
				Value(&networkOptions),
		),
	).WithShowHelp(true).WithTheme(huh.ThemeCharm())

	err := nwOptions.Run()

	if err != nil {
		if err == huh.ErrUserAborted {
			os.Exit(130)
		}
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	switch networkOptions {

	case "Start network":
		rootCmd.SetArgs([]string{"up", choosenDomain})
		rootCmd.Execute()

	case "Info":
		rootCmd.SetArgs([]string{"list"})
		rootCmd.Execute()

	case "Deploy chaincode":
		rootCmd.SetArgs([]string{"deploy", choosenDomain})
		rootCmd.Execute()

	case "Upgrade chaincode":
		rootCmd.SetArgs([]string{"upgrade", choosenDomain})
		rootCmd.Execute()

	case "Down network":
		rootCmd.SetArgs([]string{"down", choosenDomain})
		rootCmd.Execute()

	case "Remove domain":
		rootCmd.SetArgs([]string{"remove", choosenDomain})
		rootCmd.Execute()

	case "Home":
		rootCmd.SetArgs([]string{"sp"})
		rootCmd.Execute()

	case "Exit":
		fmt.Println("Exiting...")
		os.Exit(0)
	}

}

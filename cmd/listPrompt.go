package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// networkCmd represents the network command
var listPromptCmd = &cobra.Command{
	Use:     "listPrompt",
	Aliases: []string{"lp"},
	Short:   "Use this command to remove all files for a network",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		chooseDomain()
	},
}

func init() {
	rootCmd.AddCommand(listPromptCmd)
}

var choosenDomain string

func chooseDomain() {

	rootDir := "./fabrix"
	// Read the directory
	entries, err := os.ReadDir(rootDir)
	if err != nil {
		_ = fmt.Errorf("error reading directory: %v\n", err)
	}
	var domains []string
	var isDomains bool
	var goBack bool
	// List directories
	for _, entry := range entries {
		if entry.IsDir() {
			domains = append(domains, entry.Name())
		}
	}

	if len(domains) > 0 {
		isDomains = true
	}

	listOfDomains := huh.NewForm(

		huh.NewGroup(huh.NewConfirm().
			Title("No domains available, create new configuartion and come back here!!").
			Value(&goBack).
			Affirmative("Go back").
			Negative("Exit")).WithHideFunc(func() bool {
			return isDomains
		}),

		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions(domains...)...).
				Title("Available domains").
				Description("Choose the domain").
				Value(&choosenDomain),
		).WithHideFunc(func() bool {
			return !isDomains
		}),
		
	).WithShowHelp(true).WithTheme(huh.ThemeCharm())

	err = listOfDomains.Run()

	if err != nil {
		if err == huh.ErrUserAborted {
			os.Exit(130)
		}
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	if choosenDomain != "" {
		rootCmd.SetArgs([]string{"dp"})
		rootCmd.Execute()
	}

	if goBack {
		rootCmd.SetArgs([]string{"sp"})
		rootCmd.Execute()
	} else {
		fmt.Println("Exiting...")
		os.Exit(0)
	}

}

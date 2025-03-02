package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
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

type model2 struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m model2) Init() tea.Cmd {
	return nil
}

func (m model2) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model2) View() string {
	if m.choice != "" {
		choosenDomain = m.choice
		return quitTextStyle.Render(fmt.Sprintf("Ok, Choosen Network: %s", m.choice))
	}
	if m.quitting {
		return quitTextStyle.Render("Exit? ... See you later")
	}
	return "\n" + m.list.View()
}

func chooseDomain() {

	rootDir := "./fabrix"
	// Read the directory
	entries, err := os.ReadDir(rootDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
	}
	var domains []string
	// List directories
	for _, entry := range entries {
		if entry.IsDir() {
			domains = append(domains, entry.Name())
		}
	}

	

	// Append domain names as selectable items
	// for _, domain := range domains {
	// 	items = append(items, item(domain))
	// }

	networkOptions := ""

	listOfDomains := huh.NewForm(
		huh.NewGroup(

			huh.NewSelect[string]().
				Options(huh.NewOptions(domains...)...).
				Title("Select an Option").
				Description("Choose the domain").
				Value(&networkOptions),
		),
	).WithShowHelp(true).WithTheme(huh.ThemeCharm())

	err = listOfDomains.Run()

	if err != nil {
		if err == huh.ErrUserAborted {
			os.Exit(130)
		}
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

}

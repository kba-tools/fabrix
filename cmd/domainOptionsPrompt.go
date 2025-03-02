package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
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

type model3 struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m model3) Init() tea.Cmd {
	return nil
}

func (m model3) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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
			m.quitting = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model3) View() string {
	if m.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("You have choosen : %s", m.choice))
	}
	if m.quitting {
		return quitTextStyle.Render("Exit? ... See you later")
	}
	return "\n" + m.list.View()
}

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

// var choosenDomain string

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
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

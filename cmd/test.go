package cmd

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/bubbles/list"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Use this command to test different functionalities.",
	Run: func(cmd *cobra.Command, args []string) {
		tabs := []string{"Create New Domain", "Choose Existing Domain", "Exit"}
		m := newModel(tabs)
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}

// Define a custom list item struct
type listItem struct {
	title string
}

// Implement fmt.Stringer for rendering
func (i listItem) String() string {
	return i.title
}

// Implement FilterValue() required by list.Item interface
func (i listItem) FilterValue() string {
	return i.title
}

type model struct {
	Tabs      []string
	activeTab int
	list      list.Model
	showList  bool
}

func newModel(tabs []string) model {
	// Create list items using our struct
	items := []list.Item{
		listItem{"xfsafsg"},
		listItem{"asdfhn"},
		listItem{"wqetytuyiuk"},
	}

	// Create list model
	listModel := list.New(items, list.NewDefaultDelegate(), 20, 20)
	listModel.Title = "Select an Option"

	return model{
		Tabs:     tabs,
		list:     listModel,
		showList: false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "right", "l", "n", "tab":
			m.activeTab = min(m.activeTab+1, len(m.Tabs)-1)
			m.showList = m.activeTab == 1 
			return m, nil
		case "left", "h", "p", "shift+tab":
			m.activeTab = max(m.activeTab-1, 0)
			m.showList = m.activeTab == 1
			return m, nil
		}

		// Pass input to the list only if it's active
		if m.showList {
			var cmd tea.Cmd
			m.list, cmd = m.list.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}

func (m model) View() string {
	doc := strings.Builder{}

	// Render Tabs
	var renderedTabs []string
	for i, t := range m.Tabs {
		style := inactiveTabStyle
		if i == m.activeTab {
			style = activeTabStyle
		}
		renderedTabs = append(renderedTabs, style.Render(t))
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row + "\n")

	// Render content based on the active tab
	if m.showList {
		doc.WriteString(windowStyle.Width(30).Render(m.list.View()))
	} else {
		content := fmt.Sprintf("Content for: %s", m.Tabs[m.activeTab])
		doc.WriteString(windowStyle.Width(30).Render(content))
	}

	return docStyle.Render(doc.String())
}

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	highlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	activeTabStyle    = inactiveTabStyle.Border(activeTabBorder, true)
	windowStyle       = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()
)

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

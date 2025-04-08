package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"

	"github.com/spf13/cobra"
)

// networkCmd represents the network command
var upgradePromptCmd = &cobra.Command{
	Use:     "upgradePrompt",
	Aliases: []string{"uccp"},
	Short:   "Use this command to deploy chaincode.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {

		chaincodeParams := ChaincodeParams{}
		// currentDir, erro := os.Getwd()
		// if erro!= nil{
		// 	fmt.Println("error cd", erro)
		// }

		// fmt.Println("cc dir", currentDir)
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return
		}

		// fmt.Println("full path:", nwPath)

		// fmt.Println("Relative path:", relativePath)

		deployCC := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Chaincode Name").
					Description("What's the name for your Chaincode?").Value(&chaincodeParams.ccName),

				huh.NewSelect[string]().
					Options(huh.NewOptions("golang", "JavaScript", "Java")...).
					Title("Chaincode Language").
					Description("Choose your chaincode language").
					Validate(func(t string) error {
						if t == "Java" {
							return fmt.Errorf("Uh oh! sorry we don't support Java right now")
						}
						return nil
					}).
					Value(&chaincodeParams.ccLang),

				huh.NewInput().
					Title("Chaincode Label").
					Description("What's the Label for your Chaincode?").Value(&chaincodeParams.ccLabel),

				huh.NewFilePicker().
					Title("Chaincode Path").
					Description("Select your chaincode directory.").
					DirAllowed(true).
					FileAllowed(false).
					CurrentDirectory(homeDir).
					ShowPermissions(false).
					Value(&chaincodeParams.ccPath),

				huh.NewInput().
					Title("Chaincode Sequence").
					Description("What's the sequence for your Chaincode?").Value(&chaincodeParams.ccSequence),

				huh.NewInput().
					Title("Chaincode Version").
					Description("What's the version for your Chaincode?").Value(&chaincodeParams.ccVersion),

				huh.NewConfirm().
					Title("Lets deploy your chaincode?").
					Value(&chaincodeParams.deploy).
					Affirmative("Yes!").
					Negative("No."),
			),
		).WithShowHelp(true).WithTheme(huh.ThemeDracula())

		err = deployCC.Run()
		if err != nil {
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}

		nwPath := fmt.Sprintf("./fabrix/%v/Network", choosenDomain)

		// Resolve to absolute (full) path
		nwPath, err = filepath.Abs(nwPath)
		if err != nil {
			log.Fatalf("Failed to get absolute path: %v", err)
		}

		// ccRelativePath, err := filepath.Rel(nwPath, chaincodeParams.ccPath)
		// if err != nil {
		// 	log.Fatalf("Failed to get relative path: %v", err)
		// }

		action := func() {
			// configs.InstallChaincode(choosenDomain, ccRelativePath, chaincodeParams.ccLang, chaincodeParams.ccLabel, chaincodeParams.ccName, chaincodeParams.ccVersion, chaincodeParams.ccSequence)
			// utils.GetAllChaincodeDefenitiions(choosenDomain)
			width, _, err := term.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				fmt.Println("Failed to get terminal size:", err)
				return
			}

			// Create styled text
			styledText := lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render("this is a test message")

			// Calculate the left padding to shift text to the right side
			textWidth := lipgloss.Width(styledText)
			leftPadding := width - textWidth

			// Apply the padding (spaces on the left)
			finalOutput := lipgloss.NewStyle().PaddingLeft(leftPadding).Render(styledText)

			// Print
			fmt.Println(finalOutput)
		}

		if err := spinner.New().Title("Deployment in progress...").Action(action).Run(); err != nil {
			fmt.Println("Failed:", err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(upgradePromptCmd)
}

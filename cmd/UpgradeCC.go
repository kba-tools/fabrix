package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/vineshtk/fabrix/pkg/utils"
	"golang.org/x/term"
)

// networkCmd represents the network command
var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Aliases: []string{"ucc"},
	Short:   "Use this command to deploy chaincode.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {

		ccDefinitions := utils.GetAllChaincodeDefenitions(choosenDomain)
		fmt.Println(ccDefinitions)
		// Define struct for the chaincode definition
		type ChaincodeDefinition struct {
			Name     string `json:"name"`
			Sequence int    `json:"sequence"`
			Version  string `json:"version"`
		}

		// Struct to hold top-level structure
		type ChaincodeDefinitions struct {
			Definitions []ChaincodeDefinition `json:"chaincode_definitions"`
		}

		var defs ChaincodeDefinitions

		err := json.Unmarshal([]byte(ccDefinitions), &defs)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}
		ccNames := []huh.Option[string]{}
		ccMap := make(map[string]ChaincodeDefinition)

		// Print the required fields
		for _, def := range defs.Definitions {

			ccNames = append(ccNames, huh.Option[string]{Key: def.Name, Value: def.Name})
			ccMap[def.Name] = def
		}

		chaincodeParams := ChaincodeParams{}

		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return
		}

		deployCC := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Installed Chaincodes").
					Options(ccNames...).
					Description("Choose chaincode to upgrade").Value(&chaincodeParams.ccName),

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
	rootCmd.AddCommand(upgradeCmd)
}

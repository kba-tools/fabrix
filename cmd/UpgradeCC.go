package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/vineshtk/fabrix/pkg/utils"
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

		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return
		}

		var (
			ccName     string
			ccSequence string
			ccVersion  string
			ccLabel    string
			ccLang     string
			ccPath     string
		)

		versionDescription := func() string {
			selectedCC := ccMap[ccName]
			version := selectedCC.Version
			return fmt.Sprintf("Current version is: %v", version)
		}

		sequenceDescription := func() string {
			selectedCC := ccMap[ccName]
			sequence := selectedCC.Sequence
			return fmt.Sprintf("Current sequence is: %v", sequence)
		}

		versionOption := func() []huh.Option[string] {
			selectedCC := ccMap[ccName]
			version, _ := strconv.ParseFloat(selectedCC.Version, 64)
			return huh.NewOptions(fmt.Sprintf("Upgrade version to %v", version+0.1))
		}

		sequenceOption := func() []huh.Option[string] {
			selectedCC := ccMap[ccName]
			return huh.NewOptions(fmt.Sprintf("Upgrade sequence to %v", float64(selectedCC.Sequence)+1))
		}

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Options(ccNames...).
					Value(&ccName).
					Title("Installed Chaincode(s)").
					Height(5),
					
				huh.NewSelect[string]().
					Value(&ccSequence).
					Title("Chaincode Sequence").
					Height(5).
					DescriptionFunc(sequenceDescription, &ccName).
					OptionsFunc(sequenceOption, &ccName),

				huh.NewSelect[string]().
					Value(&ccVersion).
					Height(5).
					Title("Chaincode Version").
					DescriptionFunc(versionDescription, &ccName).
					OptionsFunc(versionOption, &ccName),

				huh.NewSelect[string]().
					Options(huh.NewOptions("golang", "JavaScript", "Java")...).
					Title("Chaincode Language").
					Height(5).
					Description("Choose your chaincode language").
					Validate(func(t string) error {
						if t == "Java" {
							return fmt.Errorf("uh oh! sorry we don't support Java right now")
						}
						return nil
					}).
					Value(&ccLang),

				huh.NewInput().
					Title("Chaincode Label").
					Description("What's the Label for your Chaincode?").Value(&ccLabel),

				huh.NewFilePicker().
					Title("Chaincode Path").
					Description("Select your chaincode directory.").
					DirAllowed(true).
					FileAllowed(false).
					CurrentDirectory(homeDir).
					ShowPermissions(false).
					Value(&ccPath),
			),
			huh.NewGroup(),
		)

		err = form.Run()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s, %s\n", ccName, ccSequence)

		// nwPath := fmt.Sprintf("./fabrix/%v/Network", choosenDomain)

		// // Resolve to absolute (full) path
		// nwPath, err = filepath.Abs(nwPath)
		// if err != nil {
		// 	log.Fatalf("Failed to get absolute path: %v", err)
		// }

		// ccRelativePath, err := filepath.Rel(nwPath, chaincodeParams.ccPath)
		// if err != nil {
		// 	log.Fatalf("Failed to get relative path: %v", err)
		// }

		// action := func() {
		// 	// configs.InstallChaincode(choosenDomain, ccRelativePath, chaincodeParams.ccLang, chaincodeParams.ccLabel, chaincodeParams.ccName, chaincodeParams.ccVersion, chaincodeParams.ccSequence)
		// 	width, _, err := term.GetSize(int(os.Stdout.Fd()))
		// 	if err != nil {
		// 		fmt.Println("Failed to get terminal size:", err)
		// 		return
		// 	}

		// 	// Create styled text
		// 	styledText := lipgloss.NewStyle().
		// 		BorderStyle(lipgloss.RoundedBorder()).
		// 		BorderForeground(lipgloss.Color("63")).
		// 		Padding(1, 2).
		// 		Render("this is a test message")

		// 	// Calculate the left padding to shift text to the right side
		// 	textWidth := lipgloss.Width(styledText)
		// 	leftPadding := width - textWidth

		// 	// Apply the padding (spaces on the left)
		// 	finalOutput := lipgloss.NewStyle().PaddingLeft(leftPadding).Render(styledText)

		// 	// Print
		// 	fmt.Println(finalOutput)
		// }

		// if err := spinner.New().Title("Deployment in progress...").Action(action).Run(); err != nil {
		// 	fmt.Println("Failed:", err)
		// 	return
		// }

	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}

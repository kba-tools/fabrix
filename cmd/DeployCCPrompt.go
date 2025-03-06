package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/vineshtk/fabrix/pkg/configs"

	"github.com/spf13/cobra"
)

type ChaincodeParams struct {
	ccPath     string
	ccLang     string
	ccName     string
	ccLabel    string
	ccVersion  string
	ccSequence string
	deploy     bool
}

// networkCmd represents the network command
var deployPromptCmd = &cobra.Command{
	Use:     "deployPrompt",
	Aliases: []string{"dccp"},
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

		fmt.Println("PATH FROM SELECTOR", chaincodeParams.ccPath)

		nwPath := fmt.Sprintf("./fabrix/%v/Network", choosenDomain)

		// Resolve to absolute (full) path
		nwPath, err = filepath.Abs(nwPath)
		if err != nil {
			log.Fatalf("Failed to get absolute path: %v", err)
		}

		ccRelativePath, err := filepath.Rel(nwPath, chaincodeParams.ccPath)
		if err != nil {
			log.Fatalf("Failed to get relative path: %v", err)
		}

		fmt.Println("rel  path", ccRelativePath)

		fmt.Println(chaincodeParams.ccLabel, chaincodeParams.ccLang, chaincodeParams.ccName, ccRelativePath, chaincodeParams.ccSequence, chaincodeParams.ccVersion)

		// call the function to deploy chaincode to network
		configs.InstallChaincode(choosenDomain, ccRelativePath, chaincodeParams.ccLang, chaincodeParams.ccLabel, chaincodeParams.ccName, chaincodeParams.ccVersion, chaincodeParams.ccSequence)

	},
}

func init() {
	rootCmd.AddCommand(deployPromptCmd)
}

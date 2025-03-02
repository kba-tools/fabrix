package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"

	"github.com/spf13/cobra"
)

// type mainChoice struct {
// 	choice int
// }

// networkCmd represents the network command
var applicationStart = &cobra.Command{
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

		err := deployCC.Run()
		if err != nil {
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}

		fmt.Println("PATH FROM SELECTOR", chaincodeParams.ccPath)

		// ccRelativePath, err := filepath.Rel(currentDir, chaincodeParams.ccPath)

		// if err != nil {
		// 	fmt.Println("error in rel path finding", err)

		// }

		fmt.Println(chaincodeParams.ccLabel, chaincodeParams.ccLang, chaincodeParams.ccName, chaincodeParams.ccPath, chaincodeParams.ccSequence, chaincodeParams.ccVersion)

		// call the function to deploy chaincode to network
		// configs.InstallChaincode("auto.com", ccRelativePath, chaincodeParams.ccLang, chaincodeParams.ccLabel, chaincodeParams.ccName, chaincodeParams.ccVersion, chaincodeParams.ccSequence)

	},
}

func init() {
	rootCmd.AddCommand(applicationStart)
}

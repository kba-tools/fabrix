package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vineshtk/fabrix/pkg/utils"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "fabrix",
	Version: version,
	Short:   "fabrix is a tool to create a fabric network",
	Long: `fabrix is a tool,
	that helps chaincode developers to setup a fabric network easily,
	for deploying and testing the chaincode.`,

	Run: func(cmd *cobra.Command, args []string) {
		utils.ShowMainMenu()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
	rootCmd.SetArgs([]string{"sp"})
	rootCmd.Execute()

}

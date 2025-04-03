
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ccPath string
// var ccVersion string
var ccLang string

// networkCmd represents the network command
var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"dcc"},
	Short:   "Use this command to deploy chaincode.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("choosen domain:", args[0])
		rootCmd.SetArgs([]string{"dccp"})
		rootCmd.Execute()
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}

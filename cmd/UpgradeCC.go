
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


// networkCmd represents the network command
var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Aliases: []string{"ucc"},
	Short:   "Use this command to deploy chaincode.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("choosen domain:", args[0])
		rootCmd.SetArgs([]string{"test"})
		rootCmd.Execute()
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}

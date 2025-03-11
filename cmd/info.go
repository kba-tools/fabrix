package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vineshtk/fabrix/pkg/configs"
)

// networkCmd represents the network command
var infoCmd = &cobra.Command{
	Use: "info",
	// Aliases: []string{"sn"},
	Args:  cobra.ExactArgs(1),
	Short: "Use this command to start a network and install chaincode",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {

		configs.PrintNetworkInfo(args[0])

	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// shakeCmd represents the shake command
var shakeCmd = &cobra.Command{
	Use:   "shake",
	Short: "Looks up the average price for a watch",
	Long:  `Looks up the average price of the specified watch on all supported marketplaces.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shake called")
	},
}

func init() {
	rootCmd.AddCommand(shakeCmd)

	shakeCmd.Flags().StringP("manufacturer", "m", "", "Manufacturer of the desired watch")
	shakeCmd.Flags().StringP("ref", "r", "", "Reference Nr. of the desired watch")
}

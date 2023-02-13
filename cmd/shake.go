package cmd

import (
	"github.com/ruegerj/chrono-shaker/internal/common"
	"github.com/ruegerj/chrono-shaker/internal/shaker"
	"github.com/spf13/cobra"
)

const MANUFACTURER_ARG = "manufacturer"
const REF_NO_ARG = "ref"

// shakeCmd represents the shake command
var shakeCmd = &cobra.Command{
	Use:   "shake",
	Short: "Looks up the average price for a watch",
	Long:  `Looks up the average price of the specified watch on all supported marketplaces.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: replace through proper cli arg
		const platform = common.CHRONO_24

		brand := cmd.Flag(MANUFACTURER_ARG).Value.String()
		refNo := cmd.Flag(REF_NO_ARG).Value.String()

		filter := common.NewFilterOptions(brand, refNo)
		shaker, err := shaker.NewShaker(platform, filter)

		if err != nil {
			return err
		}

		results := shaker.ShakeListings()

		for _, res := range results {
			cmd.Println(res.String())
		}

		return nil
	},
}

//nolint:errcheck
func init() {
	rootCmd.AddCommand(shakeCmd)

	shakeCmd.Flags().StringP(MANUFACTURER_ARG, "m", "", "Manufacturer of the desired watch")
	shakeCmd.Flags().StringP(REF_NO_ARG, "r", "", "Reference Nr. of the desired watch")

	shakeCmd.MarkFlagRequired(MANUFACTURER_ARG)
	shakeCmd.MarkFlagRequired(REF_NO_ARG)
}

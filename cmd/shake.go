package cmd

import (
	"fmt"

	"github.com/ruegerj/chrono-shaker/internal/common"
	"github.com/ruegerj/chrono-shaker/internal/shaker"
	"github.com/ruegerj/chrono-shaker/internal/utils"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

const MANUFACTURER_ARG = "manufacturer"
const REF_NO_ARG = "ref"
const DECIMAL_PLACES = 3

// shakeCmd represents the shake command
var shakeCmd = &cobra.Command{
	Use:   "shake",
	Short: "Looks up the average price for a watch",
	Long:  `Looks up the average price of the specified watch on all supported platforms.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		brand := cmd.Flag(MANUFACTURER_ARG).Value.String()
		refNo := cmd.Flag(REF_NO_ARG).Value.String()

		filter := common.NewFilterOptions(brand, refNo, common.AllPlatforms)
		shaker := shaker.NewShaker(filter)

		results := shaker.ShakeListings()

		platformResults := utils.GroupBy(results, func(item common.WatchListing) common.Platform {
			return item.Platform
		})

		allAveragePrices := make([]decimal.Decimal, 0)

		for key, listings := range platformResults {
			prices := make([]decimal.Decimal, len(listings))

			for i := 0; i < len(listings); i++ {
				prices[i] = listings[i].Price.Value
			}

			avgPrice := decimal.Avg(prices[0], prices[1:]...).Round(DECIMAL_PLACES)
			allAveragePrices = append(allAveragePrices, avgPrice)

			fmt.Printf("[%s]: $%d, %d listing(s)\n", key, avgPrice.CoefficientInt64(), len(prices))
		}

		totalAvg := decimal.Avg(allAveragePrices[0], allAveragePrices[1:]...).Round(DECIMAL_PLACES)

		fmt.Printf("\nTotal: $%d\n", totalAvg.Abs().CoefficientInt64())

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

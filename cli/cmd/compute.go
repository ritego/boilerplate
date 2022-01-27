package cmd

import (
	"fmt"
	"time"

	"github.com/aellacredit/jara/service"
	"github.com/aellacredit/jara/utils"
	"github.com/spf13/cobra"
)

var Compute = &cobra.Command{
	Use:   "compute",
	Short: "...",
	Long:  `...`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		entity := args[0]

		class, _ := cmd.Flags().GetString("class")
		if !utils.InSlice(class, []string{utils.ALL_CLASS, utils.HAS_INTEREST_CLASS, utils.NO_INTEREST_CLASS}) {
			fmt.Printf("Class %s is not valid. class must be one of %s, %s, %s", class, utils.ALL_CLASS, utils.HAS_INTEREST_CLASS, utils.NO_INTEREST_CLASS)
			return
		}

		period, _ := cmd.Flags().GetString("period")
		if !utils.InSlice(period, []string{utils.DAILY_PERIOD, utils.WEEKLY_PERIOD, utils.MONTHLY_PERIOD, utils.QUATERLY_PERIOD, utils.YEARLY_PERIOD}) {
			fmt.Printf("Period %s is not valid. period must be one of %s, %s, %s, %s, %s", period, utils.DAILY_PERIOD, utils.WEEKLY_PERIOD, utils.MONTHLY_PERIOD, utils.QUATERLY_PERIOD, utils.YEARLY_PERIOD)
			return
		}

		from, _ := cmd.Flags().GetString("from")
		_, err := time.Parse("2006-01-02 15:04:05", from)
		if err == nil {
			fmt.Printf("From (%s) is not valid. Time format must be in YYYY-MM-DD HH:MM:SS", from)
		}

		to, _ := cmd.Flags().GetString("to")
		_, err = time.Parse("2006-01-02 15:04:05", to)
		if err == nil {
			fmt.Printf("To (%s) is not valid. Time format must be in YYYY-MM-DD HH:MM:SS", to)
		}

		chunk, _ := cmd.Flags().GetInt64("chunk")

		fmt.Printf("Compute %s %s %s %s %d\n", class, period, from, to, chunk)

		switch entity {
		case "settlement":
			service.SettlementService.Run(class, period, from, to, chunk)
		case "payout":
			service.PayoutService.Run(class, period, from, to, chunk)
		}
		fmt.Println("compute called")
	},
}

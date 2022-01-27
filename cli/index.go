package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/aellacredit/jara/cli/cmd"
	"github.com/spf13/cobra"
)

func Init() {

	root := &cobra.Command{
		Use:   "jara",
		Short: "Jara is a wallet interest calculator",
		Long:  `...`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
				Welcome to Jara!

				Jara is a wallet interest calculator.

				To get started, run: jara --help
			`)
		},
	}

	root.AddCommand(cmd.Server)

	root.AddCommand(cmd.Compute)
	cmd.Compute.Flags().String("class", "all", "all|has-interest|no-interest")
	cmd.Compute.Flags().String("period", "daily", "daily|weekly|monthly|yearly")
	cmd.Compute.Flags().String("from", time.Now().Format("01-JAN-2006 15:04:00"), "yyyy-mm-dd hh:mm:ss")
	cmd.Compute.Flags().String("to", time.Now().Format("01-JAN-2006 15:04:00"), "yyyy-mm-dd hh:mm:ss")
	cmd.Compute.Flags().Int("chunk", 12, "how to chuck the records during the computation")

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

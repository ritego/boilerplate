package cmd

import (
	"github.com/aellacredit/jara/http"
	"github.com/spf13/cobra"
)

var Server = &cobra.Command{
	Use:   "server",
	Short: "...",
	Long:  `...`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		http.Init()
	},
}

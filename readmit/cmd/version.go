package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Readmit",
	Long:  `All software has versions. This is Readmit's.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Readmit CLI v%s\n", rootCmd.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

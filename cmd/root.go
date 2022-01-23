package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocmdb",
	Short: "gocmdb",
	Long:  `gocmdb`,
}

var verbose bool

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose")
}

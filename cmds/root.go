/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-15 23:33:27
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-15 23:37:20
 */
package cmds

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "cmdb",
	Short: "cmdb program",
	Long:  "cmdb programs",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("cmdb")
		return nil
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}

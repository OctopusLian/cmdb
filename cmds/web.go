/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-15 23:33:36
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-15 23:38:44
 */
package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var webCommand = &cobra.Command{
	Use:   "web",
	Short: "web console",
	Long:  "web console",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("web")
		return nil
	},
}

func init() {
	rootCommand.AddCommand(webCommand)
}

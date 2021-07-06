package cmd

import (
	"fmt"
	"github.com/jellycheng/gosupport/xversion"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd) //添加命令
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "查看版本",
	Long:  `查看版本`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(xversion.Version)
	},
}


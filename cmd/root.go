package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "go-markdowndoc",
	Short: "md文件转html展示",
	Long: `markdown文档转html、文档管理
        `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("请使用查看 ./go-markdowndoc -h 命令帮助，")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


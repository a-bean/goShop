package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

// rootCmd代表base 命令
var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "a brief description of your application",
	Long:  "a longer description",
}

// 命令1
var mockMsgCmd = &cobra.Command{
	Use:   "mockMsg",
	Short: "批量发送测试数据",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("call mockMsg")
		return nil
	},
}

// 命令2
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "导出数据",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("call export")
		return nil
	},
}

func main() {
	rootCmd.AddCommand(mockMsgCmd)
	rootCmd.AddCommand(exportCmd)

	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/senyu-up/toolbox/tool/cron"
	"github.com/spf13/cobra"
	"template/boot"
	"template/global"
	"template/index"
)

// cronjobCmd represents the cronjob command
var cronjobCmd = &cobra.Command{
	Use:              "cronjob",
	TraverseChildren: true,
	Short:            "template cron job, can only run as a single instance!",
	Long:             ``,
	PreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cronjob called")
		// 初始化资源
		err := boot.Boot(global.ConfigPath)
		if err != nil {
			panic(err)
		}

		// 注册 cron job
		err = index.RegisterCronJob()
		if err != nil {
			panic(err)
		}
		// 启动 cron job
		cron.Start()
		go func() {
			// 健康检查
			global.ErrChan <- global.GetFacade().GetHealthChecker().Start()
		}()
	},
}

func init() {
	rootCmd.AddCommand(cronjobCmd)
}

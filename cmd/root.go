package cmd

import (
	"context"
	"fmt"
	"github.com/senyu-up/toolbox/tool/http/gin_server/middleware"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
	"template/boot"
	"template/global"
	"template/index"
	"template/internal/model"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:              "",
	TraverseChildren: true,
	Short:            "root cmd, start http",
	Long:             ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("server pre run called")
		fmt.Printf("root cmd init called, %s \n", global.ConfigPath)

		// 初始化服务的资源
		err := boot.Boot(global.ConfigPath)
		if err != nil {
			global.ErrChan <- fmt.Errorf("boot err %v", err)
			return
		}

		// 获取 fiber app
		var app = global.GetFacade().GetGin()
		if app == nil {
			global.ErrChan <- fmt.Errorf("gin app is nil")
			return
		}

		// gin app 注册中间件
		app.Gin().Use(
			middleware.CorsMiddleware(),
			middleware.PanicRecoverMiddleware,
			middleware.SetRequestId())

		// 注册路由
		index.RegisterRouter(app.Gin())
		model.DbAutoMigrate(global.GetFacade().GetEnv().Stage, "")
	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		go func() {
			// 启动服务, 接收错误并投递到全局错误通道
			global.ErrChan <- global.GetFacade().StartGin()
		}()
		go func() {
			global.ErrChan <- global.GetFacade().StartHealthChecker()
		}()
	},

	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		for {
			select {
			case err := <-global.ErrChan:
				log.Printf("get global error %v\n", err)
				global.GetFacade().Shutdown(global.Ctx)
				global.Cancel() // 全局 ctx 取消
				return
			case <-global.Ctx.Done():
				log.Printf("get global ctx canceld\n")
				global.GetFacade().Shutdown(context.Background())
				return
			case s := <-c:
				log.Printf("get a signal %s\n", s.String())
				global.GetFacade().Shutdown(global.Ctx)
				global.Cancel() // 全局 ctx 取消
				return
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&global.ConfigPath, "conf", "c", ".", "Config file path")
	return
}

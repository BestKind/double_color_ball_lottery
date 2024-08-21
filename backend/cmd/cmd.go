package cmd

import (
	"double_color_ball_lottery/backend/app"
	"double_color_ball_lottery/backend/db"
	"double_color_ball_lottery/backend/server"
	"double_color_ball_lottery/backend/services"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var (
	version    bool
	versionCmd = &cobra.Command{
		Use:     "version",
		Aliases: []string{"--version", "-v"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("application %s version command get %s\n", app.Name, app.Version)
		},
		Short: "Show application version",
		Args:  cobra.MaximumNArgs(0),
	}
	initDataCmd = &cobra.Command{
		Use:     "init",
		Aliases: []string{"--init", "-i"},
		PreRun: func(cmd *cobra.Command, args []string) {
			if version {
				return
			}
			// 做一些初始化的事情
			preRun()
		},
		Run: func(cmd *cobra.Command, args []string) {
			services.NewLotteryService().InitHistoryData()
		},
		Short: "initialize data",
		Args:  cobra.MaximumNArgs(0),
	}
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Long:  "application long",
	Short: "application short",
	PreRun: func(cmd *cobra.Command, args []string) {
		if version {
			return
		}
		// 做一些初始化的事情
		preRun()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Printf("application %s version is %s\n", app.Name, app.Version)
			return
		}
		run()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	},
	DisableAutoGenTag: true,
	SilenceUsage:      true,
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "Version flag show application version")
}

func preRun() {
	// 1. init Config
	// fmt.Println("config init success")
	// 2. init log
	// fmt.Println("log init success")
	// 3. init trace
	// fmt.Println("trace init success")
	// 4. init mysql
	if db.DB == nil {
		db.WithDB(db.NewMysql(
			db.WithDBType("mysql"),
			db.WithHost("localhost"),
			db.WithPort("3306"),
			db.WithDBName("lottery"),
			db.WithTablePrefix(""),
			db.WithUsername("admin"),
			db.WithPassword("admin"),
			db.WithMaxConns(300),
			db.WithIdelConns(100),
		))
		fmt.Println("mysql init success")
	}
	// 5. init redis
	// fmt.Println("redis init success")
	// 6. init oss 如果供应商有很多个,考虑在使用时再创建
	// fmt.Println("oss init success")
	// 7. init kafka
	// 8. publish-subscribe
}

func run() {
	// start server & stop graceful
	httpServer := server.NewHTTPServer(9500)
	httpServer.Start()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	sig := <-ch
	fmt.Println("Start stop close service, sig: ", sig)
	httpServer.Stop()
	// 等待各个服务有序 close
	time.Sleep(3 * time.Second)
	fmt.Println("Close service finish")
}

func Execute() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initDataCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

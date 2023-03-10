/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"saniloader/config"
	"saniloader/metrics"
	"saniloader/server"

	"github.com/spf13/cobra"
)


func startNormalMode() error {
	go metrics.RunMetrics()
	server.Serve()
	return nil
}

func startDynamicMode() error {
	go metrics.RunMetrics()
	go server.Serve()

	go config.MakeConfigDynamic()

	for {
		cfg := <- config.ConfigChannel
		server.ServerConfig = cfg
	}
}

func RunStart (cmd *cobra.Command, args []string) {
	if config.OnlyConfig && config.ConfigPath == "none" {
		fmt.Println("You should use --config flag with --only flag.")
		os.Exit(1)
	}

	if config.DynamicMode {
		err := startDynamicMode()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		err := startNormalMode()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start SaniLoader",
	Long: `Use start command to start saniloader.
if you want to enter dynamic mode, you must specify the proper flag. By default it's disabled.`,
	Run: RunStart,
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.PersistentFlags().StringVarP(&config.ConfigPath, "config", "c", "none", "use config file within path/to/config.json")
	startCmd.Flags().BoolVarP(&config.DynamicMode, "dynamic", "d", false, "use this flag to enter dynamic mode.")
	startCmd.Flags().BoolVarP(&config.OnlyConfig, "only", "o", false, "use this flag to only use the specified config file to proceed")
}

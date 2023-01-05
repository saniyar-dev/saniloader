/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"saniloader/config"
	"saniloader/server"

	"github.com/spf13/cobra"
)


func readConfigFile() (config.ConfigType, error) {
	if config.ConfigPath != "none" {
		return config.ReadConfig(config.ConfigPath)
	}
	return config.ConfigType{}, nil
}

func getCfgMade() (config.ConfigType, error) {
	if config.OnlyConfig {
		return config.ConfigType{}, nil
	} 
	return config.MakeConfig()
}

func RunStart (cmd *cobra.Command, args []string) {
	if config.OnlyConfig && config.ConfigPath == "none" {
		fmt.Println("You should use --config flag with --only flag.")
		os.Exit(1)
	}

	cfgFile, err := readConfigFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfgMade, err := getCfgMade()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfg := config.CombineConfigs(cfgFile, cfgMade)
	server.Serve(cfg)
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

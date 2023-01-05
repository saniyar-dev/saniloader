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

var ConfigPath string = "none"
var DynamicMode bool = false

func readConfigFile() (config.ConfigType, error) {
	if ConfigPath != "none" {
		return config.ReadConfig(ConfigPath)
	}
	return config.ConfigType{}, nil
}


func RunStart (cmd *cobra.Command, args []string) {
	cfgFile, err := readConfigFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfgMade, err := config.MakeConfig()

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

	startCmd.PersistentFlags().StringVarP(&ConfigPath, "config", "c", "none", "use config file within path/to/config.json")
	startCmd.Flags().BoolVarP(&DynamicMode, "dynamic", "d", false, "use this flag to enter dynamic mode.")
}

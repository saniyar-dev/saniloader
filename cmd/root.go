/*
Copyright © 2023 saniyar.dev
*/
package cmd

import (
	"fmt"
	"os"
	"saniloader/config"
	"saniloader/server"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "saniloader",
	Short: "Load Balancer",
	Long: `SaniLoader is a custom load balancer tool written in golang`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// err := rootCmd.Execute()
	cfg, err := config.MakeConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	server.Serve(cfg)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.saniloader.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



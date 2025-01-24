/*
Copyright © 2025 Herman Gonçalves hermangoncalves@outlook.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/hermangoncalves/routerflow/pkg/config"
	"github.com/hermangoncalves/routerflow/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const logo = `

██████╗  ██████╗ ██╗   ██╗████████╗███████╗██████╗ ███████╗██╗      ██████╗ ██╗    ██╗
██╔══██╗██╔═══██╗██║   ██║╚══██╔══╝██╔════╝██╔══██╗██╔════╝██║     ██╔═══██╗██║    ██║
██████╔╝██║   ██║██║   ██║   ██║   █████╗  ██████╔╝█████╗  ██║     ██║   ██║██║ █╗ ██║
██╔══██╗██║   ██║██║   ██║   ██║   ██╔══╝  ██╔══██╗██╔══╝  ██║     ██║   ██║██║███╗██║
██║  ██║╚██████╔╝╚██████╔╝   ██║   ███████╗██║  ██║██║     ███████╗╚██████╔╝╚███╔███╔╝
╚═╝  ╚═╝ ╚═════╝  ╚═════╝    ╚═╝   ╚══════╝╚═╝  ╚═╝╚═╝     ╚══════╝ ╚═════╝  ╚══╝╚══╝ 
                                                                                      

`

var (
	logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
)

var (
	cfgFile string
)

var cfg config.Config

var rootCmd = &cobra.Command{
	Use:   "routerflow",
	Short: "A CLI tool to manage MikroTik routers",
	Long:  `RouterFlow CLI Tool: Automate RouterOS configurations and monitor your network`,
	Run:   func(cmd *cobra.Command, args []string) {},
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
	utils.PrinLogo(logo)
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.routerflow.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		configName := ".routerflow"
		configType := "yaml"

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Configuration file not found.")
			fmt.Println("Please run `config` command to set up the application.")
		} else {
			fmt.Printf("Error loading configuration: %v\n", err)
			os.Exit(1)
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Printf("Failed to unmarshal config: %v\n", err)
		os.Exit(1)
	}

}

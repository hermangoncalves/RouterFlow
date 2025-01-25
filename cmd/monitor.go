/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor system, connections, and bandwidt",
	Long: `Commands for monitoring system resources, active connections, and bandwidth usage on MikroTik routers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("monitor called")
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
}

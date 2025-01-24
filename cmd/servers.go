/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hermangoncalves/routerflow/pkg/ui"
	"github.com/hermangoncalves/routerflow/pkg/ui/multiselect"
	"github.com/spf13/cobra"
)

// serversCmd represents the servers command
var serversCmd = &cobra.Command{
	Use:   "servers",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var tProgram *tea.Program
		var choices []string

		for _, server := range cfg.Servers {
			choices = append(choices, fmt.Sprintf("%s - %s", server.Name, server.Host))
		}

		intialModel := multiselect.InitialModelMultiSelect(
			"Selecione um servidor",
			choices,
			fmt.Sprintf("Press %s to selecte a server or press %s to create new.\n", ui.FocusedStyle.Render("enter"), ui.FocusedStyle.Render("+")),
		)
		tProgram = tea.NewProgram(intialModel)
		if _, err := tProgram.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serversCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serversCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serversCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

package cmd

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hermangoncalves/routerflow/api"
	"github.com/spf13/cobra"
)

var duration int // Duration in seconds

// Bubble Tea model
type model struct {
	cpuUsage    string
	memoryUsage string
	uptime      string
	err         error
	done        bool
}

// Initialize the model
func (m model) Init() tea.Cmd {
	return fetchResources()
}

// Update function to handle messages
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case resourceMsg:
		if msg.err != nil {
			m.err = msg.err
			m.done = true
			return m, nil
		}

		m.cpuUsage = msg.cpuUsage
		m.memoryUsage = msg.memoryUsage
		m.uptime = msg.uptime
		return m, fetchResources()

	case tea.KeyMsg:
		// Exit if the user presses "q"
		if msg.String() == "q" {
			m.done = true
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		// Handle resizing if needed
		return m, nil
	}

	return m, nil
}

// View function to render the UI
func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n\nPress q to quit.\n", m.err)
	}

	if m.done {
		return "Monitoring completed. Press q to quit.\n"
	}

	return fmt.Sprintf(
		"System Resource Monitoring\n\n"+
			"CPU Usage: %s%%\n"+
			"Memory Usage: %s\n"+
			"Uptime: %s\n\n"+
			"Press q to quit.\n",
		m.cpuUsage,
		m.memoryUsage,
		m.uptime,
	)
}

// Command to fetch system resources
type resourceMsg struct {
	cpuUsage    string
	memoryUsage string
	uptime      string
	err         error
}

func fetchResources() tea.Cmd {
	return func() tea.Msg {
		routerosClient, err := api.NewRouterOsClient("192.168.13.1:8728", "OLan", "syspasswd")
		if err != nil {
			return resourceMsg{err: err}
		}
		defer routerosClient.Close()

		// Fetch system resources
		reply, err := routerosClient.RunCommand("/system/resource/print")
		if err != nil {
			return resourceMsg{err: err}
		}

		// Extract data
		for _, re := range reply.Re {
			return resourceMsg{
				cpuUsage:    re.Map["cpu-load"],
				memoryUsage: fmt.Sprintf("%s / %s", re.Map["free-memory"], re.Map["total-memory"]),
				uptime:      re.Map["uptime"],
				err:         nil,
			}
		}
		return resourceMsg{err: fmt.Errorf("no data received")}
	}
}

var monitorResourcesCmd = &cobra.Command{
	Use:   "resources",
	Short: "Monitor system resources (CPU, memory, uptime) dynamically",
	Run: func(cmd *cobra.Command, args []string) {
		// Run Bubble Tea program
		p := tea.NewProgram(model{})
		if err := p.Start(); err != nil {
			log.Fatalf("Error starting Bubble Tea program: %v", err)
		}
	},
}

func init() {
	monitorResourcesCmd.Flags().IntVarP(&duration, "duration", "d", 10, "Duration in seconds to run the command")
	monitorCmd.AddCommand(monitorResourcesCmd)
}

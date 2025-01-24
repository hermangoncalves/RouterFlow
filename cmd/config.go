package cmd

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hermangoncalves/routerflow/pkg/ui/textinput"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// Config represents the server configuration.
type Config struct {
	Servers []Server `yaml:"servers"`
}

// Server represents a single server's details.
type Server struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// LoadConfig loads the existing configuration from the YAML file, if it exists.
func LoadConfig() (Config, error) {
	configPath := ".routerflow.yaml"

	// Check if the file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return Config{}, nil // No file, return an empty config
	}

	// Open the file
	file, err := os.Open(configPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	// Decode the YAML
	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

// PersistConfig saves the configuration to a YAML file.
func PersistConfig(config Config) error {
	configPath := ".routerflow.yaml"
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()
	return encoder.Encode(config)
}

// collectInput runs a text input UI to collect a single field.
func collectInput(prompt string) string {
	output := &textinput.Output{}
	tProgram := tea.NewProgram(textinput.InitialTextInputModel(prompt, output))
	if _, err := tProgram.Run(); err != nil {
		log.Fatalf("Error during input collection: %v", err)
	}
	return output.Value
}

// configCmd represents the `config` command.
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set up server configuration",
	Long:  "This command allows you to configure MikroTik servers directly in the terminal.",
	Run: func(cmd *cobra.Command, args []string) {
		// Load existing configuration
		config, err := LoadConfig()
		if err != nil {
			log.Fatalf("Failed to load existing configuration: %v", err)
		}

		// Collect new server details
		for {
			name := collectInput("Enter server name:")
			host := collectInput("Enter host (IP or domain):")
			port := collectInput("Enter port (e.g., 8728):")
			username := collectInput("Enter username:")
			password := collectInput("Enter password:")

			// Add the new server to the configuration
			config.Servers = append(config.Servers, Server{
				Name:     name,
				Host:     host,
				Port:     port,
				Username: username,
				Password: password,
			})

			// Ask if the user wants to add another server
			addMore := collectInput("Do you want to add another server? (yes/no):")
			if addMore != "yes" {
				break
			}
		}

		// Save the updated configuration
		if err := PersistConfig(config); err != nil {
			log.Fatalf("Failed to save configuration: %v", err)
		}

		fmt.Println("Configuration saved successfully!")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

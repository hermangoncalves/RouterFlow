package utils

import (
	"fmt"
	"regexp"

	"github.com/charmbracelet/lipgloss"
)

var (
	logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
)

func ValidateModuleName(moduleName string) bool {
	matched, _ := regexp.Match("^[a-zA-Z0-9_-]+(?:[\\/.][a-zA-Z0-9_-]+)*$", []byte(moduleName))
	return matched
}

func PrinLogo(logo string) {
	fmt.Printf("%s\n", logoStyle.Render(logo))
}

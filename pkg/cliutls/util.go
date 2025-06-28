package cliutils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func LoadEnvFiles() {

	// Try multiple locations in order of preference
	cwd, _ := os.Getwd()
	homeDir, _ := os.UserHomeDir()

	envFiles := []string{
		filepath.Join(cwd, ".env"),
		filepath.Join(cwd, ".env.local"),
		filepath.Join(homeDir, ".env"),
		filepath.Join(homeDir, ".altalune.env"),
	}

	// Load the first available .env file
	for _, envFile := range envFiles {
		if _, err := os.Stat(envFile); err == nil {
			if err := godotenv.Load(envFile); err != nil {
				log.Warn("Could not load env file", "file", envFile, "err", err)
			} else {
				log.Debug("Loaded environment from", "file", envFile)
				break
			}
		}
	}
}

func MaskToken(token string) string {
	if len(token) > 4 {
		result := token[:4]
		result += strings.Repeat("*", len(token)-4)
		return result
	}
	return "****"
}

func Contains(s string, substrings ...string) bool {
	s = strings.ToLower(s)
	for _, substring := range substrings {
		substring = strings.ToLower(substring)
		if strings.Contains(s, substring) {
			return true
		}
	}
	return false
}

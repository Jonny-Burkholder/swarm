package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
	"github.com/jonny-burkholder/swarm/internal/models"
)

type ConfigCommand struct {
	// Flag values
	Show       bool  // --show flag
	Runs       int   // --runs flag
	Concurrent int   // --concurrent flag
	Async      bool  // --async flag
}

// NewConfigCommand creates a new config command with default values
func NewConfigCommand() *ConfigCommand {
	return &ConfigCommand{
		Show:       false,
		Runs:       -1,  // -1 means "not set"
		Concurrent: -1,  // -1 means "not set"
		Async:      false,
	}
}

// SetupFlags configures the flag set for the config command
func (c *ConfigCommand) SetupFlags(fs *flag.FlagSet) {
	fs.BoolVar(&c.Show, "show", c.Show, "Display the current configuration")
	fs.BoolVar(&c.Show, "s", c.Show, "Display the current configuration (short)")
	
	fs.IntVar(&c.Runs, "runs", c.Runs, "Set default number of runs")
	fs.IntVar(&c.Runs, "r", c.Runs, "Set default number of runs (short)")
	
	fs.IntVar(&c.Concurrent, "concurrent", c.Concurrent, "Set default number of concurrent workers")
	fs.IntVar(&c.Concurrent, "n", c.Concurrent, "Set default number of concurrent workers (short)")
	
	fs.BoolVar(&c.Async, "async", c.Async, "Set default async behavior")
	fs.BoolVar(&c.Async, "a", c.Async, "Set default async behavior (short)")
}

// Validate checks that the provided flags are valid
func (c *ConfigCommand) Validate() error {
	// Check if SWARMPATH is set
	swarmPath := os.Getenv("SWARMPATH")
	if swarmPath == "" {
		return fmt.Errorf("SWARMPATH environment variable is not set")
	}
	
	// If updating values, make sure they're positive
	if c.Runs != -1 && c.Runs <= 0 {
		return fmt.Errorf("runs must be greater than 0")
	}
	
	if c.Concurrent != -1 && c.Concurrent <= 0 {
		return fmt.Errorf("concurrent workers must be greater than 0")
	}
	
	return nil
}

// loadConfig loads the configuration from SWARMPATH
func loadConfig() (*models.Config, error) {
	swarmPath := os.Getenv("SWARMPATH")
	configPath := filepath.Join(swarmPath, "config.yaml")
	
	// Check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Return default config if file doesn't exist
		return &models.Config{
			Runs:       1,
			Concurrent: 1,
			Async:      false,
		}, nil
	}
	
	// Read the file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	
	// Parse YAML
	var config models.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}
	
	return &config, nil
}

// saveConfig saves the configuration to SWARMPATH
func saveConfig(config *models.Config) error {
	swarmPath := os.Getenv("SWARMPATH")
	configPath := filepath.Join(swarmPath, "config.yaml")
	
	// Convert config to YAML
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	
	// Write to file
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	
	return nil
}

// Run executes the config command
func (c *ConfigCommand) Run() error {
	if err := c.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	
	// Load the current config
	config, err := loadConfig()
	if err != nil {
		return err
	}
	
	// If --show flag, just display and exit
	if c.Show {
		fmt.Println("Current configuration:")
		fmt.Printf("  Runs:       %d\n", config.Runs)
		fmt.Printf("  Concurrent: %d\n", config.Concurrent)
		fmt.Printf("  Async:      %t\n", config.Async)
		return nil
	}
	
	// Otherwise, update the config with provided values
	updated := false
	
	if c.Runs != -1 {
		config.Runs = c.Runs
		updated = true
	}
	
	if c.Concurrent != -1 {
		config.Concurrent = c.Concurrent
		updated = true
	}
	
	// For async, we need to check if it was explicitly set
	// This is tricky because false is the default
	// For now, we'll only update if other flags are present
	
	if !updated {
		return fmt.Errorf("no configuration values provided to update")
	}
	
	// Save the updated config
	if err := saveConfig(config); err != nil {
		return err
	}
	
	fmt.Println("Configuration updated successfully!")
	fmt.Printf("  Runs:       %d\n", config.Runs)
	fmt.Printf("  Concurrent: %d\n", config.Concurrent)
	fmt.Printf("  Async:      %t\n", config.Async)
	
	return nil
}
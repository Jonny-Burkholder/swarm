package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jonny-burkholder/swarm/internal/models"
)

func TestNewConfigCommand(t *testing.T) {
	cmd := NewConfigCommand()
	
	if cmd.Show != false {
		t.Errorf("Expected Show to be false, got %v", cmd.Show)
	}
	
	if cmd.Runs != -1 {
		t.Errorf("Expected Runs to be -1, got %d", cmd.Runs)
	}
	
	if cmd.Concurrent != -1 {
		t.Errorf("Expected Concurrent to be -1, got %d", cmd.Concurrent)
	}
	
	if cmd.Async != false {
		t.Errorf("Expected Async to be false, got %v", cmd.Async)
	}
}

func TestConfigCommand_Validate(t *testing.T) {
	// Save original SWARMPATH and restore after test
	originalPath := os.Getenv("SWARMPATH")
	defer os.Setenv("SWARMPATH", originalPath)
	
	tests := []struct {
		name       string
		cmd        ConfigCommand
		swarmPath  string
		wantErr    bool
		errMessage string
	}{
		{
			name:      "valid show command",
			cmd:       ConfigCommand{Show: true, Runs: -1, Concurrent: -1},
			swarmPath: "/tmp/test",
			wantErr:   false,
		},
		{
			name:       "SWARMPATH not set",
			cmd:        ConfigCommand{Show: true},
			swarmPath:  "",
			wantErr:    true,
			errMessage: "SWARMPATH environment variable is not set",
		},
		{
			name:       "negative runs",
			cmd:        ConfigCommand{Runs: -5, Concurrent: 10},
			swarmPath:  "/tmp/test",
			wantErr:    true,
			errMessage: "runs must be greater than 0",
		},
		{
			name:       "zero runs",
			cmd:        ConfigCommand{Runs: 0, Concurrent: 10},
			swarmPath:  "/tmp/test",
			wantErr:    true,
			errMessage: "runs must be greater than 0",
		},
		{
			name:       "negative concurrent",
			cmd:        ConfigCommand{Runs: 10, Concurrent: -5},
			swarmPath:  "/tmp/test",
			wantErr:    true,
			errMessage: "concurrent workers must be greater than 0",
		},
		{
			name:       "zero concurrent",
			cmd:        ConfigCommand{Runs: 10, Concurrent: 0},
			swarmPath:  "/tmp/test",
			wantErr:    true,
			errMessage: "concurrent workers must be greater than 0",
		},
		{
			name:      "valid update command",
			cmd:       ConfigCommand{Runs: 100, Concurrent: 10},
			swarmPath: "/tmp/test",
			wantErr:   false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set SWARMPATH for this test
			os.Setenv("SWARMPATH", tt.swarmPath)
			
			// Run validation
			err := tt.cmd.Validate()
			
			// Check if error matches expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			// If we expect an error, check the message
			if tt.wantErr && err != nil && err.Error() != tt.errMessage {
				t.Errorf("Validate() error message = %v, want %v", err.Error(), tt.errMessage)
			}
		})
	}
}

func TestLoadConfig(t *testing.T) {
	// Save original SWARMPATH
	originalPath := os.Getenv("SWARMPATH")
	defer os.Setenv("SWARMPATH", originalPath)
	
	// Create temporary directory for testing
	tempDir := t.TempDir()
	os.Setenv("SWARMPATH", tempDir)
	
	t.Run("load config when file doesn't exist", func(t *testing.T) {
		config, err := loadConfig()
		
		if err != nil {
			t.Errorf("loadConfig() error = %v, want nil", err)
			return
		}
		
		// Should return default config
		if config.Runs != 1 {
			t.Errorf("Expected default Runs = 1, got %d", config.Runs)
		}
		if config.Concurrent != 1 {
			t.Errorf("Expected default Concurrent = 1, got %d", config.Concurrent)
		}
		if config.Async != false {
			t.Errorf("Expected default Async = false, got %v", config.Async)
		}
	})
	
	t.Run("load config when file exists", func(t *testing.T) {
		// Create a config file
		configPath := filepath.Join(tempDir, "config.yaml")
		yamlContent := []byte("runs: 50\nconcurrent: 5\nasync: true\n")
		err := os.WriteFile(configPath, yamlContent, 0644)
		if err != nil {
			t.Fatalf("Failed to create test config file: %v", err)
		}
		
		config, err := loadConfig()
		
		if err != nil {
			t.Errorf("loadConfig() error = %v, want nil", err)
			return
		}
		
		// Should return values from file
		if config.Runs != 50 {
			t.Errorf("Expected Runs = 50, got %d", config.Runs)
		}
		if config.Concurrent != 5 {
			t.Errorf("Expected Concurrent = 5, got %d", config.Concurrent)
		}
		if config.Async != true {
			t.Errorf("Expected Async = true, got %v", config.Async)
		}
	})
	
	t.Run("load config with invalid YAML", func(t *testing.T) {
		// Create invalid YAML
		configPath := filepath.Join(tempDir, "config.yaml")
		invalidYAML := []byte("this is not: valid: yaml:")
		err := os.WriteFile(configPath, invalidYAML, 0644)
		if err != nil {
			t.Fatalf("Failed to create test config file: %v", err)
		}
		
		_, err = loadConfig()
		
		if err == nil {
			t.Error("loadConfig() should error with invalid YAML")
		}
	})
}

func TestSaveConfig(t *testing.T) {
	// Save original SWARMPATH
	originalPath := os.Getenv("SWARMPATH")
	defer os.Setenv("SWARMPATH", originalPath)
	
	// Create temporary directory for testing
	tempDir := t.TempDir()
	os.Setenv("SWARMPATH", tempDir)
	
	t.Run("save config successfully", func(t *testing.T) {
		config := &models.Config{
			Runs:       100,
			Concurrent: 20,
			Async:      true,
		}
		
		err := saveConfig(config)
		
		if err != nil {
			t.Errorf("saveConfig() error = %v, want nil", err)
			return
		}
		
		// Verify file was created
		configPath := filepath.Join(tempDir, "config.yaml")
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			t.Error("Config file was not created")
		}
		
		// Load it back and verify values
		loadedConfig, err := loadConfig()
		if err != nil {
			t.Errorf("Failed to load saved config: %v", err)
			return
		}
		
		if loadedConfig.Runs != 100 {
			t.Errorf("Expected Runs = 100, got %d", loadedConfig.Runs)
		}
		if loadedConfig.Concurrent != 20 {
			t.Errorf("Expected Concurrent = 20, got %d", loadedConfig.Concurrent)
		}
		if loadedConfig.Async != true {
			t.Errorf("Expected Async = true, got %v", loadedConfig.Async)
		}
	})
}
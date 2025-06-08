package compare

import (
	"flag"
	"fmt"
)

type CompareCommand struct {
	// Flag values
	Out    string
	Format string
}

// NewCompareCommand creates a new compare command with default values
func NewCompareCommand() *CompareCommand {
	return &CompareCommand{
		Out:    "stdout",
		Format: "html",
	}
}

// SetupFlags configures the flag set for the compare command
func (c *CompareCommand) SetupFlags(fs *flag.FlagSet) {
	fs.StringVar(&c.Out, "out", c.Out, "Output destination (stdout, file path)")
	fs.StringVar(&c.Out, "o", c.Out, "Output destination (short)")

	fs.StringVar(&c.Format, "format", c.Format, "Output format (html, json, csv)")
	fs.StringVar(&c.Format, "f", c.Format, "Output format (short)")
}

// Validate checks that the provided flags are valid
func (c *CompareCommand) Validate() error {
	validFormats := map[string]bool{
		"html": true,
		"json": true,
		"csv":  true,
	}

	if !validFormats[c.Format] {
		return fmt.Errorf("invalid format '%s', must be one of: html, json, csv", c.Format)
	}

	return nil
}

// Run executes the compare command
func (c *CompareCommand) Run(args []string) error {
	if err := c.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	if len(args) < 2 {
		return fmt.Errorf("compare requires at least 2 benchmark result files")
	}

	// TODO: Implement the actual compare logic
	fmt.Printf("Comparing benchmark results with:\n")
	fmt.Printf("  Format: %s\n", c.Format)
	fmt.Printf("  Output: %s\n", c.Out)
	fmt.Printf("  Files: %v\n", args)

	return nil
}

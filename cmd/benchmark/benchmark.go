package benchmark

import (
	"flag"
	"fmt"
	"time"

	"github.com/jonny-burkholder/swarm/internal/logger"
)

type BenchmarkCommand struct {
	Runner Runner
	Logger logger.Logger

	// Flag values
	Collection string
	Config     string
	LogLevel   string
	Runs       int
	Concurrent int
	Duration   time.Duration
	Async      bool
	Save       bool
	Out        string
}

// NewBenchmarkCommand creates a new benchmark command with default values
func NewBenchmarkCommand() *BenchmarkCommand {
	return &BenchmarkCommand{
		// Set sensible defaults
		LogLevel:   "info",
		Runs:       1,
		Concurrent: 1,
		Duration:   0, // 0 means use runs instead of duration
		Async:      false,
		Save:       false,
		Out:        "stdout",
	}
}

// SetupFlags configures the flag set for the benchmark command
func (b *BenchmarkCommand) SetupFlags(fs *flag.FlagSet) {
	// Required flags
	fs.StringVar(&b.Collection, "collection", b.Collection, "Collection file to run benchmarks against")
	fs.StringVar(&b.Collection, "c", b.Collection, "Collection file to run benchmarks against (short)")

	fs.StringVar(&b.Config, "config", b.Config, "Configuration file for the benchmark")
	fs.StringVar(&b.Config, "f", b.Config, "Configuration file for the benchmark (short)")

	// Performance flags
	fs.IntVar(&b.Runs, "runs", b.Runs, "Number of runs to execute")
	fs.IntVar(&b.Runs, "r", b.Runs, "Number of runs to execute (short)")

	fs.IntVar(&b.Concurrent, "concurrent", b.Concurrent, "Number of concurrent workers")
	fs.IntVar(&b.Concurrent, "n", b.Concurrent, "Number of concurrent workers (short)")

	fs.DurationVar(&b.Duration, "duration", b.Duration, "Duration to run tests (e.g., 30s, 5m). If set, overrides --runs")
	fs.DurationVar(&b.Duration, "d", b.Duration, "Duration to run tests (short)")

	fs.BoolVar(&b.Async, "async", b.Async, "Run requests asynchronously within each worker")
	fs.BoolVar(&b.Async, "a", b.Async, "Run requests asynchronously within each worker (short)")

	// Output flags
	fs.StringVar(&b.LogLevel, "log-level", b.LogLevel, "Log level (debug, info, warn, error)")
	fs.StringVar(&b.LogLevel, "l", b.LogLevel, "Log level (short)")

	fs.BoolVar(&b.Save, "save", b.Save, "Save benchmark results to disk")
	fs.BoolVar(&b.Save, "s", b.Save, "Save benchmark results to disk (short)")

	fs.StringVar(&b.Out, "out", b.Out, "Output destination (stdout, file path, or format)")
	fs.StringVar(&b.Out, "o", b.Out, "Output destination (short)")
}

// Validate checks that the provided flags are valid
func (b *BenchmarkCommand) Validate() error {
	if b.Collection == "" {
		return fmt.Errorf("collection file is required (use -c or --collection)")
	}

	if b.Runs <= 0 && b.Duration <= 0 {
		return fmt.Errorf("either --runs or --duration must be specified and greater than 0")
	}

	if b.Concurrent <= 0 {
		return fmt.Errorf("concurrent workers must be greater than 0")
	}

	validLogLevels := map[string]bool{
		"debug": true,
		"info":  true,
		"warn":  true,
		"error": true,
	}

	if !validLogLevels[b.LogLevel] {
		return fmt.Errorf("invalid log level '%s', must be one of: debug, info, warn, error", b.LogLevel)
	}

	return nil
}

func (b *BenchmarkCommand) Run() error {
	if err := b.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// TODO: Implement the actual benchmark logic
	fmt.Printf("Running benchmark with:\n")
	fmt.Printf("  Collection: %s\n", b.Collection)
	fmt.Printf("  Config: %s\n", b.Config)
	fmt.Printf("  Runs: %d\n", b.Runs)
	fmt.Printf("  Duration: %v\n", b.Duration)
	fmt.Printf("  Concurrent: %d\n", b.Concurrent)
	fmt.Printf("  Async: %t\n", b.Async)
	fmt.Printf("  Save: %t\n", b.Save)
	fmt.Printf("  Output: %s\n", b.Out)
	fmt.Printf("  Log Level: %s\n", b.LogLevel)

	return nil
}

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jonny-burkholder/swarm/cmd/benchmark"
	"github.com/jonny-burkholder/swarm/cmd/compare"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Global flags (applied to all commands)
	var verbose bool
	var quiet bool

	// Parse global flags before subcommand
	globalFlags := flag.NewFlagSet("swarm", flag.ExitOnError)
	globalFlags.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	globalFlags.BoolVar(&verbose, "v", false, "Enable verbose output (short)")
	globalFlags.BoolVar(&quiet, "quiet", false, "Suppress all output except errors")
	globalFlags.BoolVar(&quiet, "q", false, "Suppress all output except errors (short)")

	// Get the subcommand
	subcommand := os.Args[1]

	// Parse the remaining arguments based on subcommand
	var err error
	switch subcommand {
	case "benchmark", "bench":
		err = runBenchmark(os.Args[2:], verbose, quiet)
	case "compare", "comp":
		err = runCompare(os.Args[2:], verbose, quiet)
	case "help", "-h", "--help":
		printUsage()
		return
	case "version", "-v", "--version":
		printVersion()
		return
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", subcommand)
		printUsage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func runBenchmark(args []string, verbose, quiet bool) error {
	cmd := benchmark.NewBenchmarkCommand()

	// Create flag set for benchmark command
	fs := flag.NewFlagSet("benchmark", flag.ExitOnError)
	cmd.SetupFlags(fs)

	// Add global flags to the command flag set
	fs.BoolVar(&verbose, "verbose", verbose, "Enable verbose output")
	fs.BoolVar(&verbose, "v", verbose, "Enable verbose output (short)")
	fs.BoolVar(&quiet, "quiet", quiet, "Suppress all output except errors")
	fs.BoolVar(&quiet, "q", quiet, "Suppress all output except errors (short)")

	// Parse flags
	if err := fs.Parse(args); err != nil {
		return err
	}

	// Apply global flags
	if verbose && cmd.LogLevel == "info" {
		cmd.LogLevel = "debug"
	}
	if quiet {
		cmd.LogLevel = "error"
	}

	return cmd.Run()
}

func runCompare(args []string, verbose, quiet bool) error {
	cmd := compare.NewCompareCommand()

	// Create flag set for compare command
	fs := flag.NewFlagSet("compare", flag.ExitOnError)
	cmd.SetupFlags(fs)

	// Add global flags to the command flag set
	fs.BoolVar(&verbose, "verbose", verbose, "Enable verbose output")
	fs.BoolVar(&verbose, "v", verbose, "Enable verbose output (short)")
	fs.BoolVar(&quiet, "quiet", quiet, "Suppress all output except errors")
	fs.BoolVar(&quiet, "q", quiet, "Suppress all output except errors (short)")

	// Parse flags
	if err := fs.Parse(args); err != nil {
		return err
	}

	// Get remaining arguments (benchmark files to compare)
	remainingArgs := fs.Args()

	return cmd.Run(remainingArgs)
}

func printUsage() {
	fmt.Println("swarm - The ultimate API testing and benchmarking tool")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  swarm <command> [flags] [args]")
	fmt.Println()
	fmt.Println("Available Commands:")
	fmt.Println("  benchmark, bench    Run API benchmarks")
	fmt.Println("  compare, comp       Compare benchmark results")
	fmt.Println("  help               Show this help message")
	fmt.Println("  version            Show version information")
	fmt.Println()
	fmt.Println("Global Flags:")
	fmt.Println("  -v, --verbose      Enable verbose output")
	fmt.Println("  -q, --quiet        Suppress all output except errors")
	fmt.Println()
	fmt.Println("Use 'swarm <command> --help' for more information about a command.")
}

func printVersion() {
	fmt.Println("swarm version 0.1.0")
}

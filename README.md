# swarm
The ultimate API testing tool

NOTE: some of this functionality is still in progress. If you are reading this at the time of install, not all SWARM features may be available.

## Installation

### Linux

Run go install:

``` bash
go install https://github.com/Jonny-Burkholder/swarm
```

Set the `SWARMPATH` environment variable on your system:

```bash
export SWARMPATH=usr/local/swarm
```

Your swarmpath can be any directory you choose. This is the default location swarm will look for things like config and test collections. Alternatively, you can pass those values in as flags.

## Usage

The main function of swarm is `test`. If you have your SWARMPATH set up, swarm will first look there for a default test suite and config. If no config is set, swarm will use the default config settings. If no test suite is selected, swarm will error and print usage instructions.

For additional information, run `swarm --help`, or for subcommand usage run `swarm {subcommand} --help`

### Configuration Management

SWARM uses a configuration file to store default values for benchmark runs. This allows you to set preferences once and use them across all your test runs.

#### Viewing Current Configuration

To view your current default configuration:

```bash
swarm config --show
```

**Example Output:**
```
Current configuration:
  Runs:       1
  Concurrent: 1
  Async:      false
```

#### Updating Configuration

Update your default configuration using flags:

```bash
swarm config --runs 100 --concurrent 10
```

**Example Output:**
```
Configuration updated successfully!
  Runs:       100
  Concurrent: 10
  Async:      false
```

#### Configuration Options

- `--runs` or `-r`: Set the default number of benchmark runs
- `--concurrent` or `-n`: Set the default number of concurrent workers
- `--async` or `-a`: Set the default async behavior
- `--show` or `-s`: Display the current configuration

#### Configuration Storage

Configuration is stored as a YAML file at `$SWARMPATH/config.yaml`. You can also edit this file directly:

```yaml
runs: 100
concurrent: 10
async: false
```

**Note:** The `SWARMPATH` environment variable must be set for configuration management to work. If not set, swarm will return an error message

### Available Commands

SWARM provides several commands for different testing and benchmarking needs:

#### `swarm benchmark` (or `swarm bench`)
Run API benchmarks against your test collections.

**Example:**
```bash
swarm benchmark --collection tests.yml --runs 100 --concurrent 10
```

#### `swarm compare` (or `swarm comp`)
Compare results from different benchmark runs.

**Example:**
```bash
swarm compare result1.json result2.json
```

#### `swarm config`
View or update default configuration settings.

**Examples:**
```bash
# View current config
swarm config --show

# Update defaults
swarm config --runs 100 --concurrent 10
```

#### `swarm help`
Display help information about available commands.

#### `swarm version`
Display the current version of swarm.

### API Testing

SWARM exists to test your API in every conceivable way! Made primarily for load testing, SWARM is also more than capable of running functional test suites, replacing cumbersome software with simple YAML suites

### Benchmarking

In addition to sending requests and receiving responses, SWARM collects benchmarking data to see where your API weak and strong points are. Couple this with robuts profiling to optimize your APIs.

Save runs to compare api versions

### Compare

Create and server html charts comparing different run

### Config

Manage configuration

### Easy test suites

Instead of hand-writing test runs or using curl (nothing wrong with curl!), define simple YAML test suites and run them in one line from the terminal

## Looking for contributors!

Development of open-source software is hard, especially when we all have day jobs. We do it because we love free tech and sharing knowledge. If you like this project idea and would like to help, please reach out!

## Philosophy

I am a human, and I make software for other humans. Despite having a copilot.md file, not a drop of this code is AI generated, and I'd like to keep it that way. Hand-writing code keeps me sharp, and it helps to ensure that code is actually functional, and not just plausible. LLM-generated pull requests will be rejected and the offending contributor will be banned. Feel free to use AI to help you understand the code and architecture, but please write the code yourself if you're looking to contribute

## Contributing

For more information on contributing, please refer to the [contribution guidelines](contribution_guidelines.md)
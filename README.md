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

The main function of swarm is `test`. If you have your SWARMPATH set up, swarm will first look there for a default test suite and config. If no config is set, swarm will use the default config settings (print your config settings with `swarm config`. If no test suite is selected, swarm will error and print usage instructions.

For additional information, run `swarm --help`, or for subcommand usage run `swarm {subcommand} --help`

### API Testing

SWARM exists to test your API in every conceivable way! Made primarily for load testing, SWARM is also more than capable of running functional test suites, replacing cumbersome software with simple YAML suites

### Benchmarking

In addition to sending requests and receiving responses, SWARM collects benchmarking data to see where your API weak and strong points are. Couple this with robuts profiling to optimize your APIs.

Save runs to compare api versions

### Compare

Create and server html charts comparing different run

### Easy test suites

Instead of hand-writing test runs or using curl (nothing wrong with curl!), define simple YAML test suites and run them in one line from the terminal

## Looking for contributors!

Development of open-source software is hard, especially when we all have day jobs. We do it because we love free tech and sharing knowledge. If you like this project idea and would like to help, please reach out!

## Philosophy

I am a human, and I make software for other humans. Despite having a copilot.md file, not a drop of this code is AI generated, and I'd like to keep it that way. Hand-writing code keeps me sharp, and it helps to ensure that code is actually functional, and not just plausible. LLM-generated pull requests will be rejected and the offending contributor will be banned. Feel free to use AI to help you understand the code and architecture, but please write the code yourself if you're looking to contribute

## Contributing

For more information on contributing, please refer to the [contribution guidelines](contribution_guidelines.md)
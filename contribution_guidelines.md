# SWARM Contributions Guidelines

So you're interested in contributing? Great! Glad to have you aboard. We want to work together to create a great tool.

## Solving Issues

1. Planned work for the project is kept in the form of github issues (look on the "issues" tab). If you see an issue that you'd like to tackle, ask to be assigned to it, and I'll assign out tasks as I can.
1. Once an issue is assigned to you, create a branch with the issue number (e.g. 42-the-meaning-of-life)
1. Once you've written and committed some code, open a pull request. If the code is still in progress, please open the PR as a draft
1. I'll try to get to code reviews as I can. If you feel I'm being too slow, give me a nudge by requesting a review on the PR
1. All checks should pass before a PR is merged into main, to help ensure best code quality
1. Please write all the code yourself! LLMs are great tools for many things, and feel free to use them for help with documentation and config like YAML files (but don't go overboard, they tend to write *way* more documentation than is actually necessary or helpful). In this project, we prefer code that is written by humans. This goes double for tests, which are talked more about in the next section.

## Writing tests

We want all code that can be unit tested to be unit tested. Tests should be clearly named and table-driven. E.G. There should be a slice of test values that contain input and expected result. For more information on my preferences for writing tests, please refer to my [golang test tutorial](https://github.com/Jonny-Burkholder/go-test-tutorial)

## Local development

In order to not run over on github actions usage, please try to run `go test ./...` and `golangci-lint run` locally to make sure there are no issues before pushing to an open PR. Follow the [instructions](https://golangci-lint.run/docs/welcome/install/local/) for installing golangci-lint

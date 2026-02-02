# General Instructions

- if a prompt is unclear, ask for clarification before proceeding

- focus on one task at a time, completing the task before moving to the next one

- share a detailed plan before starting any implementation

- always provide a summary of changes made after completing a task

- you are a collaborator. bounce ideas when it seems appropriate

- always ask for user confirmation before making changes

- use git in @terminal to find out what changes were made recently

- use line numbers to refer to specific parts of the code

- do not leave unfinished code in implementations

- explain any todos left in your implementations


# Project Context

- project name: swarm

- swarm is a load testing cli that runs benchmarks and creates comparative reports

- commands live in `cmd/`

- the entrypoint to swarm is `main.go`

# Code Quality

- Write idiomatic Go code that is clean, efficient, and easy to read


# Style and Best Practices

- This project uses the latest version of Go (1.24). Use the latest stylings and modernizations

- Do not, under any circumstances, use stdlib packages that have been deprecated

- Follow the Go Proverbs [^1]

- Follow the Effective Go guidelines [^2]

- Avoid reflection and complex concurrency where possible

- Use clear names and clarifying comments. Create or update documentation when necessary


# TODOs and Technical Debt

- Use clear, actionable `TODO` or `FIXME` comments in code for small, contextual improvements or technical debt.
- Format:  
  `// TODO(username): Description. LIN-123`  
  Where `LIN-123` is the Linear issue ID, if available.
- For larger or cross-cutting technical debt, create a Linear issue and reference it in the code comment.
- Optionally, include a direct link to the Linear issue for clarity.
- Regularly review and triage TODOs and tech debt during planning or code review.
- Major or recurring tech debt can be summarized in a `TECH_DEBT.md` or in the `README.md`, with links to Linear issues.

# References

[1] https://go-proverbs.github.io/
[2] https://go.dev/doc/effective_go
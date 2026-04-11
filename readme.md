# Commitment Issues

A command-line tool that audits a Git repository's commit history and assigns a humorous personality type to the project's maintainer. It analyzes commit frequency, message styles, and timestamps to generate a profile.

## Installation

### Prerequisites

- Go 1.21 or higher
- Git

### Build from source

```bash
git clone https://github.com/olia-software/commitment-issues.git
cd commitment-issues
go mod tidy
go build -o commitment-issues
```

Usage

Run the tool inside any Git repository:

```bash
./commitment-issues
```

Specify a custom repository path:

```bash
./commitment-issues -path /path/to/your/repo
```

How It Works

The tool reads the entire commit history of the current branch and calculates several metrics:

· Total number of commits
· Commit types (feat, fix, docs, style, refactor, test, chore, other)
· Night commits (between 22:00 and 05:00)
· Weekend commits (Saturday and Sunday)
· Conventional Commits compliance rate
· Average commit message length
· Commit interval consistency (standard deviation of time between commits)

Based on these metrics, a personality profile is assigned to the main author of the repository.

Personality Types

Type Description
Documentation Prophet Writes extensive documentation and treats READMEs as sacred texts.
The Midnight Committer Does most work late at night, fueled by caffeine and darkness.
Rogue Feature Implementer Prioritizes new features over bug fixes. Tests are optional.
Guardian of the Main Branch Maintains pristine commit messages and disciplined rebasing.
The Weekend Warrior Commits heavily on Saturdays and Sundays.
The Agile Devotee Balanced, predictable commit pattern that follows standard workflows.

Example Output

```
======================================
  Commitment Issues - Maintainer Profile
======================================
  Guardian of the Main Branch
--------------------------------------
  Your commit messages are pristine, your rebases surgical. The main branch has never been safer.
--------------------------------------
  Key Stats:
    Total commits: 142
    Top author:   john@example.com
    Commit types: chore:12, docs:23, feat:45, fix:38, refactor:15, test:9
    Night commits:  8.5%
    Weekend ratio:  12.0%
    Conventional:   76.1%
    Avg msg length: 68 chars
--------------------------------------
  Personality Traits:
    - Disciplined
    - Methodical
    - Protective
======================================
```

Development

Dependencies

The project uses go-git/v5 for Git repository interaction:

```bash
go get github.com/go-git/go-git/v5
```

Project Structure

```
commitment-issues/
├── go.mod
├── main.go           # CLI entry point
├── stats.go          # Commit collection and analysis
└── personality.go    # Profile generation and output
```

Running tests

```bash
go test ./...
```

License

MIT License. See LICENSE file for details.

Disclaimer

This tool is intended for entertainment and educational purposes. Personality profiles are generated algorithmically and should not be taken as serious psychological assessments.
# Todo CLI Application

## Overview
A command-line todo application written in Go that allows managing tasks with CRUD operations. The application is organized in the `todo_package` package.

## Installation
1. Ensure you have Go 1.16+ installed
2. Clone/download the repository
3. Build the application:
```bash
make build  # Creates executable named 'todo'
```

## Usage
Run commands using the built executable:
```bash
./todo [command] [arguments]
```

### Available Commands
| Command | Description | Example |
|---------|-------------|---------|
| `--add "title"` | Add new task | `./todo --add "Buy milk"` |
| `--edit "id:new_title"` | Edit existing task | `./todo --edit "1:Buy organic milk"` |
| `--delete id` | Delete task by ID | `./todo --delete 1` |
| `--toggle id` | Toggle task completion | `./todo --toggle 1` |
| `--list` | Show all tasks | `./todo --list` |

## Development
The project includes a Makefile with useful targets:

```bash
make format    # Format all Go code
make validate # Run vet checks
make build    # Build executable
make clean    # Remove build artifacts
```

## Implementation Details
The application uses:
- Standard library packages: `flag`, `fmt`, `os`, `strconv`, `strings`
- `CmdFlags` struct to handle command line arguments
- `Execute()` method to route commands to appropriate operations

The main components are:
- `commands.go`: Contains all CLI command handling logic
- `Makefile`: Provides build automation

## Example Workflow
1. Add tasks:
```bash
./todo --add "Task 1"
./todo --add "Task 2"
```

2. List tasks:
```bash
./todo --list
```

3. Edit task 1:
```bash
./todo --edit "1:Updated Task 1"
```

4. Toggle completion for task 2:
```bash
./todo --toggle 2
```

5. Delete task 1:
```bash
./todo --delete 1

# Log File Search Tool

## Overview
A command-line tool for searching log files with timestamp filtering capabilities. The application searches for keywords in log files and provides flexible filtering options.

## Features
- 🔍 Keyword search in log files
- ⏱️ Timestamp range filtering
- 📏 Head/tail line limiting
- 📊 Formatted output display
- 🛠️ Case-insensitive search

## Installation
1. Ensure you have Go 1.16+ installed
2. Clone/download the repository
3. Build the application:
```bash
make build  # Creates executable named 'finder'
```

## Usage
```bash
./finder [required flags] [optional filters]
```

### Required Flags
| Flag | Description | Example |
|------|-------------|---------|
| `--filepath` | Path to log file | `--filepath "logs.txt"` |
| `--keyword`  | Search term | `--keyword "Error"` |
| `--print`    | Display results | `--print` |

### Optional Filters
| Flag | Description | Format | Example |
|------|-------------|--------|---------|
| `--from` | Start timestamp | "YYYY-MM-DD HH:MM:SS" | `--from "2025-06-14 14:00:00"` |
| `--to`   | End timestamp | "YYYY-MM-DD HH:MM:SS" | `--to "2025-06-14 16:00:00"` |
| `--head` | Show first N matches | Number | `--head 5` |
| `--tail` | Show last N matches | Number | `--tail 10` |

## Examples

### Basic Search
```bash
./finder --filepath logs.txt --keyword "Error" --print
```

### Time Range Search
```bash
./finder --filepath logs.txt --keyword "Warning" \
  --from "2025-06-14 14:00:00" \
  --to "2025-06-14 16:00:00" \
  --print
```

### Get First/Last Matches
```bash
# First 3 matches
./finder --filepath logs.txt --keyword "Info" --head 3 --print

# Last 5 matches
./finder --filepath logs.txt --keyword "Debug" --tail 5 --print
```

## Development
```bash
make format    # Format all Go code
make validate  # Run vet checks
make run       # Run without building
make clean     # Remove build artifacts
```

## Implementation
- **commands.go**: Handles CLI flag parsing and validation
- **finder.go**: Contains core search and filtering logic
- **main.go**: Application entry point

The tool uses standard Go libraries including:
- `bufio` for file scanning
- `time` for timestamp parsing
- `flag` for command-line arguments

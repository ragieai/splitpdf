# PDF Splitter

A simple Go tool to split PDF files into multiple smaller PDFs with a specified number of pages per file.

## Installation

### Option 1: Install with Go

If you have Go installed, you can install directly:

```bash
go install github.com/ragieai/splitpdf/cmd/splitpdf@latest
```

This will install the `splitpdf` binary to your `$GOPATH/bin` directory, which should be in your PATH.

### Option 2: Build from source

Compile the tool:

```bash
# Clone the repository
git clone https://github.com/ragieai/splitpdf.git
cd splitpdf

# Resolve dependencies
go mod tidy

# Build the executable
go build -o splitpdf cmd/splitpdf/main.go
```

## Troubleshooting

If you encounter build errors related to missing dependencies, run:

```bash
go mod tidy
```

This will download the required dependencies and update the go.sum file.

## Usage

```bash
splitpdf [path/to/input.pdf] [flags]
```

### Flags

- `--output-dir`: Directory to save the output files (defaults to current directory)
- `--pages`: Number of pages per split (default: 500)
- `--help`, `-h`: Display help information

### Examples

Split a PDF into chunks of 100 pages each:

```bash
splitpdf large_document.pdf --pages 100
```

Save the output files to a specific directory:

```bash
splitpdf large_document.pdf --output-dir ./split_output --pages 200
```

The command supports flexible flag positioning:

```bash
# Flags before the argument
splitpdf --pages 100 large_document.pdf

# Flags after the argument
splitpdf large_document.pdf --pages 100
```

## Dependencies

This tool uses the following libraries:
- [pdfcpu](https://github.com/pdfcpu/pdfcpu) for PDF manipulation
- [cobra](https://github.com/spf13/cobra) for command-line interface 
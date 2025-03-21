package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/spf13/cobra"
)

var (
	outputDir string
	pages     int
)

func splitPDF(inputFile string, outputDir string, pagesPerSplit int) error {
	// If output directory is not specified, use the current directory
	if outputDir == "" {
		var err error
		outputDir, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("error getting current directory: %w", err)
		}
	}

	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	// Extract the base name of the input file (without .pdf extension)
	baseName := filepath.Base(inputFile)
	if len(baseName) > 4 && filepath.Ext(baseName) == ".pdf" {
		baseName = baseName[:len(baseName)-4]
	}

	// Get page count of the PDF
	ctx, err := api.ReadContextFile(inputFile)
	if err != nil {
		return fmt.Errorf("error reading PDF file: %w", err)
	}

	totalPages := ctx.PageCount
	partNumber := 1

	for startPage := 1; startPage <= totalPages; startPage += pagesPerSplit {
		// Determine end page for this split
		endPage := startPage + pagesPerSplit - 1
		if endPage > totalPages {
			endPage = totalPages
		}

		// Create page selection string for pdfcpu (as a slice of strings)
		pageSelection := []string{fmt.Sprintf("%d-%d", startPage, endPage)}

		// Output filename format: NAME-part0N.pdf
		outputFilename := fmt.Sprintf("%s-part%02d.pdf", baseName, partNumber)
		outputPath := filepath.Join(outputDir, outputFilename)

		// Configure the split
		conf := model.NewDefaultConfiguration()

		// Extract pages to new PDF
		if err := api.ExtractPagesFile(inputFile, outputPath, pageSelection, conf); err != nil {
			return fmt.Errorf("error extracting pages %v to %s: %w", pageSelection, outputPath, err)
		}

		fmt.Printf("Created %s with pages %d to %d\n", outputPath, startPage, endPage)
		partNumber++
	}

	return nil
}

func main() {
	// Create the root command
	rootCmd := &cobra.Command{
		Use:   "splitpdf [path/to/input.pdf]",
		Short: "Split a PDF into multiple PDFs with a fixed number of pages",
		Long:  `A tool that splits a PDF file into multiple PDFs with a specified number of pages per split.`,
		Args:  cobra.ExactArgs(1), // Require exactly one argument (the input file)
		Run: func(cmd *cobra.Command, args []string) {
			inputFile := args[0]

			err := splitPDF(inputFile, outputDir, pages)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		},
	}

	// Define command-line flags
	rootCmd.Flags().StringVar(&outputDir, "output-dir", "", "Directory to save the output files (defaults to current directory)")
	rootCmd.Flags().IntVar(&pages, "pages", 500, "Number of pages per split (default: 500)")

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

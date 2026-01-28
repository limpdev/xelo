/*
Created © 2026 Limp Cheney limpdev@proton.me
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/limpdev/xelo/internal/cleaner"

	"github.com/spf13/cobra"
)

var (
	inputFile       string
	cleanOutputFile string
	keepHeaders     []string
	dryRun          bool
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean and reformat Excel files",
	Long: `Clean and reformat Excel files by:
- Removing empty columns (except those with specific headers)
- Reformatting date columns to dd-mmm-yyyy format
- Saving the cleaned file to a new location

Examples:
  # Clean a file, keeping columns with specific headers
  xelo clean -i input.xlsx -o cleaned.xlsx --keep "Employee_Id,Stockholder_Id"

  # Preview changes without saving
  xelo clean -i input.xlsx --dry-run

  # Clean with default settings
  xelo clean -i input.xlsx -o output.xlsx`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runClean(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	persistentHeaders := []string{"As_Of_Date", "Internal_Participant_Id", "External_Participant_Id"}
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input Excel file path (required)")
	cleanCmd.Flags().StringVarP(&cleanOutputFile, "output", "o", "", "Output file path (required unless --dry-run)")
	cleanCmd.Flags().StringSliceVar(&keepHeaders, "keep", persistentHeaders, "Headers to keep even if column is empty (comma-separated)")
	cleanCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview changes without saving")

	cleanCmd.MarkFlagRequired("input")
}

func runClean() error {
	// Validate flags
	if !dryRun && cleanOutputFile == "" {
		return fmt.Errorf("--output is required unless using --dry-run")
	}

	if inputFile == "" {
		return fmt.Errorf("--input is required")
	}

	// Check if input file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist: %s", inputFile)
	}

	fmt.Printf("Cleaning Excel file...\n")
	fmt.Printf("  Input: %s\n", inputFile)

	if dryRun {
		fmt.Printf("  Mode: DRY RUN (no file will be saved)\n")
	} else {
		fmt.Printf("  Output: %s\n", cleanOutputFile)
	}

	if len(keepHeaders) > 0 {
		fmt.Printf("  Keeping headers: %v\n", keepHeaders)
	}

	// Create cleaner config
	config := cleaner.CleanConfig{
		InputFile:   inputFile,
		OutputFile:  cleanOutputFile,
		KeepHeaders: keepHeaders,
		DryRun:      dryRun,
	}

	// Run the cleaner
	result, err := cleaner.CleanFile(config)
	if err != nil {
		return fmt.Errorf("failed to clean file: %w", err)
	}

	// Print summary
	fmt.Printf("\nCleaning Summary:\n")
	fmt.Printf("  Columns removed: %d\n", result.ColumnsRemoved)
	fmt.Printf("  Date columns reformatted: %d\n", result.DateColumnsReformatted)

	if len(result.RemovedColumns) > 0 {
		fmt.Printf("\nRemoved columns:\n")
		for _, col := range result.RemovedColumns {
			fmt.Printf("  - %s\n", col)
		}
	}

	if len(result.DateColumns) > 0 {
		fmt.Printf("\nReformatted date columns:\n")
		for _, col := range result.DateColumns {
			fmt.Printf("  - %s\n", col)
		}
	}

	if !dryRun {
		fmt.Printf("\n✓ Successfully cleaned and saved to: %s\n", cleanOutputFile)
	} else {
		fmt.Printf("\n✓ Dry run complete (no file saved)\n")
	}

	return nil
}

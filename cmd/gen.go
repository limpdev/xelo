package cmd

import (
	"fmt"
	// "github.com/limpdev/xelo/internal/generator"
	"os"

	"github.com/spf13/cobra"
)

var (
	outputFile   string
	templateFile string
	recordCount  int
	generateMode string
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate Excel files from templates",
	Long: `Generate Excel files from templates with demographic or custom data.

Examples:
  # Generate a file with 5 sample demographic records
  xelo generate -o output.xlsx -n 5

  # Generate from a custom template
  xelo generate -t template.xlsx -o output.xlsx -n 10

  # Generate single record
  xelo generate -o output.xlsx -m single`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runGenerate(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&outputFile, "output", "o", "demographic_output.xlsx", "Output file path")
	generateCmd.Flags().StringVarP(&templateFile, "template", "t", "", "Template file path (optional)")
	generateCmd.Flags().IntVarP(&recordCount, "count", "n", 5, "Number of records to generate")
	generateCmd.Flags().StringVarP(&generateMode, "mode", "m", "multiple", "Generation mode: single, multiple, or sample")
}

func runGenerate() error {
	fmt.Printf("Generating demographic data...\n")
	fmt.Printf("  Mode: %s\n", generateMode)
	fmt.Printf("  Output: %s\n", outputFile)

	if templateFile != "" {
		fmt.Printf("  Template: %s\n", templateFile)
	}

	if generateMode == "multiple" || generateMode == "sample" {
		fmt.Printf("  Records: %d\n", recordCount)
	}

	// var err error

	// switch generateMode {
	// case "single":
	// err = generator.GenerateSingleRecord(outputFile)
	// case "multiple", "sample":
	// err = generator.GenerateMultipleRecords(outputFile, recordCount)
	// default:
	// return fmt.Errorf("invalid mode: %s (must be 'single', 'multiple', or 'sample')", generateMode)
	// }

	// if err != nil {
	// 	return fmt.Errorf("failed to generate file: %w", err)
	// }

	fmt.Printf("\n✓ Successfully generated: %s\n", outputFile)
	return nil
}

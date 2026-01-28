package cleaner

import (
	"fmt"
	"strings"
	"time"

	"github.com/tealeg/xlsx/v3"
)

// CleanConfig contains configuration for cleaning an Excel file
type CleanConfig struct {
	InputFile   string
	OutputFile  string
	KeepHeaders []string
	DryRun      bool
}

// CleanResult contains information about the cleaning operation
type CleanResult struct {
	ColumnsRemoved         int
	DateColumnsReformatted int
	RemovedColumns         []string
	DateColumns            []string
}

// CleanFile processes an Excel file according to the configuration
func CleanFile(config CleanConfig) (*CleanResult, error) {
	// Open the workbook
	wb, err := xlsx.OpenFile(config.InputFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	result := &CleanResult{
		RemovedColumns: []string{},
		DateColumns:    []string{},
	}

	// Process each sheet
	for _, sheet := range wb.Sheets {
		if err := processSheet(sheet, config, result); err != nil {
			return nil, fmt.Errorf("failed to process sheet '%s': %w", sheet.Name, err)
		}
	}

	// Save if not dry run
	if !config.DryRun {
		if err := wb.Save(config.OutputFile); err != nil {
			return nil, fmt.Errorf("failed to save file: %w", err)
		}
	}

	return result, nil
}

// processSheet processes a single sheet
func processSheet(sheet *xlsx.Sheet, config CleanConfig, result *CleanResult) error {
	// First, identify headers and empty columns
	maxRow := sheet.MaxRow
	maxCol := sheet.MaxCol

	if maxRow == 0 {
		return nil // Empty sheet
	}

	// Get headers from first row
	headers := make([]string, maxCol)
	headerRow, err := sheet.Row(0)
	if err != nil {
		return fmt.Errorf("failed to get header row: %w", err)
	}

	for colIdx := 0; colIdx < maxCol; colIdx++ {
		cell := headerRow.GetCell(colIdx)
		headers[colIdx] = cell.String()
	}

	// Identify empty columns and date columns
	emptyColumns := make(map[int]bool)
	dateColumns := make(map[int]bool)

	for colIdx := 0; colIdx < maxCol; colIdx++ {
		isEmpty := isColumnEmpty(sheet, colIdx, maxRow)
		headerName := headers[colIdx]

		// Check if this column should be kept even if empty
		shouldKeep := contains(config.KeepHeaders, headerName)

		if isEmpty && !shouldKeep {
			emptyColumns[colIdx] = true
			result.RemovedColumns = append(result.RemovedColumns, headerName)
		}

		// Check if column name contains "Date"
		if strings.Contains(strings.ToLower(headerName), "date") {
			dateColumns[colIdx] = true
			result.DateColumns = append(result.DateColumns, headerName)
		}
	}

	result.ColumnsRemoved = len(emptyColumns)
	result.DateColumnsReformatted = len(dateColumns)

	// Process date columns - reformat to dd-mmm-yyyy
	for colIdx := range dateColumns {
		if err := reformatDateColumn(sheet, colIdx, maxRow); err != nil {
			return fmt.Errorf("failed to reformat date column %d: %w", colIdx, err)
		}
	}

	// Remove empty columns (do this after date formatting)
	// We need to remove columns in reverse order to maintain indices
	columnsToRemove := make([]int, 0, len(emptyColumns))
	for colIdx := range emptyColumns {
		columnsToRemove = append(columnsToRemove, colIdx)
	}

	// Sort in descending order
	for i := 0; i < len(columnsToRemove); i++ {
		for j := i + 1; j < len(columnsToRemove); j++ {
			if columnsToRemove[i] < columnsToRemove[j] {
				columnsToRemove[i], columnsToRemove[j] = columnsToRemove[j], columnsToRemove[i]
			}
		}
	}

	// Remove columns
	for _, colIdx := range columnsToRemove {
		if err := removeColumn(sheet, colIdx, maxRow); err != nil {
			return fmt.Errorf("failed to remove column %d: %w", colIdx, err)
		}
	}

	return nil
}

// isColumnEmpty checks if a column is empty (excluding header)
func isColumnEmpty(sheet *xlsx.Sheet, colIdx int, maxRow int) bool {
	// Start from row 1 (skip header)
	for rowIdx := 1; rowIdx < maxRow; rowIdx++ {
		row, err := sheet.Row(rowIdx)
		if err != nil {
			continue
		}

		cell := row.GetCell(colIdx)
		if cell.String() != "" {
			return false
		}
	}
	return true
}

// reformatDateColumn reformats all dates in a column to dd-mmm-yyyy format
func reformatDateColumn(sheet *xlsx.Sheet, colIdx int, maxRow int) error {
	// Start from row 1 (skip header)
	for rowIdx := 1; rowIdx < maxRow; rowIdx++ {
		row, err := sheet.Row(rowIdx)
		if err != nil {
			continue
		}

		cell := row.GetCell(colIdx)
		cellValue := cell.String()

		if cellValue == "" {
			continue
		}

		// Try to parse the date in various formats
		date, err := parseDate(cellValue)
		if err != nil {
			// If we can't parse it, leave it as is
			continue
		}

		// Format as dd-mmm-yyyy
		formatted := date.Format("02-Jan-2006")
		cell.SetString(formatted)
	}

	return nil
}

// parseDate attempts to parse a date string in various formats
func parseDate(dateStr string) (time.Time, error) {
	// Common date formats to try
	formats := []string{
		"2006-01-02",          // ISO format
		"01/02/2006",          // US format
		"02/01/2006",          // UK format
		"2006/01/02",          // Alt ISO
		"01-02-2006",          // US with dashes
		"02-01-2006",          // UK with dashes
		"2006-01-02 15:04:05", // ISO with time
		"01/02/2006 15:04:05", // US with time
		"02-Jan-2006",         // dd-mmm-yyyy
		time.RFC3339,          // RFC3339
	}

	for _, format := range formats {
		if date, err := time.Parse(format, dateStr); err == nil {
			return date, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse date: %s", dateStr)
}

// removeColumn removes a column from the sheet
func removeColumn(sheet *xlsx.Sheet, colIdx int, maxRow int) error {
	for rowIdx := 0; rowIdx < maxRow; rowIdx++ {
		row, err := sheet.Row(rowIdx)
		if err != nil {
			continue
		}

		// Collect all cell values except the one we're removing
		var cellValues []string
		cellCount := 0
		row.ForEachCell(func(c *xlsx.Cell) error {
			if cellCount != colIdx {
				cellValues = append(cellValues, c.String())
			}
			cellCount++
			return nil
		})

		// Clear the entire row
		for i := 0; i < cellCount; i++ {
			cell := row.GetCell(i)
			if cell != nil {
				cell.SetString("")
			}
		}

		// Repopulate with the values (excluding the removed column)
		for i, value := range cellValues {
			cell := row.GetCell(i)
			if cell != nil {
				cell.SetString(value)
			}
		}
	}

	return nil
}

// contains checks if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

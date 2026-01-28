# xelo

*Automated manipulation of excel files*  **this is primarily just a testing grounds**

## Install Dependencies

```bash
go mod init demographic-generator
go get github.com/ivahaev/go-xlsx-templater
```

### Write Your Code

**Minimal Example:**

```go
package main

import (
    "log"
    xlst "github.com/ivahaev/go-xlsx-templater"
)

func main() {
    // Load template
    doc := xlst.New()
    doc.ReadTemplate("demographic_template_with_placeholders.xlsx")

    // Prepare data (maps to the placeholders in Row 2)
    data := map[string]interface{}{
        "As_Of_Date": "2026-01-27",
        "Internal_Participant_Id": "INT001",
        "External_Participant_Id": "EXT001",
        "Employee_Id": "EMP001",
        "Stockholder_Id": "STK001",
        "First_Name": "John",
        "Middle_Name": "A",
        "Last_Name": "Doe",
    }

    // Render and save
    doc.Render(data)
    doc.Save("output.xlsx")
}
```

### Run

```bash
go run complete_example.go
```

This will generate three example files:
- `demographic_single.xlsx` - Single record
- `demographic_multiple.xlsx` - Multiple records
- `demographic_structs.xlsx` - Using Go structs

## Key Concepts

### 1. Placeholder Syntax
- Use `{{FieldName}}` in your template cells
- Field names are case-sensitive
- Must exactly match your data keys

### 2. Multiple Records
For multiple rows, wrap your data in an array:

```go
data := map[string]interface{}{
    "Records": []map[string]string{
        {
            "As_Of_Date": "2026-01-27",
            "Internal_Participant_Id": "INT001",
            // ... more fields
        },
        {
            "As_Of_Date": "2026-01-27",
            "Internal_Participant_Id": "INT002",
            // ... more fields
        },
    },
}
```

### 3. Data Flow

```
Template File (with {{placeholders}})
         ↓
    Load with doc.ReadTemplate()
         ↓
    Prepare Data (map or struct)
         ↓
    Render with doc.Render(data)
         ↓
    Save with doc.Save()
         ↓
    Output File (populated with data)
```

## Common Patterns

### Reading from Database

```go
// Fetch from database
rows, _ := db.Query("SELECT * FROM demographics")
defer rows.Close()

var records []map[string]string
for rows.Next() {
    var asOfDate, internalID, externalID, empID, stockID, firstName, middleName, lastName string
    rows.Scan(&asOfDate, &internalID, &externalID, &empID, &stockID, &firstName, &middleName, &lastName)
    
    records = append(records, map[string]string{
        "As_Of_Date": asOfDate,
        "Internal_Participant_Id": internalID,
        "External_Participant_Id": externalID,
        "Employee_Id": empID,
        "Stockholder_Id": stockID,
        "First_Name": firstName,
        "Middle_Name": middleName,
        "Last_Name": lastName,
    })
}

// Render template
data := map[string]interface{}{"Records": records}
doc.Render(data)
```

### Reading from CSV

```go
import "encoding/csv"

file, _ := os.Open("data.csv")
reader := csv.NewReader(file)
records, _ := reader.ReadAll()

var data []map[string]string
for i, record := range records {
    if i == 0 { continue } // Skip header
    
    data = append(data, map[string]string{
        "As_Of_Date": record[0],
        "Internal_Participant_Id": record[1],
        "External_Participant_Id": record[2],
        "Employee_Id": record[3],
        "Stockholder_Id": record[4],
        "First_Name": record[5],
        "Middle_Name": record[6],
        "Last_Name": record[7],
    })
}

templateData := map[string]interface{}{"Records": data}
```

### Using with REST API

```go
type DemographicData struct {
    Records []map[string]string `json:"records"`
}

http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
    var data DemographicData
    json.NewDecoder(r.Body).Decode(&data)
    
    doc := xlst.New()
    doc.ReadTemplate("template.xlsx")
    doc.Render(map[string]interface{}{"Records": data.Records})
    doc.Save("output.xlsx")
    
    // Serve file
    http.ServeFile(w, r, "output.xlsx")
})
```

## Troubleshooting

**Problem:** Placeholders not being replaced
**Solution:** Check that placeholder names exactly match data keys (case-sensitive)

**Problem:** Only first record appears
**Solution:** Make sure you're passing data as `{"Records": [...]}` not just `[...]`

**Problem:** Template not found
**Solution:** Use absolute paths or ensure file is in working directory

**Problem:** Formatting lost
**Solution:** The template preserves formatting from the original file; check your template file

## Next Steps

1. Review `complete_example.go` for full working examples
2. Modify the data structure to match your actual data source
3. Customize the template with your desired formatting
4. Add error handling for production use
5. Consider adding validation for required fields

## Resources

- Full documentation: See `README.md`
- Example code: `complete_example.go`, `simple_example.go`, `demographic_generator.go`
- Package docs: [xlsx-templater](https://github.com/ivahaev/go-xlsx-templater) & [xlsx v3](https://github.com/tealeg/xlsx/v3)
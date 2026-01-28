## `xlsx` - VERSION 4

### Constants [¶](#pkg-constants "Go to Constants")

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/data_validation.go#L13)

```
const (
	DataValidationTypeCustom
	DataValidationTypeDate
	DataValidationTypeDecimal

	DataValidationTypeTextLeng
	DataValidationTypeTime
	// DataValidationTypeWhole Integer
	DataValidationTypeWhole
)
```

Data validation types

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/data_validation.go#L54)

```
const (
	DataValidationOperatorBetween
	DataValidationOperatorEqual
	DataValidationOperatorGreaterThan
	DataValidationOperatorGreaterThanOrEqual
	DataValidationOperatorLessThan
	DataValidationOperatorLessThanOrEqual
	DataValidationOperatorNotBetween
	DataValidationOperatorNotEqual
)
```

Data validation operators

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/date.go#L8)

```
const (
	MJD_0      float64 = 2400000.5
	MJD_JD2000 float64 = 51544.5
)
```

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L16)

```
const (
	TRUE  = 0x01
	FALSE = 0x00
	US    = 0x1f // Unit Separator
	RS    = 0x1e // Record Separator
	GS    = 0x1d // Group Separator
)
```

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/richtext.go#L13)

```
const (
	// RichTextFontFamilyUnspecified indicates that the font family was not specified
	RichTextFontFamilyUnspecified   RichTextFontFamily = -1
	RichTextFontFamilyNotApplicable RichTextFontFamily = 0
	RichTextFontFamilyRoman         RichTextFontFamily = 1
	RichTextFontFamilySwiss         RichTextFontFamily = 2
	RichTextFontFamilyModern        RichTextFontFamily = 3
	RichTextFontFamilyScript        RichTextFontFamily = 4
	RichTextFontFamilyDecorative    RichTextFontFamily = 5

	// RichTextCharsetUnspecified indicates that the font charset was not specified
	RichTextCharsetUnspecified RichTextCharset = -1
	RichTextCharsetANSI        RichTextCharset = 0
	RichTextCharsetDefault     RichTextCharset = 1
	RichTextCharsetSymbol      RichTextCharset = 2
	RichTextCharsetMac         RichTextCharset = 77
	RichTextCharsetShiftJIS    RichTextCharset = 128
	RichTextCharsetHangul      RichTextCharset = 129
	RichTextCharsetJohab       RichTextCharset = 130
	RichTextCharsetGB2312      RichTextCharset = 134
	RichTextCharsetBIG5        RichTextCharset = 136
	RichTextCharsetGreek       RichTextCharset = 161
	RichTextCharsetTurkish     RichTextCharset = 162
	RichTextCharsetVietnamese  RichTextCharset = 163
	RichTextCharsetHebrew      RichTextCharset = 177
	RichTextCharsetArabic      RichTextCharset = 178
	RichTextCharsetBaltic      RichTextCharset = 186
	RichTextCharsetRussian     RichTextCharset = 204
	RichTextCharsetThai        RichTextCharset = 222
	RichTextCharsetEastEurope  RichTextCharset = 238
	RichTextCharsetOEM         RichTextCharset = 255

	RichTextVertAlignSuperscript RichTextVertAlign = "superscript"
	RichTextVertAlignSubscript   RichTextVertAlign = "subscript"

	RichTextUnderlineSingle RichTextUnderline = "single"
	RichTextUnderlineDouble RichTextUnderline = "double"
)
```

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L6)

```
const (
	Helvetica     = "Helvetica"
	Baskerville   = "Baskerville Old Face"
	TimesNewRoman = "Times New Roman"
	Bodoni        = "Bodoni MT"
	GillSans      = "Gill Sans MT"
	Courier       = "Courier"
)
```

Several popular font names that can be used to create fonts

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L15)

```
const (
	RGB_Light_Green = "FFC6EFCE"
	RGB_Dark_Green  = "FF006100"
	RGB_Light_Red   = "FFFFC7CE"
	RGB_Dark_Red    = "FF9C0006"
	RGB_White       = "FFFFFFFF"
	RGB_Black       = "00000000"
)
```

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/col.go#L4)

```
const ColWidth = 9.5
```

Default column width in excel

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/reftable.go#L3)

```
const DEFAULT_REFTABLE_SIZE = 500
```

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/col.go#L5)

```
const Excel2006MaxRowCount = 1048576
```

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/col.go#L6)

```
const Excel2006MaxRowIndex = Excel2006MaxRowCount - 1
```

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L34)

```
const NoColLimit int = -1
```

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L33)

```
const NoRowLimit int = -1
```

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L24)

```
const (
	Solid_Cell_Fill = "solid"
)
```

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/templates.go#L13)

```
const TEMPLATE_DOCPROPS_APP = `` /* 308-byte string literal not displayed */
```

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/templates.go#L19)

```
const TEMPLATE_DOCPROPS_CORE = `` /* 364-byte string literal not displayed */
```

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/templates.go#L22)

```
const TEMPLATE_XL_THEME_THEME = `` /* 10940-byte string literal not displayed */
```

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/templates.go#L6)

```
const TEMPLATE__RELS_DOT_RELS = `` /* 580-byte string literal not displayed */
```

### Variables [¶](#pkg-variables "Go to Variables")

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L304)

```
var (
	DefaultDateFormat     = builtInNumFmt[14]
	DefaultDateTimeFormat = builtInNumFmt[22]

	DefaultDateOptions = DateTimeOptions{
		Location:        timeLocationUTC,
		ExcelTimeFormat: DefaultDateFormat,
	}

	DefaultDateTimeOptions = DateTimeOptions{
		Location:        timeLocationUTC,
		ExcelTimeFormat: DefaultDateTimeFormat,
	}
)
```

[View Source](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/hsl.go#L39)

```
var HSLModel = color.ModelFunc(hslModel)
```

HSLModel converts any color.Color to a HSL color.

### Functions [¶](#pkg-functions "Go to Functions")

#### func [ColIndexToLetters](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/lib.go#L94) [¶](#ColIndexToLetters "Go to ColIndexToLetters")

```
func ColIndexToLetters(n int) string
```

ColIndexToLetters is used to convert a zero based, numeric column indentifier into a character code.

#### func [ColLettersToIndex](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/lib.go#L77) [¶](#ColLettersToIndex "Go to ColLettersToIndex")

```
func ColLettersToIndex(letters string) int
```

ColLettersToIndex is used to convert a character based column reference to a zero based numeric column identifier.

#### func [DefaultAutoWidth](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L429) [¶](#DefaultAutoWidth "Go to DefaultAutoWidth")

```
func DefaultAutoWidth(s string) float64
```

This can be use as the default scale function for the autowidth. It works well with the default font sizes.

#### func [FileToSlice](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L135) [¶](#FileToSlice "Go to FileToSlice")

```
func FileToSlice(path string, options ...FileOption) ([][][]string, error)
```

A convenient wrapper around File.ToSlice, FileToSlice will return the raw data contained in an Excel XLSX file as three dimensional slice. The first index represents the sheet number, the second the row number, and the third the cell number.

For example:

```
var mySlice [][][]string
var value string
mySlice = xlsx.FileToSlice("myXLSX.xlsx")
value = mySlice[0][0][0]
```

Here, value would be set to the raw value of the cell A1 in the first sheet in the XLSX file.

#### func [FileToSliceUnmerged](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L147) [¶](#FileToSliceUnmerged "Go to FileToSliceUnmerged")

```
func FileToSliceUnmerged(path string, options ...FileOption) ([][][]string, error)
```

FileToSliceUnmerged is a wrapper around File.ToSliceUnmerged. It returns the raw data contained in an Excel XLSX file as three dimensional slice. Merged cells will be unmerged. Covered cells become the values of theirs origins.

#### func [GetCellIDStringFromCoords](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/lib.go#L155) [¶](#GetCellIDStringFromCoords "Go to GetCellIDStringFromCoords")

```
func GetCellIDStringFromCoords(x, y int) string
```

GetCellIDStringFromCoords returns the Excel format cell name that represents a pair of zero based cartesian coordinates.

#### func [GetCellIDStringFromCoordsWithFixed](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/lib.go#L162) [¶](#GetCellIDStringFromCoordsWithFixed "Go to GetCellIDStringFromCoordsWithFixed")

```
func GetCellIDStringFromCoordsWithFixed(x, y int, xFixed, yFixed bool) string
```

GetCellIDStringFromCoordsWithFixed returns the Excel format cell name that represents a pair of zero based cartesian coordinates. It can specify either value as fixed.

#### func [GetCoordsFromCellIDString](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/lib.go#L139) [¶](#GetCoordsFromCellIDString "Go to GetCoordsFromCellIDString")

```
func GetCoordsFromCellIDString(cellIDString string) (x, y int, err error)
```

GetCoordsFromCellIDString returns the zero based cartesian coordinates from a cell name in Excel format, e.g. the cellIDString "A1" returns 0, 0 and the "B3" return 1, 2.

#### func [HSLToRGB](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/hsl.go#L105) [¶](#HSLToRGB "Go to HSLToRGB")

```
func HSLToRGB(h, s, l float64) (r, g, b uint8)
```

HSLToRGB converts an HSL triple to a RGB triple.

Ported from [http://goo.gl/Vg1h9](http://goo.gl/Vg1h9)

#### func [IsSaneSheetName](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L975) [¶](#IsSaneSheetName "Go to IsSaneSheetName")

```
func IsSaneSheetName(sheetName string) error
```

#### func [MakeDefaultContentTypes](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/xmlContentTypes.go#L24) [¶](#MakeDefaultContentTypes "Go to MakeDefaultContentTypes")

```
func MakeDefaultContentTypes() (types xlsxTypes)
```

#### func [NewDataValidation](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/data_validation.go#L67) [¶](#NewDataValidation "Go to NewDataValidation")

```
func NewDataValidation(startRow, startCol, endRow, endCol int, allowBlank bool) *xlsxDataValidation
```

NewDataValidation return data validation struct

#### func [RGBToHSL](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/hsl.go#L68) [¶](#RGBToHSL "Go to RGBToHSL")

```
func RGBToHSL(r, g, b uint8) (h, s, l float64)
```

RGBToHSL converts an RGB triple to a HSL triple.

Ported from [http://goo.gl/Vg1h9](http://goo.gl/Vg1h9)

#### func [RowIndexToString](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/lib.go#L111) [¶](#RowIndexToString "Go to RowIndexToString")

```
func RowIndexToString(rowRef int) string
```

RowIndexToString is used to convert a zero based, numeric row indentifier into its string representation.

#### func [SetDefaultFont](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L191) [¶](#SetDefaultFont "Go to SetDefaultFont")

```
func SetDefaultFont(size float64, name string)
```

#### func [SkipEmptyCells](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/row.go#L116) [¶](#SkipEmptyCells "Go to SkipEmptyCells")

```
func SkipEmptyCells(flags *cellVisitorFlags)
```

SkipEmptyCells can be passed as an option to [Row.ForEachCell](#Row.ForEachCell) in order to make it skip over empty cells in the sheet.

#### func [SkipEmptyRows](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L151) [¶](#SkipEmptyRows "Go to SkipEmptyRows")

```
func SkipEmptyRows(flags *rowVisitorFlags)
```

SkipEmptyRows can be passed to the Sheet.ForEachRow function to cause it to skip over empty Rows.

#### func [TimeFromExcelTime](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/date.go#L104) [¶](#TimeFromExcelTime "Go to TimeFromExcelTime")

```
func TimeFromExcelTime(excelTime float64, date1904 bool) time.Time
```

Convert an excelTime representation (stored as a floating point number) to a time.Time.

#### func [TimeToExcelTime](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/date.go#L133) [¶](#TimeToExcelTime "Go to TimeToExcelTime")

```
func TimeToExcelTime(t time.Time, date1904 bool) float64
```

TimeToExcelTime will convert a time.Time into Excel's float representation, in either 1900 or 1904 mode. If you don't know which to use, set date1904 to false. TODO should this should handle Julian dates?

#### func [TimeToUTCTime](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/date.go#L32) [¶](#TimeToUTCTime "Go to TimeToUTCTime")

```
func TimeToUTCTime(t time.Time) time.Time
```

#### func [UseDiskVCellStore](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L348) [¶](#UseDiskVCellStore "Go to UseDiskVCellStore")

```
func UseDiskVCellStore(f *File)
```

UseDiskVCellStore is a FileOption that makes all Sheet instances for a File use DiskV as their backing store. You can use this option when handling very large Sheets that would otherwise require allocating vast amounts of memory.

#### func [UseMemoryCellStore](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L146) [¶](#UseMemoryCellStore "Go to UseMemoryCellStore")

```
func UseMemoryCellStore(f *File)
```

UseMemoryCellStore is a FileOption that makes all Sheet instances for a File use memory as their backing store. This is the default backing store. You can use this option when you are comfortable keeping the contents of each Sheet in memory. This is faster than using a disk backed store, but can easily use a large amount of memory and, if you exhaust the available system memory, it'll actualy be slower than using a disk backed store (e.g. DiskV).

### Types [¶](#pkg-types "Go to Types")

#### type [Alignment](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L179) [¶](#Alignment "Go to Alignment")

```
type Alignment struct {
	Horizontal   string
	Indent       int
	ShrinkToFit  bool
	TextRotation int
	Vertical     string
	WrapText     bool
}
```

#### func [DefaultAlignment](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L209) [¶](#DefaultAlignment "Go to DefaultAlignment")

```
func DefaultAlignment() *Alignment
```

#### type [AutoFilter](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L95) [¶](#AutoFilter "Go to AutoFilter")

```
type AutoFilter struct {
	TopLeftCell     string
	BottomRightCell string
}
```

#### type [Border](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L127) [¶](#Border "Go to Border")

```
type Border struct {
	Left        string
	LeftColor   string
	Right       string
	RightColor  string
	Top         string
	TopColor    string
	Bottom      string
	BottomColor string
}
```

Border is a high level structure intended to provide user access to the contents of Border Style within an Sheet.

#### func [DefaultBorder](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L205) [¶](#DefaultBorder "Go to DefaultBorder")

```
func DefaultBorder() *Border
```

#### func [NewBorder](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L138) [¶](#NewBorder "Go to NewBorder")

```
func NewBorder(left, right, top, bottom string) *Border
```

#### type [Cell](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L60) [¶](#Cell "Go to Cell")

```
type Cell struct {
	Row      *Row
	Value    string
	RichText []RichTextRun

	NumFmt string

	Hidden bool
	HMerge int
	VMerge int

	DataValidation *xlsxDataValidation
	Hyperlink      Hyperlink
	// contains filtered or unexported fields
}
```

Cell is a high level structure intended to provide user access to the contents of Cell within an xlsx.Row.

#### func (\*Cell) [Bool](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L483) [¶](#Cell.Bool "Go to Cell.Bool")

```
func (c *Cell) Bool() bool
```

Bool returns a boolean from a cell's value. TODO: Determine if the current return value is appropriate for types other than CellTypeBool.

#### func (\*Cell) [Float](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L349) [¶](#Cell.Float "Go to Cell.Float")

```
func (c *Cell) Float() (float64, error)
```

Float returns the value of cell as a number.

#### func (\*Cell) [FormattedValue](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L558) [¶](#Cell.FormattedValue "Go to Cell.FormattedValue")

```
func (c *Cell) FormattedValue() (string, error)
```

FormattedValue returns a value, and possibly an error condition from a Cell. If it is possible to apply a format to the cell value, it will do so, if not then an error will be returned, along with the raw value of the Cell.

#### func (\*Cell) [Formula](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L517) [¶](#Cell.Formula "Go to Cell.Formula")

```
func (c *Cell) Formula() string
```

Formula returns the formula string for the cell.

#### func (\*Cell) [GeneralNumeric](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L375) [¶](#Cell.GeneralNumeric "Go to Cell.GeneralNumeric")

```
func (c *Cell) GeneralNumeric() (string, error)
```

GeneralNumeric returns the value of the cell as a string. It is formatted very closely to the the XLSX spec for how to display values when the storage type is Number and the format type is General. It is not 100% identical to the spec but is as close as you can get using the built in Go formatting tools.

#### func (\*Cell) [GeneralNumericWithoutScientific](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L382) [¶](#Cell.GeneralNumericWithoutScientific "Go to Cell.GeneralNumericWithoutScientific")

```
func (c *Cell) GeneralNumericWithoutScientific() (string, error)
```

GeneralNumericWithoutScientific returns numbers that are always formatted as numbers, but it does not follow the rules for when XLSX should switch to scientific notation, since sometimes scientific notation is not desired, even if that is how the document is supposed to be formatted.

#### func (\*Cell) [GetCoordinates](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L580) [¶](#Cell.GetCoordinates "Go to Cell.GetCoordinates")

```
func (c *Cell) GetCoordinates() (int, int)
```

GetCoordinates returns a pair of integers representing the cartesian coorindates of the Cell within the Sheet. The coordinates are zero based and a returned in order x,y where x is the Column number and y is the Row number. If you need to convert these numbers to a Excel cellID (i.e. B15) then please see the GetCellIDStringFromCoords function.

#### func (\*Cell) [GetNumberFormat](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L541) [¶](#Cell.GetNumberFormat "Go to Cell.GetNumberFormat")

```
func (c *Cell) GetNumberFormat() string
```

GetNumberFormat returns the number format string for a cell.

#### func (\*Cell) [GetStyle](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L526) [¶](#Cell.GetStyle "Go to Cell.GetStyle")

```
func (c *Cell) GetStyle() *Style
```

GetStyle returns the Style associated with a Cell

#### func (\*Cell) [GetTime](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L260) [¶](#Cell.GetTime "Go to Cell.GetTime")

```
func (c *Cell) GetTime(date1904 bool) (t time.Time, err error)
```

GetTime returns the value of a Cell as a time.Time

#### func (\*Cell) [HasFormula](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L512) [¶](#Cell.HasFormula "Go to Cell.HasFormula") added in v4.1.0

```
func (c *Cell) HasFormula() bool
```

HasFormula returns true if the Cell has an associated, cell specific formula

#### func (\*Cell) [Int](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L460) [¶](#Cell.Int "Go to Cell.Int")

```
func (c *Cell) Int() (int, error)
```

Int returns the value of cell as integer. Has max 53 bits of precision See: float64(int64(math.MaxInt))

#### func (\*Cell) [Int64](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L364) [¶](#Cell.Int64 "Go to Cell.Int64")

```
func (c *Cell) Int64() (int64, error)
```

Int64 returns the value of cell as 64-bit integer.

#### func (\*Cell) [IsTime](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L254) [¶](#Cell.IsTime "Go to Cell.IsTime")

```
func (c *Cell) IsTime() bool
```

IsTime returns true if the cell stores a time value.

#### func (Cell) [MarshalBinary](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L83) [¶](#Cell.MarshalBinary "Go to Cell.MarshalBinary")

```
func (c Cell) MarshalBinary() ([]byte, error)
```

Return a representation of the Cell as a slice of bytes

#### func (\*Cell) [Merge](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L204) [¶](#Cell.Merge "Go to Cell.Merge")

```
func (c *Cell) Merge(hcells, vcells int)
```

Merge with other cells, horizontally and/or vertically.

#### func (\*Cell) [Modified](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L149) [¶](#Cell.Modified "Go to Cell.Modified")

```
func (c *Cell) Modified() bool
```

Modified returns True if a cell has been modified since it was last persisted.

#### func (\*Cell) [SetBool](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L469) [¶](#Cell.SetBool "Go to Cell.SetBool")

```
func (c *Cell) SetBool(b bool)
```

SetBool sets a cell's value to a boolean.

#### func (\*Cell) [SetDataValidation](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L568) [¶](#Cell.SetDataValidation "Go to Cell.SetDataValidation")

```
func (c *Cell) SetDataValidation(dd *xlsxDataValidation)
```

SetDataValidation set data validation

#### func (\*Cell) [SetDate](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L320) [¶](#Cell.SetDate "Go to Cell.SetDate")

```
func (c *Cell) SetDate(t time.Time)
```

SetDate sets the value of a cell to a float.

#### func (\*Cell) [SetDateTime](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L325) [¶](#Cell.SetDateTime "Go to Cell.SetDateTime")

```
func (c *Cell) SetDateTime(t time.Time)
```

#### func (\*Cell) [SetDateTimeWithFormat](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L339) [¶](#Cell.SetDateTimeWithFormat "Go to Cell.SetDateTimeWithFormat")

```
func (c *Cell) SetDateTimeWithFormat(n float64, format string)
```

#### func (\*Cell) [SetDateWithOptions](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L331) [¶](#Cell.SetDateWithOptions "Go to Cell.SetDateWithOptions")

```
func (c *Cell) SetDateWithOptions(t time.Time, options DateTimeOptions)
```

SetDateWithOptions allows for more granular control when exporting dates and times

#### func (\*Cell) [SetFloat](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L248) [¶](#Cell.SetFloat "Go to Cell.SetFloat")

```
func (c *Cell) SetFloat(n float64)
```

SetFloat sets the value of a cell to a float.

#### func (\*Cell) [SetFloatWithFormat](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L282) [¶](#Cell.SetFloatWithFormat "Go to Cell.SetFloatWithFormat")

```
func (c *Cell) SetFloatWithFormat(n float64, format string)
```

SetFloatWithFormat sets the value of a cell to a float and applies formatting to the cell.

#### func (\*Cell) [SetFormat](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L290) [¶](#Cell.SetFormat "Go to Cell.SetFormat")

```
func (c *Cell) SetFormat(format string)
```

SetCellFormat set cell value format

#### func (\*Cell) [SetFormula](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L497) [¶](#Cell.SetFormula "Go to Cell.SetFormula")

```
func (c *Cell) SetFormula(formula string)
```

SetFormula sets the format string for a cell.

#### func (\*Cell) [SetHyperlink](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L398) [¶](#Cell.SetHyperlink "Go to Cell.SetHyperlink")

```
func (c *Cell) SetHyperlink(hyperlink string, displayText string, tooltip string)
```

SetHyperlink sets this cell to contain the given hyperlink, displayText and tooltip. If the displayText or tooltip are an empty string, they will not be set. The hyperlink provided must be a valid URL starting with http:// or https:// or excel will not recognize it as an external link. All other hyperlink formats will be treated as internal link between sheets. Official format in form of \`#Sheet!A123\`. Maximum number of hyperlinks per sheet is 65530, according to specification.

#### func (\*Cell) [SetInt](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L387) [¶](#Cell.SetInt "Go to Cell.SetInt")

```
func (c *Cell) SetInt(n int)
```

SetInt sets a cell's value to an integer.

#### func (\*Cell) [SetInt64](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L358) [¶](#Cell.SetInt64 "Go to Cell.SetInt64")

```
func (c *Cell) SetInt64(n int64)
```

SetInt64 sets a cell's value to a 64-bit integer.

#### func (\*Cell) [SetNumeric](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L448) [¶](#Cell.SetNumeric "Go to Cell.SetNumeric")

```
func (c *Cell) SetNumeric(s string)
```

SetNumeric sets a cell's value to a number

#### func (\*Cell) [SetRichText](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L227) [¶](#Cell.SetRichText "Go to Cell.SetRichText")

```
func (c *Cell) SetRichText(r []RichTextRun)
```

SetRichText sets the value of a cell to a set of the rich text.

#### func (\*Cell) [SetString](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L217) [¶](#Cell.SetString "Go to Cell.SetString")

```
func (c *Cell) SetString(s string)
```

SetString sets the value of a cell to a string.

#### func (\*Cell) [SetStringFormula](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L504) [¶](#Cell.SetStringFormula "Go to Cell.SetStringFormula")

```
func (c *Cell) SetStringFormula(formula string)
```

#### func (\*Cell) [SetStyle](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L534) [¶](#Cell.SetStyle "Go to Cell.SetStyle")

```
func (c *Cell) SetStyle(style *Style)
```

SetStyle sets the style of a cell.

#### func (\*Cell) [SetValue](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L418) [¶](#Cell.SetValue "Go to Cell.SetValue")

```
func (c *Cell) SetValue(n interface{})
```

SetValue sets a cell's value to any type.

#### func (\*Cell) [String](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L239) [¶](#Cell.String "Go to Cell.String")

```
func (c *Cell) String() string
```

String returns the value of a Cell as a string. If you'd like to see errors returned from formatting then please use Cell.FormattedValue() instead.

#### func (\*Cell) [Type](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L212) [¶](#Cell.Type "Go to Cell.Type")

```
func (c *Cell) Type() CellType
```

Type returns the CellType of a cell. See CellType constants for more details.

#### func (\*Cell) [UnmarshalBinary](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L110) [¶](#Cell.UnmarshalBinary "Go to Cell.UnmarshalBinary")

```
func (c *Cell) UnmarshalBinary(data []byte) error
```

Read a slice of bytes, produced by MarshalBinary, into a Cell

#### type [CellInterface](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L184) [¶](#CellInterface "Go to CellInterface")

```
type CellInterface interface {
	String() string
	FormattedValue() string
}
```

CellInterface defines the public API of the Cell.

#### type [CellStore](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cellstore.go#L16) [¶](#CellStore "Go to CellStore")

```
type CellStore interface {
	MakeRow(sheet *Sheet) *Row
	MakeRowWithLen(sheet *Sheet, len int) *Row
	ReadRow(key string, sheet *Sheet) (*Row, error)
	WriteRow(r *Row) error
	MoveRow(r *Row, newIndex int) error
	RemoveRow(key string) error
	Close() error
}
```

CellStore provides an interface for interacting with backend cell storage. For example, this allows us, as required, to persist cells to some store instead of holding them in memory. This tactic allows us a degree of control around the characteristics of our programs when handling large spreadsheets - we can choose to run more slowly, but without exhausting system memory.

If you wish to implement a custom CellStore you must not only support this interface, but also a CellStoreConstructor and a FileOption that set's the File's cellStoreConstructor to the right constructor.

#### func [NewDiskVCellStore](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L356) [¶](#NewDiskVCellStore "Go to NewDiskVCellStore")

```
func NewDiskVCellStore() (CellStore, error)
```

NewDiskVCellStore is a CellStoreConstructor than returns a CellStore in terms of DiskV.

#### func [NewMemoryCellStore](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L151) [¶](#NewMemoryCellStore "Go to NewMemoryCellStore")

```
func NewMemoryCellStore() (CellStore, error)
```

NewMemoryCellStore returns a pointer to a newly allocated MemoryCellStore

#### type [CellStoreConstructor](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cellstore.go#L29) [¶](#CellStoreConstructor "Go to CellStoreConstructor")

```
type CellStoreConstructor func() (CellStore, error)
```

CellStoreConstructor defines the signature of a function that will be used to return a new instance of the CellStore implmentation, you must pass this into

#### type [CellStoreRow](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cellstore.go#L32) [¶](#CellStoreRow "Go to CellStoreRow")

```
type CellStoreRow interface {
	AddCell() *Cell
	GetCell(colIdx int) *Cell
	PushCell(c *Cell)
	ForEachCell(cvf CellVisitorFunc, option ...CellVisitorOption) error
	MaxCol() int
	CellCount() int
	Updatable()
	CellUpdatable(c *Cell)
}
```

CellStoreRow is the interface used to interact with the currently loaded Row from the CellStore. Different backends can choose whether to hold the whole row in memory, or persist and load the cell

#### type [CellType](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L19) [¶](#CellType "Go to CellType")

```
type CellType int
```

CellType is an int type for storing metadata about the data type in the cell.

```
const (
	CellTypeString CellType = iota
	// CellTypeStringFormula is a specific format for formulas that return string values. Formulas that return numbers
	// and booleans are stored as those types.
	CellTypeStringFormula
	CellTypeNumeric
	CellTypeBool
	// CellTypeInline is not respected on save, all inline string cells will be saved as SharedStrings
	// when saving to an XLSX file. This the same behavior as that found in Excel.
	CellTypeInline
	CellTypeError
	// d (Date): Cell contains a date in the ISO 8601 format.
	// That is the only mention of this format in the XLSX spec.
	// Date seems to be unused by the current version of Excel, it stores dates as Numeric cells with a date format string.
	// For now these cells will have their value output directly. It is unclear if the value is supposed to be parsed
	// into a number and then formatted using the formatting or not.
	CellTypeDate
)
```

These are the cell types from the ST\_CellType spec

#### func (CellType) [Ptr](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L41) [¶](#CellType.Ptr "Go to CellType.Ptr")

```
func (ct CellType) Ptr() *CellType
```

#### type [CellVisitorFunc](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cellstore.go#L45) [¶](#CellVisitorFunc "Go to CellVisitorFunc")

```
type CellVisitorFunc func(c *Cell) error
```

CellVisitorFunc defines the signature of a function that will be called when visiting a Cell using CellStore.ForEachInRow.

#### type [CellVisitorOption](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/row.go#L112) [¶](#CellVisitorOption "Go to CellVisitorOption")

```
type CellVisitorOption func(flags *cellVisitorFlags)
```

CellVisitorOption is an option for [Row.ForEachCell](#Row.ForEachCell).

Available options: [SkipEmptyCells](#SkipEmptyCells).

#### type [Col](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/col.go#L8) [¶](#Col "Go to Col")

```
type Col struct {
	Min          int
	Max          int
	Hidden       *bool
	Width        *float64
	Collapsed    *bool
	OutlineLevel *uint8
	BestFit      *bool
	CustomWidth  *bool
	Phonetic     *bool
	// contains filtered or unexported fields
}
```

#### func [NewColForRange](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/col.go#L29) [¶](#NewColForRange "Go to NewColForRange")

```
func NewColForRange(min, max int) *Col
```

NewColForRange return a pointer to a new Col, which will apply to columns in the range min to max (inclusive). Note, in order for this Col to do anything useful you must set some of its parameters and then apply it to a Sheet by calling sheet.SetColParameters. Column numbers start from 1.

#### func (\*Col) [GetStyle](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/col.go#L75) [¶](#Col.GetStyle "Go to Col.GetStyle")

```
func (c *Col) GetStyle() *Style
```

GetStyle returns the Style associated with a Col

#### func (\*Col) [SetOutlineLevel](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/col.go#L84) [¶](#Col.SetOutlineLevel "Go to Col.SetOutlineLevel")

```
func (c *Col) SetOutlineLevel(outlineLevel uint8)
```

#### func (\*Col) [SetStyle](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/col.go#L80) [¶](#Col.SetStyle "Go to Col.SetStyle")

```
func (c *Col) SetStyle(style *Style)
```

SetStyle sets the style of a Col

#### func (\*Col) [SetType](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/col.go#L53) [¶](#Col.SetType "Go to Col.SetType")

```
func (c *Col) SetType(cellType CellType)
```

SetType will set the format string of a column based on the type that you want to set it to. This function does not really make a lot of sense.

#### func (\*Col) [SetWidth](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/col.go#L45) [¶](#Col.SetWidth "Go to Col.SetWidth")

```
func (c *Col) SetWidth(width float64)
```

SetWidth sets the width of columns that have this Col applied to them. The width is expressed as the number of characters of the maximum digit width of the numbers 0-9 as rendered in the normal style's font.

#### type [ColStore](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/col.go#L142) [¶](#ColStore "Go to ColStore")

```
type ColStore struct {
	Root *ColStoreNode
	Len  int
}
```

ColStore is the working store of Col definitions, it will simplify all Cols added to it, to ensure there ar no overlapping definitions.

#### func (\*ColStore) [Add](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/col.go#L150) [¶](#ColStore.Add "Go to ColStore.Add")

```
func (cs *ColStore) Add(col *Col) *ColStoreNode
```

Add a Col to the ColStore. If it overwrites all, or part of some existing Col's range of columns the that Col will be adjusted and/or split to make room for the new Col.

#### func (\*ColStore) [FindColByIndex](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/col.go#L161) [¶](#ColStore.FindColByIndex "Go to ColStore.FindColByIndex")

```
func (cs *ColStore) FindColByIndex(index int) *Col
```

#### func (\*ColStore) [ForEach](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/col.go#L397) [¶](#ColStore.ForEach "Go to ColStore.ForEach")

```
func (cs *ColStore) ForEach(fn func(idx int, col *Col))
```

ForEach calls the function fn for each Col defined in the ColStore.

#### type [ColStoreNode](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/col.go#L109) [¶](#ColStoreNode "Go to ColStoreNode")

```
type ColStoreNode struct {
	Col  *Col
	Prev *ColStoreNode
	Next *ColStoreNode
}
```

#### type [DataValidationErrorStyle](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/data_validation.go#L33) [¶](#DataValidationErrorStyle "Go to DataValidationErrorStyle")

```
type DataValidationErrorStyle int
```

```
const (
	StyleStop DataValidationErrorStyle
	StyleWarning
	StyleInformation
)
```

Data validation error styles

#### type [DataValidationOperator](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/data_validation.go#L51) [¶](#DataValidationOperator "Go to DataValidationOperator")

```
type DataValidationOperator int
```

DataValidationOperator operator enum

#### type [DataValidationType](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/data_validation.go#L10) [¶](#DataValidationType "Go to DataValidationType")

```
type DataValidationType int
```

#### type [DateTimeOptions](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L297) [¶](#DateTimeOptions "Go to DateTimeOptions")

```
type DateTimeOptions struct {
	// Location allows calculating times in other timezones/locations
	Location *time.Location
	// ExcelTimeFormat is the string you want excel to use to format the datetime
	ExcelTimeFormat string
}
```

DateTimeOptions are additional options for exporting times

#### type [DefinedName](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L734) [¶](#DefinedName "Go to DefinedName")

```
type DefinedName xlsxDefinedName
```

#### type [DiskVCellStore](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L338) [¶](#DiskVCellStore "Go to DiskVCellStore")

```
type DiskVCellStore struct {
	// contains filtered or unexported fields
}
```

DiskVCellStore is an implementation of the CellStore interface, backed by DiskV

#### func (\*DiskVCellStore) [Close](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L489) [¶](#DiskVCellStore.Close "Go to DiskVCellStore.Close")

```
func (cs *DiskVCellStore) Close() error
```

Close will remove the persisant storage for a given Sheet completely.

#### func (\*DiskVCellStore) [MakeRow](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L477) [¶](#DiskVCellStore.MakeRow "Go to DiskVCellStore.MakeRow")

```
func (cs *DiskVCellStore) MakeRow(sheet *Sheet) *Row
```

MakeRow returns an empty Row

#### func (\*DiskVCellStore) [MakeRowWithLen](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L482) [¶](#DiskVCellStore.MakeRowWithLen "Go to DiskVCellStore.MakeRowWithLen")

```
func (cs *DiskVCellStore) MakeRowWithLen(sheet *Sheet, len int) *Row
```

MakeRowWithLen returns an empty Row, with a preconfigured starting length.

#### func (\*DiskVCellStore) [MoveRow](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L392) [¶](#DiskVCellStore.MoveRow "Go to DiskVCellStore.MoveRow")

```
func (cs *DiskVCellStore) MoveRow(r *Row, index int) error
```

MoveRow moves a Row from one position in a Sheet (index) to another within the persistant store.

#### func (\*DiskVCellStore) [ReadRow](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L375) [¶](#DiskVCellStore.ReadRow "Go to DiskVCellStore.ReadRow")

```
func (cs *DiskVCellStore) ReadRow(key string, s *Sheet) (*Row, error)
```

ReadRow reads a row from the persistant store, identified by key, into memory and returns it, with the provided Sheet set as the Row's Sheet.

#### func (\*DiskVCellStore) [RemoveRow](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L465) [¶](#DiskVCellStore.RemoveRow "Go to DiskVCellStore.RemoveRow")

```
func (cs *DiskVCellStore) RemoveRow(key string) error
```

RemoveRow removes a Row from the Sheet's representation in the persistant store.

#### func (\*DiskVCellStore) [WriteRow](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L1237) [¶](#DiskVCellStore.WriteRow "Go to DiskVCellStore.WriteRow")

```
func (cs *DiskVCellStore) WriteRow(r *Row) error
```

WriteRow writes a Row to persistant storage.

#### type [DiskVRow](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L30) [¶](#DiskVRow "Go to DiskVRow")

```
type DiskVRow struct {
	// contains filtered or unexported fields
}
```

#### func (\*DiskVRow) [AddCell](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L62) [¶](#DiskVRow.AddCell "Go to DiskVRow.AddCell")

```
func (dvr *DiskVRow) AddCell() *Cell
```

#### func (\*DiskVRow) [CellCount](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L333) [¶](#DiskVRow.CellCount "Go to DiskVRow.CellCount")

```
func (dvr *DiskVRow) CellCount() int
```

CellCount returns the total number of cells in the row.

#### func (\*DiskVRow) [CellUpdatable](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L50) [¶](#DiskVRow.CellUpdatable "Go to DiskVRow.CellUpdatable")

```
func (dvr *DiskVRow) CellUpdatable(c *Cell)
```

#### func (\*DiskVRow) [ForEachCell](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L270) [¶](#DiskVRow.ForEachCell "Go to DiskVRow.ForEachCell")

```
func (dvr *DiskVRow) ForEachCell(cvf CellVisitorFunc, option ...CellVisitorOption) error
```

#### func (\*DiskVRow) [GetCell](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L253) [¶](#DiskVRow.GetCell "Go to DiskVRow.GetCell")

```
func (dvr *DiskVRow) GetCell(colIdx int) *Cell
```

#### func (\*DiskVRow) [MaxCol](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L328) [¶](#DiskVRow.MaxCol "Go to DiskVRow.MaxCol")

```
func (dvr *DiskVRow) MaxCol() int
```

MaxCol returns the index of the rightmost cell in the row's column.

#### func (\*DiskVRow) [PushCell](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L248) [¶](#DiskVRow.PushCell "Go to DiskVRow.PushCell")

```
func (dvr *DiskVRow) PushCell(c *Cell)
```

#### func (\*DiskVRow) [Updatable](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go#L56) [¶](#DiskVRow.Updatable "Go to DiskVRow.Updatable")

```
func (dvr *DiskVRow) Updatable()
```

#### type [File](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L17) [¶](#File "Go to File")

```
type File struct {
	Date1904 bool

	Sheets []*Sheet
	Sheet  map[string]*Sheet

	DefinedNames []*xlsxDefinedName
	// contains filtered or unexported fields
}
```

File is a high level structure providing a slice of Sheet structs to the user.

#### func [NewFile](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L66) [¶](#NewFile "Go to NewFile")

```
func NewFile(options ...FileOption) *File
```

NewFile creates a new File struct. You may pass it zero, one or many FileOption functions that affect the behaviour of the file.

#### func [OpenBinary](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L105) [¶](#OpenBinary "Go to OpenBinary")

```
func OpenBinary(bs []byte, options ...FileOption) (*File, error)
```

OpenBinary() take bytes of an XLSX file and returns a populated xlsx.File struct for it.

#### func [OpenFile](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L84) [¶](#OpenFile "Go to OpenFile")

```
func OpenFile(fileName string, options ...FileOption) (file *File, err error)
```

OpenFile will take the name of an XLSX file and returns a populated xlsx.File struct for it. You may pass it zero, one or many FileOption functions that affect the behaviour of the file.

#### func [OpenReaderAt](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L113) [¶](#OpenReaderAt "Go to OpenReaderAt")

```
func OpenReaderAt(r io.ReaderAt, size int64, options ...FileOption) (*File, error)
```

OpenReaderAt() take io.ReaderAt of an XLSX file and returns a populated xlsx.File struct for it.

#### func [ReadZip](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/lib.go#L1135) [¶](#ReadZip "Go to ReadZip")

```
func ReadZip(f *zip.ReadCloser, options ...FileOption) (*File, error)
```

ReadZip() takes a pointer to a zip.ReadCloser and returns a xlsx.File struct populated with its contents. In most cases ReadZip is not used directly, but is called internally by OpenFile.

#### func [ReadZipReader](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/lib.go#L1146) [¶](#ReadZipReader "Go to ReadZipReader")

```
func ReadZipReader(r *zip.Reader, options ...FileOption) (*File, error)
```

ReadZipReader() can be used to read an XLSX in memory without touching the filesystem.

#### func (\*File) [AddDefinedName](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L737) [¶](#File.AddDefinedName "Go to File.AddDefinedName")

```
func (f *File) AddDefinedName(name DefinedName) error
```

AddDefinedName adds a new Name definition to the workbook.

#### func (\*File) [AddSheet](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L196) [¶](#File.AddSheet "Go to File.AddSheet")

```
func (f *File) AddSheet(sheetName string) (*Sheet, error)
```

AddSheet Add a new Sheet, with the provided name, to a File. The minimum sheet name length is 1 character. If the sheet name length is less an error is thrown. The maximum sheet name length is 31 characters. If the sheet name length is exceeded an error is thrown. These special characters are also not allowed: : \\ / ? * \[ ]

#### func (\*File) [AddSheetWithCellStore](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L200) [¶](#File.AddSheetWithCellStore "Go to File.AddSheetWithCellStore")

```
func (f *File) AddSheetWithCellStore(sheetName string, constructor CellStoreConstructor) (*Sheet, error)
```

#### func (\*File) [AppendSheet](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L227) [¶](#File.AppendSheet "Go to File.AppendSheet")

```
func (f *File) AppendSheet(sheet Sheet, sheetName string) (*Sheet, error)
```

Appends an existing Sheet, with the provided name, to a File

#### func (\*File) [MakeStreamParts](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L336) [¶](#File.MakeStreamParts "Go to File.MakeStreamParts")

```
func (f *File) MakeStreamParts() (map[string]string, error)
```

MakeStreamParts constructs a map of file name to XML content representing the file in terms of the structure of an XLSX file.

#### func (\*File) [MarshallParts](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L467) [¶](#File.MarshallParts "Go to File.MarshallParts")

```
func (f *File) MarshallParts(zipWriter *zip.Writer) error
```

MarshallParts constructs a map of file name to XML content representing the file in terms of the structure of an XLSX file.

#### func (\*File) [Save](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L156) [¶](#File.Save "Go to File.Save")

```
func (f *File) Save(path string) (err error)
```

Save the File to an xlsx file at the provided path.

#### func (\*File) [ToSlice](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L652) [¶](#File.ToSlice "Go to File.ToSlice")

```
func (f *File) ToSlice() (output [][][]string, err error)
```

Return the raw data contained in the File as three dimensional slice. The first index represents the sheet number, the second the row number, and the third the cell number.

For example:

```
var mySlice [][][]string
var value string
mySlice = xlsx.FileToSlice("myXLSX.xlsx")
value = mySlice[0][0][0]
```

Here, value would be set to the raw value of the cell A1 in the first sheet in the XLSX file.

#### func (\*File) [ToSliceUnmerged](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L703) [¶](#File.ToSliceUnmerged "Go to File.ToSliceUnmerged")

```
func (f *File) ToSliceUnmerged() (output [][][]string, err error)
```

ToSliceUnmerged returns the raw data contained in the File as three dimensional slice (s. method ToSlice). A covered cell become the value of its origin cell. Example: table where A1:A2 at row 0 and row 1 are merged. | 2011 | Bread | 20 | | | Fish | 70 | | 2012 | 2013 | Egg | 80 | This sheet will be converted to the slice: [

```
[2011 2011 Bread 20]
[2011 2011 Fish  70]
[2012 2013 Egg   80]
```

]

#### func (\*File) [Write](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L176) [¶](#File.Write "Go to File.Write")

```
func (f *File) Write(writer io.Writer) error
```

Write the File to io.Writer as xlsx

#### type [FileOption](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L36) [¶](#FileOption "Go to FileOption")

```
type FileOption func(f *File)
```

#### func [ColLimit](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L48) [¶](#ColLimit "Go to ColLimit")

```
func ColLimit(n int) FileOption
```

ColLimit will limit the columns handled in any given sheet to the first n, where n is the number of columns

#### func [RowLimit](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L40) [¶](#RowLimit "Go to RowLimit")

```
func RowLimit(n int) FileOption
```

RowLimit will limit the rows handled in any given sheet to the first n, where n is the number of rows.

#### func [ValueOnly](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go#L58) [¶](#ValueOnly "Go to ValueOnly")

```
func ValueOnly() FileOption
```

ValueOnly treats all NULL values as meaningless and it will delete all NULL value cells, before decode worksheet.xml. this option can save memory and time when parsing files with a large number of NULL values. But it may also cause accidental injury, because NULL may not really be meaningless. Use with caution

#### type [Fill](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L149) [¶](#Fill "Go to Fill")

```
type Fill struct {
	PatternType string
	BgColor     string
	FgColor     string
}
```

Fill is a high level structure intended to provide user access to the contents of background and foreground color index within an Sheet.

#### func [DefaultFill](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L200) [¶](#DefaultFill "Go to DefaultFill")

```
func DefaultFill() *Fill
```

#### func [NewFill](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L155) [¶](#NewFill "Go to NewFill")

```
func NewFill(patternType, fgColor, bgColor string) *Fill
```

#### type [Font](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L163) [¶](#Font "Go to Font")

```
type Font struct {
	Size      float64
	Name      string
	Family    int
	Charset   int
	Color     string
	Bold      bool
	Italic    bool
	Underline bool
	Strike    bool
}
```

#### func [DefaultFont](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L196) [¶](#DefaultFont "Go to DefaultFont")

```
func DefaultFont() *Font
```

#### func [NewFont](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L175) [¶](#NewFont "Go to NewFont")

```
func NewFont(size float64, name string) *Font
```

#### type [HSL](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/hsl.go#L44) [¶](#HSL "Go to HSL")

```
type HSL struct {
	H, S, L float64
}
```

HSL represents a cylindrical coordinate of points in an RGB color model.

Values are in the range 0 to 1.

#### func (HSL) [RGBA](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/hsl.go#L50) [¶](#HSL.RGBA "Go to HSL.RGBA")

```
func (c HSL) RGBA() (uint32, uint32, uint32, uint32)
```

RGBA returns the alpha-premultiplied red, green, blue and alpha values for the HSL.

#### type [Hyperlink](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go#L176) [¶](#Hyperlink "Go to Hyperlink")

```
type Hyperlink struct {
	DisplayString string
	Link          string
	Tooltip       string
	Location      string
}
```

Hyperlink is a structure to store link information in-workbook links to cells or defined names are stored in Location external links are stores in Link

#### type [MemoryCellStore](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L135) [¶](#MemoryCellStore "Go to MemoryCellStore")

```
type MemoryCellStore struct {
	// contains filtered or unexported fields
}
```

MemoryCellStore is the default CellStore - it holds all rows and cells in system memory. This is fast, right up until you run out of memory ;-)

#### func (\*MemoryCellStore) [Close](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L160) [¶](#MemoryCellStore.Close "Go to MemoryCellStore.Close")

```
func (mcs *MemoryCellStore) Close() error
```

Close is nullOp for the MemoryCellStore, but we have to comply with the interface.

#### func (\*MemoryCellStore) [MakeRow](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L215) [¶](#MemoryCellStore.MakeRow "Go to MemoryCellStore.MakeRow")

```
func (mcs *MemoryCellStore) MakeRow(sheet *Sheet) *Row
```

MakeRow returns an empty Row

#### func (\*MemoryCellStore) [MakeRowWithLen](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L207) [¶](#MemoryCellStore.MakeRowWithLen "Go to MemoryCellStore.MakeRowWithLen")

```
func (mcs *MemoryCellStore) MakeRowWithLen(sheet *Sheet, len int) *Row
```

MakeRowWithLen returns an empty Row, with a preconfigured starting length.

#### func (\*MemoryCellStore) [MoveRow](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L183) [¶](#MemoryCellStore.MoveRow "Go to MemoryCellStore.MoveRow")

```
func (mcs *MemoryCellStore) MoveRow(r *Row, index int) error
```

MoveRow moves the persisted Row's position in the sheet.

#### func (\*MemoryCellStore) [ReadRow](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L165) [¶](#MemoryCellStore.ReadRow "Go to MemoryCellStore.ReadRow")

```
func (mcs *MemoryCellStore) ReadRow(key string, s *Sheet) (*Row, error)
```

ReadRow returns a Row identfied by the given key.

#### func (\*MemoryCellStore) [RemoveRow](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L197) [¶](#MemoryCellStore.RemoveRow "Go to MemoryCellStore.RemoveRow")

```
func (mcs *MemoryCellStore) RemoveRow(key string) error
```

RemoveRow removes a row from the sheet, it doesn't specifically move any following rows, leaving this decision to the user.

#### func (\*MemoryCellStore) [WriteRow](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L174) [¶](#MemoryCellStore.WriteRow "Go to MemoryCellStore.WriteRow")

```
func (mcs *MemoryCellStore) WriteRow(r *Row) error
```

WriteRow pushes the Row to the MemoryCellStore.

#### type [MemoryRow](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L7) [¶](#MemoryRow "Go to MemoryRow")

```
type MemoryRow struct {
	// contains filtered or unexported fields
}
```

#### func (\*MemoryRow) [AddCell](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L32) [¶](#MemoryRow.AddCell "Go to MemoryRow.AddCell")

```
func (mr *MemoryRow) AddCell() *Cell
```

#### func (\*MemoryRow) [CellCount](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L128) [¶](#MemoryRow.CellCount "Go to MemoryRow.CellCount")

```
func (mr *MemoryRow) CellCount() int
```

CellCount returns the total number of cells in the row.

#### func (\*MemoryRow) [CellUpdatable](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L28) [¶](#MemoryRow.CellUpdatable "Go to MemoryRow.CellUpdatable")

```
func (mr *MemoryRow) CellUpdatable(c *Cell)
```

#### func (\*MemoryRow) [ForEachCell](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L81) [¶](#MemoryRow.ForEachCell "Go to MemoryRow.ForEachCell")

```
func (mr *MemoryRow) ForEachCell(cvf CellVisitorFunc, option ...CellVisitorOption) error
```

#### func (\*MemoryRow) [GetCell](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L64) [¶](#MemoryRow.GetCell "Go to MemoryRow.GetCell")

```
func (mr *MemoryRow) GetCell(colIdx int) *Cell
```

#### func (\*MemoryRow) [MaxCol](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L123) [¶](#MemoryRow.MaxCol "Go to MemoryRow.MaxCol")

```
func (mr *MemoryRow) MaxCol() int
```

MaxCol returns the index of the rightmost cell in the row's column.

#### func (\*MemoryRow) [PushCell](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L38) [¶](#MemoryRow.PushCell "Go to MemoryRow.PushCell")

```
func (mr *MemoryRow) PushCell(c *Cell)
```

#### func (\*MemoryRow) [Updatable](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go#L24) [¶](#MemoryRow.Updatable "Go to MemoryRow.Updatable")

```
func (mr *MemoryRow) Updatable()
```

#### type [Pane](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L80) [¶](#Pane "Go to Pane")

```
type Pane struct {
	XSplit      float64
	YSplit      float64
	TopLeftCell string
	ActivePane  string
	State       string // Either "split" or "frozen"
}
```

#### type [RefTable](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/reftable.go#L11) [¶](#RefTable "Go to RefTable")

```
type RefTable struct {
	// contains filtered or unexported fields
}
```

#### func [NewSharedStringRefTable](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/reftable.go#L18) [¶](#NewSharedStringRefTable "Go to NewSharedStringRefTable")

```
func NewSharedStringRefTable(size int) *RefTable
```

NewSharedStringRefTable creates a new, empty RefTable.

#### func (\*RefTable) [AddRichText](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/reftable.go#L99) [¶](#RefTable.AddRichText "Go to RefTable.AddRichText")

```
func (rt *RefTable) AddRichText(r []RichTextRun) int
```

AddRichText adds a set of rich text to the reference table and return it's numeric index. If a set of rich text already exists then it simply returns the existing index.

#### func (\*RefTable) [AddString](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/reftable.go#L78) [¶](#RefTable.AddString "Go to RefTable.AddString")

```
func (rt *RefTable) AddString(str string) int
```

AddString adds a string to the reference table and return it's numeric index. If the string already exists then it simply returns the existing index.

#### func (\*RefTable) [Length](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/reftable.go#L132) [¶](#RefTable.Length "Go to RefTable.Length")

```
func (rt *RefTable) Length() int
```

#### func (\*RefTable) [ResolveSharedString](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/reftable.go#L65) [¶](#RefTable.ResolveSharedString "Go to RefTable.ResolveSharedString")

```
func (rt *RefTable) ResolveSharedString(index int) (plainText string, richText []RichTextRun)
```

ResolveSharedString looks up a string value or the rich text by numeric index from a provided reference table (just a slice of strings in the correct order). If the rich text was found, non-empty slice will be returned in richText. This function only exists to provide clarity of purpose via it's name.

#### type [Relation](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L100) [¶](#Relation "Go to Relation")

```
type Relation struct {
	Type       RelationshipType
	Target     string
	TargetMode RelationshipTargetMode
}
```

#### type [RelationshipTargetMode](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/xlsxRelation.go#L11) [¶](#RelationshipTargetMode "Go to RelationshipTargetMode")

```
type RelationshipTargetMode string
```

```
const (
	RelationshipTargetModeExternal RelationshipTargetMode = "External"
)
```

#### type [RelationshipType](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/xlsxRelation.go#L5) [¶](#RelationshipType "Go to RelationshipType")

```
type RelationshipType string
```

```
const (
	RelationshipTypeHyperlink RelationshipType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/hyperlink"
)
```

#### type [RichTextCharset](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/richtext.go#L9) [¶](#RichTextCharset "Go to RichTextCharset")

```
type RichTextCharset int
```

#### type [RichTextColor](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/richtext.go#L58) [¶](#RichTextColor "Go to RichTextColor")

```
type RichTextColor struct {
	// contains filtered or unexported fields
}
```

RichTextColor is the color of the RichTextRun.

#### func [NewRichTextColorFromARGB](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/richtext.go#L64) [¶](#NewRichTextColorFromARGB "Go to NewRichTextColorFromARGB")

```
func NewRichTextColorFromARGB(alpha, red, green, blue int) *RichTextColor
```

NewRichTextColorFromARGB creates a new RichTextColor from ARGB component values. Each component must have a value in range of 0 to 255.

#### func [NewRichTextColorFromThemeColor](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/richtext.go#L71) [¶](#NewRichTextColorFromThemeColor "Go to NewRichTextColorFromThemeColor")

```
func NewRichTextColorFromThemeColor(themeColor int) *RichTextColor
```

NewRichTextColorFromThemeColor creates a new RichTextColor from the theme color. The argument \`themeColor\` is a zero-based index of the theme color.

#### type [RichTextFont](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/richtext.go#L76) [¶](#RichTextFont "Go to RichTextFont")

```
type RichTextFont struct {
	// Name is the font name. If Name is empty, Size, Family and Charset will be ignored.
	Name string
	// Size is the font size.
	Size float64
	// Family is a value of the font family. Use one of the RichTextFontFamily constants.
	Family RichTextFontFamily
	// Charset is a value of the charset of the font. Use one of the RichTextCharset constants.
	Charset RichTextCharset
	// Color is the text color.
	Color *RichTextColor
	// Bold specifies the bold face font style.
	Bold bool
	// Italic specifies the italic font style.
	Italic bool
	// Strike specifies a strikethrough line.
	Strike bool
	// VertAlign specifies the vertical position of the text. Use one of the RichTextVertAlign constants, or empty.
	VertAlign RichTextVertAlign
	// Underline specifies the underline style. Use one of the RichTextUnderline constants, or empty.
	Underline RichTextUnderline
}
```

RichTextFont is the font spec of the RichTextRun.

#### type [RichTextFontFamily](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/richtext.go#L8) [¶](#RichTextFontFamily "Go to RichTextFontFamily")

```
type RichTextFontFamily int
```

#### type [RichTextRun](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/richtext.go#L100) [¶](#RichTextRun "Go to RichTextRun")

```
type RichTextRun struct {
	Font *RichTextFont
	Text string
}
```

RichTextRun is a run of the decorated text.

#### func (\*RichTextRun) [Equals](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/richtext.go#L105) [¶](#RichTextRun.Equals "Go to RichTextRun.Equals")

```
func (rt *RichTextRun) Equals(other *RichTextRun) bool
```

#### type [RichTextUnderline](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/richtext.go#L11) [¶](#RichTextUnderline "Go to RichTextUnderline")

```
type RichTextUnderline string
```

#### type [RichTextVertAlign](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/richtext.go#L10) [¶](#RichTextVertAlign "Go to RichTextVertAlign")

```
type RichTextVertAlign string
```

#### type [Row](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/row.go#L8) [¶](#Row "Go to Row")

```
type Row struct {
	Hidden bool   // Hidden determines whether this Row is hidden or not.
	Sheet  *Sheet // Sheet is a reference back to the Sheet that this Row is within.
	// contains filtered or unexported fields
}
```

Row represents a single Row in the current Sheet.

#### func (\*Row) [AddCell](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/row.go#L69) [¶](#Row.AddCell "Go to Row.AddCell")

```
func (r *Row) AddCell() *Cell
```

AddCell adds a new Cell to the end of the Row

#### func (\*Row) [ForEachCell](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/row.go#L125) [¶](#Row.ForEachCell "Go to Row.ForEachCell")

```
func (r *Row) ForEachCell(cvf CellVisitorFunc, option ...CellVisitorOption) error
```

ForEachCell will call the provided [CellVisitorFunc](#CellVisitorFunc) for each currently defined cell in the Row. Optionally you may pass one or more [CellVisitorOption](#CellVisitorOption) to affect how ForEachCell operates. For example you may wish to pass [SkipEmptyCells](#SkipEmptyCells) to only visit cells which are populated.

#### func (\*Row) [GetCell](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/row.go#L99) [¶](#Row.GetCell "Go to Row.GetCell")

```
func (r *Row) GetCell(colIdx int) *Cell
```

GetCell returns the Cell at a given column index, creating it if it doesn't exist.

#### func (\*Row) [GetCoordinate](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/row.go#L20) [¶](#Row.GetCoordinate "Go to Row.GetCoordinate")

```
func (r *Row) GetCoordinate() int
```

GetCoordinate returns the y coordinate of the row (the row number). This number is zero based, i.e. the Excel CellID "A1" is in Row 0, not Row 1.

#### func (\*Row) [GetHeight](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/row.go#L47) [¶](#Row.GetHeight "Go to Row.GetHeight")

```
func (r *Row) GetHeight() float64
```

GetHeight returns the height of the Row in PostScript points.

#### func (\*Row) [GetOutlineLevel](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/row.go#L64) [¶](#Row.GetOutlineLevel "Go to Row.GetOutlineLevel")

```
func (r *Row) GetOutlineLevel() uint8
```

GetOutlineLevel returns the outline level of the Row.

#### func (\*Row) [PushCell](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/row.go#L80) [¶](#Row.PushCell "Go to Row.PushCell")

```
func (r *Row) PushCell(c *Cell)
```

PushCell adds a predefiend cell to the end of the Row

#### func (\*Row) [ReadStruct](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/read.go#L26) [¶](#Row.ReadStruct "Go to Row.ReadStruct")

```
func (r *Row) ReadStruct(ptr interface{}) error
```

ReadStruct reads a struct from r to ptr. Accepts a ptr to struct. This code expects a tag xlsx:"N", where N is the index of the cell to be used. Basic types like int,string,float64 and bool are supported

Example [¶](#example-Row.ReadStruct "Go to Example")

```
//example type
type structTest struct {
	IntVal     int     `xlsx:"0"`
	StringVal  string  `xlsx:"1"`
	FloatVal   float64 `xlsx:"2"`
	IgnoredVal int     `xlsx:"-"`
	BoolVal    bool    `xlsx:"4"`
}
structVal := structTest{
	IntVal:     16,
	StringVal:  "heyheyhey :)!",
	FloatVal:   3.14159216,
	IgnoredVal: 7,
	BoolVal:    true,
}
//create a new xlsx file and write a struct
//in a new row
f := NewFile()
sheet, _ := f.AddSheet("TestRead")
row := sheet.AddRow()
row.WriteStruct(&structVal, -1)

//read the struct from the same row
readStruct := &structTest{}
err := row.ReadStruct(readStruct)
if err != nil {
	panic(err)
}
fmt.Println(readStruct)
```

#### func (\*Row) [SetHeight](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/row.go#L33) [¶](#Row.SetHeight "Go to Row.SetHeight")

```
func (r *Row) SetHeight(ht float64)
```

SetHeight sets the height of the Row in PostScript points

#### func (\*Row) [SetHeightCM](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/row.go#L41) [¶](#Row.SetHeightCM "Go to Row.SetHeightCM")

```
func (r *Row) SetHeightCM(ht float64)
```

SetHeightCM sets the height of the Row in centimetres, inherently converting it to PostScript points.

#### func (\*Row) [SetOutlineLevel](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/row.go#L52) [¶](#Row.SetOutlineLevel "Go to Row.SetOutlineLevel")

```
func (r *Row) SetOutlineLevel(outlineLevel uint8)
```

SetOutlineLevel sets the outline level of the Row (used for collapsing rows)

#### func (\*Row) [WriteSlice](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/write.go#L14) [¶](#Row.WriteSlice "Go to Row.WriteSlice")

```
func (r *Row) WriteSlice(e interface{}, cols int) int
```

Writes a slice to row r. Accepts a slice or a pointer to a slice, and will wirte up to the provided number of columns, 'cols'. If 'cols' is &lt; 0, the entire slice will be written if possible. Returns -1 if the 'e' is not a slice type, otherwise the number of columns written.

#### func (\*Row) [WriteStruct](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/write.go#L100) [¶](#Row.WriteStruct "Go to Row.WriteStruct")

```
func (r *Row) WriteStruct(e interface{}, cols int) int
```

Writes a struct to row r. Accepts a pointer to struct type 'e', and the number of columns to write, \`cols\`. If 'cols' is &lt; 0, the entire struct will be written if possible. Returns -1 if the 'e' doesn't point to a struct, otherwise the number of columns written

#### type [RowNotFoundError](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cellstore.go#L50) [¶](#RowNotFoundError "Go to RowNotFoundError")

```
type RowNotFoundError struct {
	// contains filtered or unexported fields
}
```

RowNotFoundError is an Error that should be returned by a RowStore implementation if a call to ReadRow is made with a key that doesn't correspond to any persisted Row.

#### func [NewRowNotFoundError](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cellstore.go#L56) [¶](#NewRowNotFoundError "Go to NewRowNotFoundError")

```
func NewRowNotFoundError(key, reason string) *RowNotFoundError
```

NewRowNotFoundError creates a new RowNotFoundError, capturing the Row key and the reason this key could not be found.

#### func (RowNotFoundError) [Error](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cellstore.go#L61) [¶](#RowNotFoundError.Error "Go to RowNotFoundError.Error")

```
func (cnfe RowNotFoundError) Error() string
```

Error returns a human-readable description of the failure to find a Row. It makes RowNotFoundError comply with the Error interface.

#### type [RowVisitor](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L157) [¶](#RowVisitor "Go to RowVisitor")

```
type RowVisitor func(r *Row) error
```

A RowVisitor function should be provided by the user when calling Sheet.ForEachRow, it will be called once for every Row visited.

#### type [RowVisitorOption](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L147) [¶](#RowVisitorOption "Go to RowVisitorOption")

```
type RowVisitorOption func(flags *rowVisitorFlags)
```

RowVisitorOption defines the call signature of functions that can be passed as options to the Sheet.ForEachRow function to affect its behaviour.

#### type [Sheet](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L17) [¶](#Sheet "Go to Sheet")

```
type Sheet struct {
	Name            string
	File            *File
	Cols            *ColStore
	MaxRow          int
	MaxCol          int
	Hidden          bool
	Selected        bool
	SheetViews      []SheetView
	SheetFormat     SheetFormat
	AutoFilter      *AutoFilter
	Relations       []Relation
	DataValidations []*xlsxDataValidation
	// contains filtered or unexported fields
}
```

Sheet is a high level structure intended to provide user access to the contents of a particular sheet within an XLSX file.

#### func [NewSheet](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L39) [¶](#NewSheet "Go to NewSheet")

```
func NewSheet(name string) (*Sheet, error)
```

NewSheet constructs a Sheet with the default CellStore and returns a pointer to it.

#### func [NewSheetWithCellStore](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L45) [¶](#NewSheetWithCellStore "Go to NewSheetWithCellStore")

```
func NewSheetWithCellStore(name string, constructor CellStoreConstructor) (*Sheet, error)
```

NewSheetWithCellStore constructs a Sheet, backed by a CellStore, for which you must provide the constructor function.

#### func (\*Sheet) [AddDataValidation](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L254) [¶](#Sheet.AddDataValidation "Go to Sheet.AddDataValidation")

```
func (s *Sheet) AddDataValidation(dv *xlsxDataValidation)
```

Add a DataValidation to a range of cells

#### func (\*Sheet) [AddRow](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L204) [¶](#Sheet.AddRow "Go to Sheet.AddRow")

```
func (s *Sheet) AddRow() *Row
```

Add a new Row to a Sheet

#### func (\*Sheet) [AddRowAtIndex](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L222) [¶](#Sheet.AddRowAtIndex "Go to Sheet.AddRowAtIndex")

```
func (s *Sheet) AddRowAtIndex(index int) (*Row, error)
```

Add a new Row to a Sheet at a specific index

#### func (\*Sheet) [Cell](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L345) [¶](#Sheet.Cell "Go to Sheet.Cell")

```
func (s *Sheet) Cell(row, col int) (*Cell, error)
```

Get a Cell by passing it's cartesian coordinates (zero based) as row and column integer indexes.

For example:

```
cell := sheet.Cell(0,0)
```

... would set the variable "cell" to contain a Cell struct containing the data from the field "A1" on the spreadsheet.

#### func (\*Sheet) [Close](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L64) [¶](#Sheet.Close "Go to Sheet.Close")

```
func (s *Sheet) Close()
```

Remove Sheet's dependant resources - if you are done with operations on a sheet this should be called to clear down the Sheet's persistent cache. Note: if you call this, all further read operaton on the sheet will fail - including any attempt to save the file, or dump it's contents to a byte stream. Therefore only call this \*after* you've saved your changes, of when you're done reading a sheet in a file you don't plan to persist.

#### func (\*Sheet) [Col](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L328) [¶](#Sheet.Col "Go to Sheet.Col")

```
func (s *Sheet) Col(idx int) *Col
```

Return the Col that applies to this Column index, or return nil if no such Col exists Column numbers start from 1.

#### func (\*Sheet) [ForEachRow](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L165) [¶](#Sheet.ForEachRow "Go to Sheet.ForEachRow")

```
func (s *Sheet) ForEachRow(rv RowVisitor, options ...RowVisitorOption) error
```

#### func (\*Sheet) [MarshalSheet](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L888) [¶](#Sheet.MarshalSheet "Go to Sheet.MarshalSheet")

```
func (s *Sheet) MarshalSheet(w io.Writer, refTable *RefTable, styles *xlsxStyleSheet, relations *xlsxWorksheetRels) error
```

#### func (\*Sheet) [RemoveRowAtIndex](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L260) [¶](#Sheet.RemoveRowAtIndex "Go to Sheet.RemoveRowAtIndex")

```
func (s *Sheet) RemoveRowAtIndex(index int) error
```

Removes a row at a specific index

#### func (\*Sheet) [Row](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L299) [¶](#Sheet.Row "Go to Sheet.Row")

```
func (s *Sheet) Row(idx int) (*Row, error)
```

Make sure we always have as many Rows as we do cells.

#### func (\*Sheet) [SetColAutoWidth](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L436) [¶](#Sheet.SetColAutoWidth "Go to Sheet.SetColAutoWidth")

```
func (s *Sheet) SetColAutoWidth(colIndex int, width func(string) float64) error
```

Tries to guess the best width for a column, based on the largest cell content. A scale function needs to be provided. Column numbers start from 1.

#### func (\*Sheet) [SetColParameters](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L364) [¶](#Sheet.SetColParameters "Go to Sheet.SetColParameters")

```
func (s *Sheet) SetColParameters(col *Col)
```

Set the parameters of a column. Parameters are passed as a pointer to a Col structure which you much construct yourself. Column numbers start from 1.

#### func (\*Sheet) [SetColWidth](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L420) [¶](#Sheet.SetColWidth "Go to Sheet.SetColWidth")

```
func (s *Sheet) SetColWidth(min, max int, width float64)
```

Set the width of a range of columns. Column numbers start from 1.

#### func (\*Sheet) [SetOutlineLevel](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L463) [¶](#Sheet.SetOutlineLevel "Go to Sheet.SetOutlineLevel")

```
func (s *Sheet) SetOutlineLevel(minCol, maxCol int, outlineLevel uint8)
```

Set the outline level for a range of columns.

#### func (\*Sheet) [SetType](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L471) [¶](#Sheet.SetType "Go to Sheet.SetType")

```
func (s *Sheet) SetType(minCol, maxCol int, cellType CellType)
```

Set the type for a range of columns.

#### type [SheetFormat](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L88) [¶](#SheetFormat "Go to SheetFormat")

```
type SheetFormat struct {
	DefaultColWidth  float64
	DefaultRowHeight float64
	OutlineLevelCol  uint8
	OutlineLevelRow  uint8
}
```

#### type [SheetView](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go#L76) [¶](#SheetView "Go to SheetView")

```
type SheetView struct {
	Pane *Pane
}
```

#### type [Style](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L30) [¶](#Style "Go to Style")

```
type Style struct {
	Border          Border
	Fill            Fill
	Font            Font
	ApplyBorder     bool
	ApplyFill       bool
	ApplyFont       bool
	ApplyAlignment  bool
	Alignment       Alignment
	NamedStyleIndex *int
}
```

Style is a high level structure intended to provide user access to the contents of Style within an XLSX file.

#### func [NewStyle](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go#L43) [¶](#NewStyle "Go to NewStyle")

```
func NewStyle() *Style
```

Return a new Style structure initialised with the default values.

#### type [WorkBookRels](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/lib.go#L1054) [¶](#WorkBookRels "Go to WorkBookRels")

```
type WorkBookRels map[string]string
```

#### func (\*WorkBookRels) [MakeXLSXWorkbookRels](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/lib.go#L1056) [¶](#WorkBookRels.MakeXLSXWorkbookRels "Go to WorkBookRels.MakeXLSXWorkbookRels")

```
func (w *WorkBookRels) MakeXLSXWorkbookRels() xlsxWorkbookRels
```

#### type [XLSXReaderError](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/lib.go#L41) [¶](#XLSXReaderError "Go to XLSXReaderError")

```
type XLSXReaderError struct {
	Err string
}
```

XLSXReaderError is the standard error type for otherwise undefined errors in the XSLX reading process.

#### func (\*XLSXReaderError) [Error](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/lib.go#L47) [¶](#XLSXReaderError.Error "Go to XLSXReaderError.Error")

```
func (e *XLSXReaderError) Error() string
```

Error returns a string value from an XLSXReaderError struct in order that it might comply with the builtin.error interface.

#### type [XLSXUnmarshaler](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/read.go#L18) [¶](#XLSXUnmarshaler "Go to XLSXUnmarshaler")

```
type XLSXUnmarshaler interface {
	Unmarshal(*Row) error
}
```

XLSXUnmarshaler is the interface implemented for types that can unmarshal a Row as a representation of themselves.

## ![](/static/shared/icon/insert_drive_file_gm_grey_24dp.svg) Source Files [¶](#section-sourcefiles "Go to Source Files")

[View all Source files](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0)

- [cell.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cell.go "cell.go")
- [cellstore.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/cellstore.go "cellstore.go")
- [col.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/col.go "col.go")
- [data\_validation.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/data_validation.go "data_validation.go")
- [date.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/date.go "date.go")
- [diskv.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/diskv.go "diskv.go")
- [doc.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/doc.go "doc.go")
- [file.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/file.go "file.go")
- [format\_code.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/format_code.go "format_code.go")
- [hsl.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/hsl.go "hsl.go")
- [lib.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/lib.go "lib.go")
- [memory.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/memory.go "memory.go")
- [read.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/read.go "read.go")
- [reftable.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/reftable.go "reftable.go")
- [richtext.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/richtext.go "richtext.go")
- [row.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/row.go "row.go")
- [sheet.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/sheet.go "sheet.go")
- [style.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/style.go "style.go")
- [templates.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/templates.go "templates.go")
- [testutil.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/testutil.go "testutil.go")
- [theme.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/theme.go "theme.go")
- [utility.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/utility.go "utility.go")
- [write.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/write.go "write.go")
- [xlsxRelation.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/xlsxRelation.go "xlsxRelation.go")
- [xmlContentTypes.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/xmlContentTypes.go "xmlContentTypes.go")
- [xmlSharedStrings.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/xmlSharedStrings.go "xmlSharedStrings.go")
- [xmlStyle.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/xmlStyle.go "xmlStyle.go")
- [xmlTheme.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/xmlTheme.go "xmlTheme.go")
- [xmlWorkbook.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/xmlWorkbook.go "xmlWorkbook.go")
- [xmlWorksheet.go](https://codeberg.org/tealeg/xlsx.git/src/tag/v4.1.0/xmlWorksheet.go "xmlWorksheet.go")

Click to show internal directories.

Click to hide internal directories.
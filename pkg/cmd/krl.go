package cmd

import (
	"unicode/utf8"
)

//====================================================== the print tools
type AlignStyleType int
type RowValue []string

// The tables const
const (
	HorizontalDivider = "-" // the horizontal divider
	TitleDivider      = "="
	VerticalDivider   = "|" // the vertical divider
	CenterDivider     = "+" // the center divider

	CenterAlignType = iota // the align style : for center
	RightAlignType         // the align style : for right
	LeftAlignType          // the align style : for left

)

// Table struct save the title and value info, the tile count must equals to the info count. The one title is mapper the
// column in table, is has itself align style. For: CenterAlignType, RightAlignType and LeftAlignType.
// If the title is set, it can't be changed.
type Table struct {
	// values
	titles RowValue
	infos  []RowValue

	// data
	maxLength []int
	// the align style for every column.
	alignStyle []AlignStyleType
}

// SetAlignStyle will set align style from given styles, if is less then title count, the missing part will set to
// CenterAlignType, and the default style is CenterAlignType.
// If input is nil, will set all align style to CenterAlignType.
func (t *Table) SetAlignStyle(alignStyle []AlignStyleType) {
	if alignStyle == nil {
		alignStyle = make([]AlignStyleType, len(t.titles))
		// default align style is center align type
		for index, _ := range t.alignStyle {
			t.alignStyle[index] = CenterAlignType
		}
	} else if len(alignStyle) != len(t.titles) {
		newAlignStyle := make([]AlignStyleType, len(t.titles))
		for index, _ := range newAlignStyle {
			newAlignStyle[index] = CenterAlignType
		}
		for index, item := range alignStyle {
			newAlignStyle[index] = item
		}
		alignStyle = newAlignStyle
	}

	t.alignStyle = alignStyle
}

// SetIndexAlignStyle will set target column align style to target alignType, if given index is illegal, it will do
// nothing.
func (t *Table) SetIndexAlignStyle(index int, alignType AlignStyleType) {
	if index < 0 || index > len(t.alignStyle) {
		return
	}

	t.alignStyle[index] = alignType
}

// InitTitle will init the title for table
func (t *Table) InitTitle(titles []string) {
	// if titles is empty, it set empty string
	if titles == nil || len(titles) == 0 {
		t.titles = RowValue{}
		t.maxLength = []int{}
		return
	}

	// set title
	t.titles = titles
	t.maxLength = make([]int, len(t.titles))
	// set maxLength
	for index, title := range titles {
		t.maxLength[index] = utf8.RuneCountInString(title)
	}

	t.SetAlignStyle(nil)
}

// addRows will add 'value' to table
func (t *Table) addRows(value []string) {
	// empty value to add
	if value == nil || len(value) == 0 {
		t.infos = append(t.infos, make(RowValue, len(t.titles)))
		return
	}

	// check value to add
	if len(value) != len(t.titles) {
		// add empty info to value
		if len(value) < len(t.titles) {
			// less then title, add ""
			for len(value) != len(t.titles) {
				value = append(value, "")
			}
		}
	}

	// update length
	for index, length := range t.maxLength {
		if length < utf8.RuneCountInString(value[index]) {
			t.maxLength[index] = utf8.RuneCountInString(value[index])
		}
	}

	// add data
	t.infos = append(t.infos, value)
}

//DoPrint will return the string of table format.
func (t *Table) DoPrint() string {
	result := ""
	// update all max length : length += 2
	for index, lengthItem := range t.maxLength {
		t.maxLength[index] = lengthItem + 2
	}

	// print title
	result += CenterDivider
	for _, maxLength := range t.maxLength {
		for i := 0; i < maxLength; i++ {
			result += TitleDivider
		}
		result += "+"
	}
	result += "\n" + VerticalDivider
	for index, title := range t.titles {
		titleLength := utf8.RuneCountInString(title)
		spaceCount := t.maxLength[index] - titleLength
		if t.alignStyle[index] == RightAlignType {
			// do nothing
		} else if t.alignStyle[index] == LeftAlignType {
			spaceCount = 0
		} else {
			spaceCount = spaceCount / 2
		}

		for i := 0; i < spaceCount; i++ {
			result += " "
		}
		result += title
		for i := 0; i < (t.maxLength[index] - titleLength - spaceCount); i++ {
			result += " "
		}
		result += VerticalDivider
	}
	result += "\n" + CenterDivider
	for _, maxLength := range t.maxLength {
		for i := 0; i < maxLength; i++ {
			result += TitleDivider
		}
		result += "+"
	}

	// print value
	for _, valueItem := range t.infos {
		result += "\n" + VerticalDivider
		for index, value := range valueItem {
			valueLength := utf8.RuneCountInString(value)
			spaceCount := t.maxLength[index] - valueLength
			if t.alignStyle[index] == RightAlignType {
				// do nothing
			} else if t.alignStyle[index] == LeftAlignType {
				spaceCount = 0
			} else {
				spaceCount = spaceCount / 2
			}
			for i := 0; i < spaceCount; i++ {
				result += " "
			}
			result += value
			for i := 0; i < (t.maxLength[index] - valueLength - spaceCount); i++ {
				result += " "
			}
			result += VerticalDivider
		}
		result += "\n" + CenterDivider
		for _, maxLength := range t.maxLength {
			for i := 0; i < maxLength; i++ {
				result += HorizontalDivider
			}
			result += "+"
		}
	}
	return result
}

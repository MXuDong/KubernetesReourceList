package cmd

import (
	"unicode/utf8"
)

//====================================================== the print tools
type AlignStyleType int
type RowValue []string

const (
	HorizontalDivider = "-" // the horizontal divider
	TitleDivider      = "="
	VerticalDivider   = "|" // the vertical divider
	CenterDivider     = "+" // the center divider

	CenterAlignType = iota // the align style : for center
	RightAlignType         // the align style : for right
	LeftAlignType          // the align style : for left

)

type Table struct {
	// values
	titles RowValue
	infos  []RowValue

	// data
	maxLength  []int
	alignStyle []AlignStyleType
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

		for i := 0; i < spaceCount/2; i++ {
			result += " "
		}
		result += title
		for i := 0; i < (t.maxLength[index] - titleLength - spaceCount/2); i++ {
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
			for i := 0; i < spaceCount/2; i++ {
				result += " "
			}
			result += value
			for i := 0; i < (t.maxLength[index] - valueLength - spaceCount/2); i++ {
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

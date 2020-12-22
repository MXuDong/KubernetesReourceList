package cmd

import (
	"fmt"
	"testing"
)

func TestTable_DoPrint(t1 *testing.T) {
	t := Table{}
	t.InitTitle([]string{"test1", "tesasdf t 2"})
	t.addRows([]string{"Deployment", "apps/v1"})
	t.SetAlignStyle([]AlignStyleType{RightAlignType, LeftAlignType})
	result := t.DoPrint()
	fmt.Println(result)
}

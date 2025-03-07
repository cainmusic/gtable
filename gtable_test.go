package gtable

import (
	"strings"
	"testing"
)

func TestNormal(t *testing.T) {
	table := NewTable()
	table.AppendTitle("this is a title")
	table.AppendHead([]string{"h1", "h2", "h3"})
	table.AppendBody([]string{"1", "2", "3"})
	table.AppendBody([]string{"123456", "7890abc", "defghijklmnopqrst"})

	table.PrintData()
}

func TestMoreThanOneTitle(t *testing.T) {
	table := NewTable()
	table.AppendTitle("this is a title")
	table.AppendTitle("this is title2")
	table.AppendHead([]string{"id", "age", "height", "weight"})
	table.AppendBody([]string{"1", "19", "167", "140"})
	table.AppendBody([]string{"2", "22", "182", "202"})
	table.PrintData()
}

func TestMoreThanOneHead(t *testing.T) {
	table := NewTable()
	table.AppendTitle("this is a title")
	table.AppendHead([]string{"id", "name", "time"})
	table.AppendHead([]string{"int", "string", "time"})
	table.AppendBody([]string{"110", "abn", "2022-01-01 00:01:02"})
	table.AppendBody([]string{"111", "default", "2022-01-01 00:01:02"})
	table.PrintData()
}

func TestMoreThanOneTable(t *testing.T) {
	table := NewTable()
	table.AppendTitle("table no.1")
	table.AppendHead([]string{"h1", "h2", "h3"})
	table.AppendBody([]string{"1", "2", "3"})
	table.AppendBody([]string{"123456", "7890abc", "defghijklmnopqrst"})
	table.AppendTitle("table no.2")
	table.AppendHead([]string{"id", "age", "height", "weight"})
	table.AppendBody([]string{"1", "19", "167", "140"})
	table.AppendBody([]string{"2", "22", "182", "202"})
	table.PrintData()
}

func TestVeryLongTitle(t *testing.T) {
	table := NewTable()
	table.AppendTitle("table no.1")
	table.AppendHead([]string{"h1", "h2", "h3"})
	table.AppendBody([]string{"1", "2", "3"})
	table.AppendBody([]string{"123456", "7890abc", "defghijklmnopqrst"})
	table.AppendTitle("========================================================")
	table.AppendHead([]string{"id", "age", "height", "weight"})
	table.AppendBody([]string{"1", "19", "167", "140"})
	table.AppendBody([]string{"2", "22", "182", "202"})
	table.PrintData()
}

func TestEmptyTable(t *testing.T) {
	table := NewTable()
	table.PrintData()
}

func TestLongContent(t *testing.T) {
	table := NewTable()
	table.AppendTitle("Long Content Test")
	table.AppendHead([]string{"ID", "Description"})
	table.AppendBody([]string{"1", "This is a very long description text to test automatic wrapping or truncation functionality"})
	table.AppendBody([]string{"2", strings.Repeat("long-", 50) + "end"})
	table.PrintData()
}

func TestSpecialCharacters(t *testing.T) {
	table := NewTable()
	table.AppendTitle("Special Characters Test")
	table.AppendHead([]string{"Symbol", "Content"})
	table.AppendBody([]string{"Quotes", "Contains \"double quotes\" and 'single quotes'"})
	table.AppendBody([]string{"Newline", "Line1\nLine2"})
	table.PrintData()
}

func TestInvalidInput(t *testing.T) {
	table := NewTable()
	// 测试空标题
	table.AppendTitle("")
	// 测试空表头
	table.AppendHead([]string{"", ""})
	// 测试不一致的列数
	table.AppendBody([]string{"1", "2", "3"})
	table.AppendBody([]string{"4"})
	table.PrintData()
}

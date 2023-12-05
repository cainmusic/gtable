package gtable

import (
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

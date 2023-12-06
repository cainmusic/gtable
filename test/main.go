package main

import (
	"github.com/cainmusic/gtable"
)

func main() {
	normal()
	writeFile()
	readCsvFile()
	readJsonString()
	readJsonStringAndSetHead()
	readCsvFileAndSetHead()
	readDirTree()
	formatTree()
}

func normal() {
	table := gtable.NewTable()

	table.AppendTitle("this is a title")
	table.AppendHead([]string{"h1", "h2", "h3"})
	table.AppendBody([]string{"1", "2", "3"})
	table.AppendBody([]string{"123456", "7890abc", "defghijklmnopqrstuvwxyz"})

	table.PrintData()
}

func writeFile() {
	table := gtable.NewTable()

	table.AppendTitle("this is a title")
	table.AppendTitle("subtitle")
	table.AppendHead([]string{"h1", "h2", "h3"})
	table.AppendHead([]string{"h21", "h22", "h23"})
	table.AppendBody([]string{"123456", "7890abc", "defghijklmnopqrstuvwxyz"})
	table.AppendBody([]string{"1", "2", "3"})
	table.AppendBody([]string{"4", "5", "6"})
	table.AppendBody([]string{"7", "8", "9"})

	table.SetOutputFile("./out.txt")

	table.PrintData()
}

func readCsvFile() {
	table := gtable.NewTable()

	table.InitInputFromFile("./test.csv", gtable.DataTypeCsv)
	table.ReadFromInput()

	table.PrintData()
}

const s = `
	{"No.": 1, "Name": "Ed", "Text": "Knock knock."}
	{"No.": 2, "Name": "Sam", "Text": "Who's there?"}
	{"No.": 3, "Name": "Ed", "Text": "Go fmt."}
	{"No.": 4, "Name": "Sam", "Text": "Go fmt who?"}
	{"No.": 5, "Name": "Ed", "Text": "Go fmt yourself!"}
`

func readJsonString() {
	table := gtable.NewTable()

	table.InitInputFromString(s, gtable.DataTypeJson)
	table.ReadFromInput()

	table.PrintData()
}

func readJsonStringAndSetHead() {
	table := gtable.NewTable()

	table.InitInputFromString(s, gtable.DataTypeJson)
	table.SetTitle("some dialog")
	// SetHead会修改input，需要在Init方法后执行
	table.SetHead([]string{"No.", "Name", "Text"})
	table.ReadFromInput()

	table.PrintData()
}

func readCsvFileAndSetHead() {
	table := gtable.NewTable()

	table.InitInputFromFile("./test_no_head.csv", gtable.DataTypeCsv)
	table.SetHead([]string{"type", "count", "price", "limit"})
	table.ReadFromInput()

	table.PrintData()
}

func readDirTree() {
	table := gtable.NewTable()

	table.ReadDirTree("..")

	table.PrintData()
}

func formatTree() {
	table := gtable.NewTable()

	tls := []gtable.TreeLayer{
		gtable.TreeLayer{Layer: 0, Name: "/"},
		gtable.TreeLayer{Layer: 1, Name: "hi"},
		gtable.TreeLayer{Layer: 1, Name: "hello"},
		gtable.TreeLayer{Layer: 2, Name: "world"},
		gtable.TreeLayer{Layer: 2, Name: "china"},
		gtable.TreeLayer{Layer: 3, Name: "shanghai"},
		gtable.TreeLayer{Layer: 4, Name: "pudong"},
		gtable.TreeLayer{Layer: 3, Name: "beijing"},
		gtable.TreeLayer{Layer: 2, Name: "russia"},
		gtable.TreeLayer{Layer: 3, Name: "moscow"},
		gtable.TreeLayer{Layer: 1, Name: "see"},
		gtable.TreeLayer{Layer: 2, Name: "k"},
	}

	table.FormatTree(tls)

	table.SetNoBorder()

	table.PrintData()
}

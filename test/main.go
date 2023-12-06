package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cainmusic/gtable"
)

const outfile = "./out.txt"

type funcEntry struct {
	name string
	f    func()
	des  string
}

func main() {
	cases := []funcEntry{
		funcEntry{name: "normal", f: normal, des: "basic use"},
		funcEntry{name: "write file", f: writeFile, des: "write table into file"},
		funcEntry{name: "read csv file", f: readCsvFile, des: "read csv file to make table"},
		funcEntry{name: "read json string", f: readJsonString, des: "read json string to make table"},
		funcEntry{name: "read json string and set head", f: readJsonStringAndSetHead, des: "read json string and set head to make table"},
		funcEntry{name: "read csv file and set head", f: readCsvFileAndSetHead, des: "read csv file and set head to make table"},
		funcEntry{name: "read dir tree", f: readDirTree, des: "read dir to format tree"},
		funcEntry{name: "format tree", f: formatTree, des: "format tree"},
		funcEntry{name: "align", f: align, des: "align blocks"},
	}

	reader := bufio.NewReader(os.Stdin)
	table := gtable.NewTable()
	for i, en := range cases {
		table.AppendBody([]string{fmt.Sprint(i + 1), en.name, en.des})
	}

	for {
		table.PrintData()
		fmt.Println("type number to run func, or type quit to quit.")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		text = strings.TrimRight(text, "\r\n")
		if text == "quit" {
			return
		}
		idx, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("wrong input")
			continue
		}
		if idx < 1 || idx > len(cases) {
			fmt.Println("wrong input")
			continue
		}
		cases[idx-1].f()
	}
}

var normal = func() {
	table := gtable.NewTable()

	table.AppendTitle("this is a title")
	table.AppendHead([]string{"h1", "h2", "h3"})
	table.AppendBody([]string{"1", "2", "3"})
	table.AppendBody([]string{"123456", "7890abc", "defghijklmnopqrstuvwxyz"})

	table.PrintData()
}

var writeFile = func() {
	table := gtable.NewTable()

	table.AppendTitle("this is a title")
	table.AppendTitle("subtitle")
	table.AppendHead([]string{"h1", "h2", "h3"})
	table.AppendHead([]string{"h21", "h22", "h23"})
	table.AppendBody([]string{"123456", "7890abc", "defghijklmnopqrstuvwxyz"})
	table.AppendBody([]string{"1", "2", "3"})
	table.AppendBody([]string{"4", "5", "6"})
	table.AppendBody([]string{"7", "8", "9"})

	table.SetOutputFile(outfile)

	table.PrintData()

	fmt.Println("write data to file", outfile)
}

var readCsvFile = func() {
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

var readJsonString = func() {
	table := gtable.NewTable()

	table.InitInputFromString(s, gtable.DataTypeJson)
	table.ReadFromInput()

	table.PrintData()
}

var readJsonStringAndSetHead = func() {
	table := gtable.NewTable()

	table.InitInputFromString(s, gtable.DataTypeJson)
	table.SetTitle("some dialog")
	// SetHead会修改input，需要在Init方法后执行
	table.SetHead([]string{"No.", "Name", "Text"})
	table.ReadFromInput()

	table.PrintData()
}

var readCsvFileAndSetHead = func() {
	table := gtable.NewTable()

	table.InitInputFromFile("./test_no_head.csv", gtable.DataTypeCsv)
	table.SetHead([]string{"type", "count", "price", "limit"})
	table.ReadFromInput()

	table.PrintData()
}

var readDirTree = func() {
	table := gtable.NewTable()

	table.ReadDirTree("..")

	table.PrintData()
}

var formatTree = func() {
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

var align = func() {
	table := gtable.NewTable()

	table.AppendTitle("this is a title")
	table.AppendTitle("subtitle")
	table.AppendHead([]string{"h1", "h2", "h3"})
	table.AppendHead([]string{"h21", "h22", "h23"})
	table.AppendBody([]string{"123456", "7890abc", "defghijklmnopqrstuvwxyz"})
	table.AppendBody([]string{"1", "2", "3"})
	table.AppendBody([]string{"4", "5", "6"})
	table.AppendBody([]string{"7", "8", "9"})

	table.SetAlignTitle(gtable.AlignCenter)
	table.SetAlignHeadAll(gtable.AlignLeft)
	table.SetAlignBodyAll(gtable.AlignRight)

	table.PrintData()
}

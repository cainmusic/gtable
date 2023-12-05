package gtable

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var DefaultOutput = os.Stdout

func (t *table) SetOutput(w io.Writer) {
	t.output = w
}

// 默认append模式
func (t *table) SetOutputFile(pathToFile string) {
	f, err := os.OpenFile(pathToFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// TODO handle error
		return
	}
	t.SetOutput(f)
}

func (t *table) Print(a ...any) {
	fmt.Fprint(t.output, a...)
}

func (t *table) Println(a ...any) {
	fmt.Fprintln(t.output, a...)
}

func (t *table) PrintData() {
	if len(t.data) <= 0 {
		return
	}
	lastBType := btNull
	for _, line := range t.data {
		// 行类型变化时，才打印border
		if lastBType != line.bType {
			// 首行，且首行是title，打印不带turn的border
			t.PrintBorder(lastBType == btNull && line.bType == btTitle)
		}
		lastBType = line.bType
		t.PrintLine(line)
	}
	t.PrintBorder(false)
}

func (t *table) PrintBorder(titleTop bool) {
	buffer := new(bytes.Buffer)
	buffer.Write(turnByte)
	if titleTop {
		buffer.Write(bytes.Repeat(borderByte, max(t.status.titleWidth, t.status.lineWidth)))
	} else {
		lenWidth := len(t.status.width)
		for i, w := range t.status.width {
			buffer.Write(bytes.Repeat(borderByte, w))
			if i < lenWidth-1 {
				buffer.Write(turnByte)
			}
		}
		if t.status.titleWidth > t.status.lineWidth {
			buffer.Write(bytes.Repeat(borderByte, t.status.titleWidth-t.status.lineWidth))
		}
	}
	buffer.Write(turnByte)
	t.Println(buffer.String())
}

func (t *table) PrintLine(line line) {
	t.Print(string(splitByte))
	switch line.bType {
	case btTitle:
		t.PrintTitle(line.blocks[0])
	case btHead:
		t.PrintBlocks(line.blocks, btHead)
	case btBody:
		t.PrintBlocks(line.blocks, btBody)
	}
	t.Println(string(splitByte))
}

func (t *table) PrintTitle(title block) {
	t.Print(title.raw)
	tw := max(t.status.titleWidth, t.status.lineWidth)
	if len(title.raw) < tw {
		t.Print(string(bytes.Repeat(spaceByte, tw-len(title.raw))))
	}

}

func (t *table) PrintBlocks(blocks []block, bt blockType) {
	for i := 0; i < len(t.status.width); i++ {
		s := ""
		if i < len(blocks) {
			s = blocks[i].raw
		}
		t.Print(s)
		wid := t.status.width[i]
		if wid > len(s) {
			t.Print(string(bytes.Repeat(spaceByte, wid-len(s))))
		}
		if i < len(t.status.width)-1 {
			t.Print(string(splitByte))
		}
	}
	if t.status.titleWidth > t.status.lineWidth {
		t.Print(string(bytes.Repeat(spaceByte, t.status.titleWidth-t.status.lineWidth)))
	}
}

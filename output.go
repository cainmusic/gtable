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
	if t.config.noBorder {
		return
	}
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
	if !t.config.noBorder {
		t.Print(string(splitByte))
	}
	switch line.bType {
	case btTitle:
		t.PrintTitle(line.blocks[0])
	case btHead:
		t.PrintBlocks(line.blocks, btHead)
	case btBody:
		t.PrintBlocks(line.blocks, btBody)
	}
	if !t.config.noBorder {
		t.Print(string(splitByte))
	}
	t.Println()
}

func (t *table) PrintTitle(title block) {
	t.printBlock(title.raw, max(t.status.titleWidth, t.status.lineWidth), btTitle, 0)
}

func (t *table) PrintBlocks(blocks []block, bt blockType) {
	for i := 0; i < len(t.status.width); i++ {
		s := ""
		if i < len(blocks) {
			s = blocks[i].raw
		}
		t.printBlock(s, t.status.width[i], bt, i)
		if i < len(t.status.width)-1 {
			t.Print(string(splitByte))
		}
	}
	if t.status.titleWidth > t.status.lineWidth {
		t.Print(string(bytes.Repeat(spaceByte, t.status.titleWidth-t.status.lineWidth)))
	}
}

func (t *table) printBlock(s string, length int, bt blockType, idx int) {
	at := t.getAlignType(bt, idx)
	leftLen, rightLen := 0, 0
	switch at {
	case AlignLeft:
		if length > len(s) {
			rightLen = length - len(s)
		}
	case AlignRight:
		if length > len(s) {
			leftLen = length - len(s)
		}
	case AlignCenter:
		if length > len(s) {
			leftLen = (length - len(s)) / 2
			rightLen = length - len(s) - leftLen
		}
	}
	t.Print(string(bytes.Repeat(spaceByte, leftLen)))
	t.Print(s)
	t.Print(string(bytes.Repeat(spaceByte, rightLen)))
}

func (t *table) getAlignType(bt blockType, idx int) alignType {
	align := t.config.align
	if align[bt] != nil && len(align[bt]) > idx {
		return align[bt][idx]
	}
	return AlignLeft
}

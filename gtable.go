package gtable

import (
	"io"
)

var (
	spaceByte  = []byte{' '}
	splitByte  = []byte{'|'}
	turnByte   = []byte{'+'}
	borderByte = []byte{'-'}
)

type blockType uint8

const (
	btNull blockType = iota
	btTitle
	btHead
	btBody
)

type block struct {
	raw string
}

type line struct {
	bType  blockType
	blocks []block
}

type status struct {
	width      []int
	totalWidth int
	blockCount int
	titleWidth int
	lineWidth  int
}

type config struct {
	noBorder bool
}

type table struct {
	data   []line
	status status
	config config

	input *inputData

	output io.Writer
}

func NewTable() *table {
	return &table{
		data:   []line{},
		output: DefaultOutput,
	}
}

func (t *table) SetNoBorder() {
	t.config.noBorder = true
}

func (t *table) AppendTitle(s string) {
	t.AppendBlocks([]string{s}, btTitle)
}

func (t *table) AppendHead(ss []string) {
	t.AppendBlocks(ss, btHead)
}

func (t *table) AppendBody(ss []string) {
	t.AppendBlocks(ss, btBody)
}

func (t *table) AppendBlocks(ss []string, bt blockType) {
	blocks := make([]block, len(ss))
	width := make([]int, len(ss))
	for i, s := range ss {
		blocks[i] = block{raw: s}
		width[i] = len(s)
	}
	t.data = append(t.data, line{
		bType:  bt,
		blocks: blocks,
	})
	if bt == btTitle {
		t.updateStatusTitleWidth(width[0])
	} else {
		t.updateStatusBlockWidth(width)
	}
}

func (t *table) ClearData() {
	t.data = []line{}
}

func (t *table) updateStatusTitleWidth(titleWidth int) {
	if titleWidth > t.status.titleWidth {
		t.status.titleWidth = titleWidth
	}
}

func (t *table) updateStatusBlockWidth(width []int) {
	blockCount := len(width)
	if blockCount > t.status.blockCount {
		t.status.blockCount = blockCount
	}

	if blockCount > len(t.status.width) {
		for i, w := range t.status.width {
			if w > width[i] {
				width[i] = w
			}
		}
		t.status.width = width
	} else {
		for i, w := range width {
			if w > t.status.width[i] {
				t.status.width[i] = w
			}
		}
	}

	totalWidth := 0
	for _, w := range t.status.width {
		totalWidth += w
	}
	t.status.totalWidth = totalWidth

	t.status.lineWidth = t.status.blockCount + t.status.totalWidth - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

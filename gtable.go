package gtable

import (
	"io"
)

/*
基本名词解释：
block，块，相当于表格里的单个字段
title，标题block
head，表头block
body，数据block
block一般包含title、head、body
由于head和body结构更接近，有时block特指head和body，title单独处理

table是总表格结构，包含
data，数据
status，状态，不由外部设置，由data决定
config，配置，外部设置
input，输入模块
output，输出模块
*/

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

type alignType uint8

const (
	AlignLeft alignType = iota
	AlignRight
	AlignCenter
)

type config struct {
	noBorder bool
	align    map[blockType][]alignType
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
		data: []line{},
		config: config{
			align: map[blockType][]alignType{},
		},
		output: DefaultOutput,
	}
}

// 边框

// 去除边框
func (t *table) SetNoBorder() {
	t.config.noBorder = true
}

// 对齐

// title对齐
func (t *table) SetAlignTitle(at alignType) {
	t.config.align[btTitle] = []alignType{at}
}

// head对齐
func (t *table) SetAlignHeadAll(at alignType) {
	ats := make([]alignType, t.status.blockCount)
	for i := 0; i < t.status.blockCount; i++ {
		ats[i] = at
	}
	t.config.align[btHead] = ats
}

func (t *table) SetAlignHeadCol(ats []alignType) {
	t.config.align[btHead] = ats
}

// body对齐
func (t *table) SetAlignBodyAll(at alignType) {
	ats := make([]alignType, t.status.blockCount)
	for i := 0; i < t.status.blockCount; i++ {
		ats[i] = at
	}
	t.config.align[btBody] = ats
}

func (t *table) SetAlignBodyCol(ats []alignType) {
	t.config.align[btBody] = ats
}

// 写入data

// 写入title
func (t *table) AppendTitle(s string) {
	t.AppendBlocks([]string{s}, btTitle)
}

// 写入head
func (t *table) AppendHead(ss []string) {
	t.AppendBlocks(ss, btHead)
}

// 写入body
func (t *table) AppendBody(ss []string) {
	t.AppendBlocks(ss, btBody)
}

// 写入block
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

// 清空data
func (t *table) ClearData() {
	t.data = []line{}
}

// 更新status

// 更新title宽度
func (t *table) updateStatusTitleWidth(titleWidth int) {
	if titleWidth > t.status.titleWidth {
		t.status.titleWidth = titleWidth
	}
}

// 更新block宽度
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

// utils

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

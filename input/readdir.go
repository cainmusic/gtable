package input

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

var (
	emptySpace  = []byte{' ', ' ', ' ', ' '}
	passbySpace = []byte{'|', ' ', ' ', ' '}
	treeBorder  = []byte{'+', '-', '-', '-'}
)

type entry struct {
	layer int
	name  string
	res   string
}

type DirInfo struct {
	dir     string
	layer   int
	flags   []bool
	entries []entry
	skipDot bool
}

func NewDir(pathToDir string) *DirInfo {
	return &DirInfo{
		dir:     pathToDir,
		flags:   []bool{false},
		entries: []entry{},
		skipDot: true,
	}
}

func (dir *DirInfo) Read() {
	dir.formatEntry(dir.dir)
	dir.readdir(dir.dir)
}

func (dir *DirInfo) ReadAndPrint() {
	dir.Read()
	dir.printRes()
}

func (dir *DirInfo) ReadAndGet() []string {
	dir.Read()
	return dir.getRes()
}

func (dir *DirInfo) readdir(pathToDir string) {
	entries, _ := os.ReadDir(pathToDir)
	if dir.skipDot {
		newEntries := []os.DirEntry{}
		for _, en := range entries {
			if !strings.HasPrefix(en.Name(), ".") {
				newEntries = append(newEntries, en)
			}
		}
		entries = newEntries
	}
	dir.layer++
	if dir.layer > len(dir.flags) {
		dir.flags = append(dir.flags, false)
	}
	for i, entry := range entries {
		dir.flags[dir.layer-1] = i >= len(entries)-1
		dir.formatEntry(entry.Name())
		if entry.IsDir() {
			dir.readdir(pathToDir + "/" + entry.Name())
		}
	}
	dir.layer--
}

func (dir *DirInfo) formatEntry(path string) {
	res := ""
	if dir.layer > 1 {
		for i := 0; i < dir.layer-1; i++ {
			if (dir.flags)[i] {
				res += fmt.Sprint(string(emptySpace))
			} else {
				res += fmt.Sprint(string(passbySpace))
			}
		}
	}
	if dir.layer > 0 {
		res += fmt.Sprint(string(bytes.Repeat(treeBorder, 1)))
	}
	res += fmt.Sprint(path)
	dir.entries = append(dir.entries, entry{
		layer: dir.layer,
		name:  path,
		res:   res,
	})
}

func (dir *DirInfo) printRes() {
	for _, en := range dir.entries {
		fmt.Println(en.res)
	}
}

func (dir *DirInfo) getRes() []string {
	ss := make([]string, 0, len(dir.entries))
	for _, en := range dir.entries {
		ss = append(ss, en.res)
	}
	return ss
}

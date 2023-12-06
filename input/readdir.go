package input

import (
	//"bytes"
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

func (dir *DirInfo) ReadDot() {
	dir.skipDot = false
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
		res += fmt.Sprint(string(treeBorder))
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

type OnlyEntries struct {
	layer   int
	flags   []bool
	entries []entry
}

func NewEntries() *OnlyEntries {
	return &OnlyEntries{
		entries: []entry{},
	}
}

func (ens *OnlyEntries) Append(layer int, name string) {
	if layer > len(ens.flags) {
		ens.flags = append(ens.flags, false)
	}
	res := ""
	ens.entries = append(ens.entries, entry{
		layer: layer,
		name:  name,
		res:   res,
	})
	if layer > ens.layer {
		ens.layer = layer
	}
}

func (ens *OnlyEntries) Format() {
	ens.flags = make([]bool, ens.layer)
	passbyEnd := make([]int, ens.layer)

	lastLayer := 0
	for i, en := range ens.entries {
		if en.layer > 0 {
			if lastLayer < en.layer {
				for j := i + 1; j < len(ens.entries); j++ {
					if ens.entries[j].layer < en.layer {
						break
					}
					if ens.entries[j].layer == en.layer {
						passbyEnd[en.layer-1] = j
					}
				}
			} else if lastLayer > en.layer {
				passbyEnd[en.layer-1] = 0
			}
			if i <= passbyEnd[en.layer-1] {
				ens.flags[en.layer-1] = true
			} else {
				ens.flags[en.layer-1] = false
			}
			lastLayer = en.layer

			for k := 0; k < en.layer-1; k++ {
				if (ens.flags)[k] {
					en.res += fmt.Sprint(string(passbySpace))
				} else {
					en.res += fmt.Sprint(string(emptySpace))
				}
			}
		}
		if en.layer > 0 {
			en.res += fmt.Sprint(string(treeBorder))
		}
		en.res += fmt.Sprint(en.name)
		ens.entries[i] = en
	}
}

func (ens *OnlyEntries) FormatAndPrint() {
	ens.Format()
	ens.printRes()
}

func (ens *OnlyEntries) FormatAndGet() []string {
	ens.Format()
	return ens.getRes()
}

func (ens *OnlyEntries) printRes() {
	for _, en := range ens.entries {
		fmt.Println(en.res)
	}
}

func (ens *OnlyEntries) getRes() []string {
	ss := make([]string, 0, len(ens.entries))
	for _, en := range ens.entries {
		ss = append(ss, en.res)
	}
	return ss
}

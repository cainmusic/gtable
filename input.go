package gtable

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/cainmusic/gtable/input"
)

type DataType uint8

const (
	DataTypeNull DataType = iota
	DataTypeCsv
	DataTypeJson
)

type inputData struct {
	reader   io.Reader
	dataType DataType
	fHead    bool
	ssHead   []string
}

func (t *table) InitInputFromString(str string, dt DataType) {
	t.InitInputFromReader(strings.NewReader(str), dt)
}

func (t *table) InitInputFromFile(pathToFile string, dt DataType) {
	f, err := os.Open(pathToFile)
	if err != nil {
		// TODO handle error
		return
	}
	t.InitInputFromReader(f, dt)
}

func (t *table) InitInputFromReader(reader io.Reader, dt DataType) {
	t.input = &inputData{}
	t.SetInputReader(reader)
	t.SetDataType(dt)
}

func (t *table) SetInputReader(reader io.Reader) {
	t.input.reader = reader
}

func (t *table) SetDataType(dt DataType) {
	t.input.dataType = dt
}

func (t *table) SetTitle(s string) {
	t.AppendTitle(s)
}

func (t *table) SetHead(ss []string) {
	t.AppendHead(ss)
	t.input.fHead = true
	t.input.ssHead = ss
}

func (t *table) ReadFromInput() {
	switch t.input.dataType {
	case DataTypeNull:
		// TODO handle error
		return
	case DataTypeCsv:
		t.ReadCsvReader()
	case DataTypeJson:
		t.ReadJsonReader()
	}
}

func (t *table) ReadCsvReader() {
	r := csv.NewReader(t.input.reader)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			// TODO handle error
			return
		}
		if !t.input.fHead {
			t.AppendHead(record)
			t.input.fHead = true
		}
		t.AppendBody(record)
	}
}

func (t *table) ReadJsonReader() {
	d := json.NewDecoder(t.input.reader)
	var keys []string
	if t.input.fHead {
		keys = t.input.ssHead
	}
	for {
		var m any
		if err := d.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err.Error())
			// TODO handle error
			return
		}
		mmap, _ := m.(map[string]any)
		if !t.input.fHead {
			keys = make([]string, 0, len(mmap))
			for k, _ := range mmap {
				keys = append(keys, k)
			}
			t.AppendHead(keys)
			t.input.fHead = true
		}
		values := make([]string, len(keys))
		for i := 0; i < len(keys); i++ {
			values[i] = fmt.Sprint(mmap[keys[i]])
		}
		t.AppendBody(values)
	}
}

func (t *table) ClearDataAndReadFromInput() {
	t.ClearData()
	t.ReadFromInput()
}

func (t *table) ReadDirTree(pathToDir string) {
	dir := input.NewDir(pathToDir)
	//dir.ReadDot()
	for _, str := range dir.ReadAndGet() {
		t.AppendBody([]string{str})
	}
}

type TreeLayer struct {
	Layer int
	Name  string
}

func (t *table) FormatTree(tls []TreeLayer) {
	ens := input.NewEntries()
	for i := 0; i < len(tls); i++ {
		ens.Append(tls[i].Layer, tls[i].Name)
	}
	for _, str := range ens.FormatAndGet() {
		t.AppendBody([]string{str})
	}
}

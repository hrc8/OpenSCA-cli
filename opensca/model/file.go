package model

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/xmirrorsecurity/opensca-cli/opensca/logs"
)

type File struct {
	Abspath string
	Relpath string
}

func (file *File) Path() string {
	if file != nil {
		return file.Relpath
	}
	return ""
}

func (file File) OpenReader(do func(reader io.Reader)) {
	f, err := os.Open(file.Abspath)
	if err != nil {
		logs.Warnf("open file %s fail: %s", file.Relpath, err)
		return
	}
	defer f.Close()
	do(f)
}

func (file File) ReadLine(do func(line string)) {
	file.OpenReader(func(reader io.Reader) {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			do(strings.TrimRight(scanner.Text(), "\n\r"))
		}
	})
}

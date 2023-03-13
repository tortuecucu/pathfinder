package clipboard

import (
	"strings"

	"golang.design/x/clipboard"
)

type Clipboard struct{}

func (t Clipboard) RecordLines(lines *[]string) {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	clipboard.Write(clipboard.FmtText, []byte(strings.Join((*lines), "\n")))
}

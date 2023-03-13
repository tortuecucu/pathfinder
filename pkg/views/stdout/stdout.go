package stdout

import (
	"fmt"

	"github.com/tortuecucu/pathfinder/pkg/core"
	"github.com/tortuecucu/pathfinder/pkg/formatters/text"
)

type Stdout struct{}

func (s Stdout) DisplayLines(lines *[]string) {
	for _, line := range *lines {
		fmt.Println(line)
	}
}
func (s Stdout) Display(exe *core.FactCollection) {
	formatter := text.TextFormatter{}
	lines := formatter.Format(exe)
	s.DisplayLines(&lines)
}

package stdout

import (
	"fmt"

	"github.com/tortuecucu/pathfinder/pkg/formatters/text"
	"github.com/tortuecucu/pathfinder/pkg/plan"
)

type Stdout struct{}

func (s Stdout) DisplayLines(lines *[]string) {
	for _, line := range *lines {
		fmt.Println(line)
	}
}
func (s Stdout) Display(exe *plan.PlanExecution) {
	formatter := text.TextFormatter{}
	fmt.Println(formatter.Format(exe))
}

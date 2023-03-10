package text

import "github.com/tortuecucu/pathfinder/pkg/plan"

type TextFormatter struct {
}

func (f TextFormatter) Format(exe *plan.PlanExecution) []string {
	var r []string
	r = append(r, "text line") //TODO: code it
	return r
}

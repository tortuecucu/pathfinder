package views

import "github.com/tortuecucu/pathfinder/pkg/plan"

type View interface {
	DisplayLines(lines *[]string)
	Display(exe *plan.PlanExecution)
}

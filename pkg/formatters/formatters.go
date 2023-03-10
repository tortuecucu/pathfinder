package formatters

import "github.com/tortuecucu/pathfinder/pkg/plan"

type Formatter interface {
	Format(exe *plan.PlanExecution) []string
}

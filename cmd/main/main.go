package main

import (
	"github.com/tortuecucu/pathfinder/pkg/collectors/common"
	"github.com/tortuecucu/pathfinder/pkg/plan"
	"github.com/tortuecucu/pathfinder/pkg/views/stdout"
)

func main() {

	plan := plan.NewPlan("testplan")
	common.AddCoreCollectors(plan)

	exe := plan.Execute()
	stdout.Stdout{}.Display(exe)
}

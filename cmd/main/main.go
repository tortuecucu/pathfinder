package main

import (
	"github.com/tortuecucu/pathfinder/pkg/commands"
	"github.com/tortuecucu/pathfinder/pkg/plan"
	"github.com/tortuecucu/pathfinder/pkg/views/stdout"
)

func main() {
	plan := plan.NewPlan("testplan")
	plan.Actions["actiontest"] = commands.NewCommand("hostname")

	exe := plan.Execute()
	stdout.Stdout{}.Display(exe)
}

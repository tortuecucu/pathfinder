package commands

import (
	"strings"

	"github.com/tortuecucu/pathfinder/pkg/plan"
)

type Command struct {
	commandName       string
	commandParameters []string
}

func (t Command) Run(execution *plan.PlanExecution) {
	execution.Results[t.commandName] = plan.NewResult(t, "tbd", t.commandName)
}
func (t Command) Name() string {
	return "command '" + t.commandName + "' args:'" + strings.Join(t.commandParameters[:], ",") + "'"
}

func NewCommand(name string, parameters ...string) Command {
	return Command{commandName: name, commandParameters: parameters}
}

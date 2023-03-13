package commands

import (
	"strings"

	"github.com/tortuecucu/pathfinder/pkg/core"
)

type Command struct {
	core.Runnable
	commandName       string
	commandParameters []string
}

func (t Command) Run(facts *core.FactCollection) {
	facts.AddFact(t.commandName, "tbd", t)
}
func (t Command) Name() string {
	return "command '" + t.commandName + "' args:'" + strings.Join(t.commandParameters[:], ",") + "'"
}

func NewCommand(name string, parameters ...string) Command {
	return Command{commandName: name, commandParameters: parameters}
}

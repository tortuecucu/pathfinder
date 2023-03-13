package views

import "github.com/tortuecucu/pathfinder/pkg/core"

type View interface {
	DisplayLines(lines *[]string)
	Display(exe *core.FactCollection)
}

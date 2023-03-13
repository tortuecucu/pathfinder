package formatters

import "github.com/tortuecucu/pathfinder/pkg/core"

type Formatter interface {
	Format(exe *core.FactCollection) []string
}

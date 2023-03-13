package plan

import (
	"time"

	"github.com/google/uuid"
	"github.com/tortuecucu/pathfinder/pkg/core"
)

type CollectPlan struct {
	Name       string
	collectors map[string]core.Runnable
}

func (p CollectPlan) Collectors() map[string]core.Runnable {
	return p.collectors
}

func (p CollectPlan) Execute() *core.FactCollection {
	return NewCollection(p)
}

func (p CollectPlan) AddCollector(c core.Runnable) {
	p.collectors[c.Name()] = c
}

func NewPlan(name string) CollectPlan {
	p := CollectPlan{Name: name, collectors: make(map[string]core.Runnable)}

	return p
}

func NewCollection(plan core.Plan) *core.FactCollection {
	exe := &core.FactCollection{Plan: plan, Start: time.Now(), Uuid: uuid.New()}
	exe.Facts = make(map[string]*core.Fact)

	for _, test := range plan.Collectors() {
		test.Run(exe)
	}

	exe.End = time.Now()
	return exe
}

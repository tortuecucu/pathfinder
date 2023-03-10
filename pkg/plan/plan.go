package plan

import (
	"time"

	"github.com/google/uuid"
)

type Result struct {
	Action    Runnable
	Category  string
	Name      string
	Timestamp time.Time
	Value     interface{}
}

func NewResult(test Runnable, value interface{}, name string) *Result {
	return &Result{Name: name, Action: test, Value: value, Timestamp: time.Now()}
}

type PlanExecution struct {
	ActionPlan ActionPlan
	Start      time.Time
	End        time.Time
	Uuid       uuid.UUID
	Results    map[string]*Result
}

func NewExecution(plan ActionPlan) *PlanExecution {
	exe := &PlanExecution{ActionPlan: plan, Start: time.Now(), Uuid: uuid.New()}
	exe.Results = make(map[string]*Result)

	for _, test := range plan.Actions {
		test.Run(exe)
	}

	exe.End = time.Now()
	return exe
}

type Runnable interface {
	Run(exe *PlanExecution)
	Name() string
}

type ActionPlan struct {
	Name    string
	Actions map[string]Runnable
}

func (p ActionPlan) Execute() *PlanExecution {
	return NewExecution(p)
}

func NewPlan(name string) ActionPlan {
	p := ActionPlan{Name: name}
	p.Actions = make(map[string]Runnable)
	return p
}

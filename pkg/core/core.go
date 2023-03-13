package core

import (
	"time"

	"github.com/google/uuid"
)

type Runnable interface {
	Run(facts *FactCollection)
	Name() string
}

type Fact struct {
	Collector Runnable
	Category  string
	Name      string
	Timestamp time.Time
	Value     interface{}
}

type Plan interface {
	Execute() *FactCollection
	AddCollector(r Runnable)
	Collectors() map[string]Runnable
}

type FactCollection struct {
	Plan     Plan
	Campaign string
	Start    time.Time
	End      time.Time
	Uuid     uuid.UUID
	Facts    map[string]*Fact
}

func (f FactCollection) AddFact(name string, value interface{}, collector Runnable) error {
	f.Facts[name] = NewFact(collector, value, name)
	return nil //TODO: code it
}

func NewFact(collector Runnable, value interface{}, name string) *Fact {
	return &Fact{Name: name, Collector: collector, Value: value, Timestamp: time.Now()}
}

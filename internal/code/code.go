package code

import "github.com/duosonic62/codanalyzer-domains/internal/scale"

type Code interface {
	Name() string
	Root() *scale.Note
	Notes() []scale.Note
	Contains(other Code) bool
}

type Factory struct {
	intervals []scale.Interval
}

func NewCodeFactory(intervals []scale.Interval) *Factory {
	return &Factory{
		intervals: intervals,
	}
}

func (f Factory) CreateCodes() (error, []Code) {
	return nil, nil
}
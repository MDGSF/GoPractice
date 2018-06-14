package StrategyPattern

type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

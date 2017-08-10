package abcgo

import (
	"math"

	"github.com/atgjack/prob"
)

// Proposer generates parameter values randomly.
type Proposer interface {
	Moments() []float64
	Propose() float64
	Prob(...float64) []float64
	LogProb(...float64) []float64
}

type NormalProposer struct {
	prob.Normal
}

func (n *NormalProposer) Moments(names ...string) []float64 {
	var moments []float64
	for _, name := range names {
		switch name {
		case "Mean":
			moments = append(moments, n.Mean())
		case "Variance":
			moments = append(moments, n.Variance())
		case "Skewness":
			moments = append(moments, n.Skewness())
		case "Kurtosis":
			moments = append(moments, n.Kurtosis())
		case "StdDev":
			moments = append(moments, n.StdDev())
		case "RelStdDev":
			moments = append(moments, n.RelStdDev())
		}
	}
	return moments
}

func (n *NormalProposer) Propose() float64 {
	return n.Random()
}

func (n *NormalProposer) Prob(values ...float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		res[i] = n.Pdf(v)
	}
	return res
}

func (n *NormalProposer) LogProb(values ...float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		res[i] = math.Log(n.Pdf(v))
	}
	return res
}

type Proposers []Proposer

func (p Proposers) Propose() []float64 {
	params := make([]float64, len(p))
	for i, proposer := range p {
		params[i] = proposer.Propose()
	}
	return params
}

package abcgo

import (
	"math"

	"github.com/atgjack/prob"
)

// Proposer generates parameter values randomly.
type Proposer interface {
	Moments(...string) []float64
	Propose() float64
	Prob(float64) float64
	LogProb(float64) float64
	Probs(...float64) []float64
	LogProbs(...float64) []float64
	UpdateMoments(...float64)
}

type NormalProposer struct {
	prob.Normal
}

func NewNormalProposer(mu, sigma float64) *NormalProposer {
	return &NormalProposer{prob.Normal{mu, sigma}}
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

func (n *NormalProposer) Prob(value float64) float64 {
	return n.Pdf(value)
}

func (n *NormalProposer) Probs(values ...float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		res[i] = n.Pdf(v)
	}
	return res
}

func (n *NormalProposer) LogProb(value float64) float64 {
	return math.Log(n.Pdf(value))
}

func (n *NormalProposer) LogProbs(values ...float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		res[i] = math.Log(n.Pdf(v))
	}
	return res
}

func (n *NormalProposer) UpdateMoments(moments ...float64) {
	if len(moments) != 2 {
		panic("Requires 2 values to update mu and sigma respectively")
	}
	n.Mu = moments[0]
	n.Sigma = moments[1]
}

type Proposers []Proposer

func (p Proposers) Propose() []float64 {
	params := make([]float64, len(p))
	for i, proposer := range p {
		params[i] = proposer.Propose()
	}
	return params
}

func (p Proposers) Probs(values ...float64) []float64 {
	probs := make([]float64, len(values))
	for i, v := range values {
		probs[i] = p[i].Prob(v)
	}
	return probs
}

func (p Proposers) LogProbs(values ...float64) []float64 {
	probs := make([]float64, len(values))
	for i, v := range values {
		probs[i] = p[i].LogProb(v)
	}
	return probs
}

func (p Proposers) TotalProb(values ...float64) float64 {
	prob := float64(1)
	for i, v := range values {
		prob *= p[i].Prob(v)
	}
	return prob
}

func (p Proposers) TotalLogProb(values ...float64) float64 {
	prob := float64(0)
	for i, v := range values {
		prob += p[i].LogProb(v)
	}
	return prob
}

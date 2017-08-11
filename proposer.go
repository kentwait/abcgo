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

// NormalProposer is a Proposer based on a normal distribution with a mean
// mu and standard deviation sigma.
type NormalProposer struct {
	prob.Normal
}

// NewNormalProposer is a constructor that returns a pointer to
// a new NormalProposer.
func NewNormalProposer(mu, sigma float64) *NormalProposer {
	n, _ := prob.NewNormal(mu, sigma)
	return &NormalProposer{n}
}

// Moments returns the statistics of the shape of the normal distribution
// proposer.
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

// Propose returns a random value from the initialized normal distribution.
func (n *NormalProposer) Propose() float64 {
	return n.Random()
}

// Prob returns the probability of a value in the initialized normal
// distribution.
func (n *NormalProposer) Prob(value float64) float64 {
	return n.Pdf(value)
}

// Probs returns probabilities of one or more given values based on the
// initialized normal distribution.
func (n *NormalProposer) Probs(values ...float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		res[i] = n.Pdf(v)
	}
	return res
}

// LogProb returns the log probability of a value in the initialized normal
// distribution.
func (n *NormalProposer) LogProb(value float64) float64 {
	return math.Log(n.Pdf(value))
}

// LogProbs returns the log probabilities of one or more given values based on
// the initialized normal distribution.
func (n *NormalProposer) LogProbs(values ...float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		res[i] = math.Log(n.Pdf(v))
	}
	return res
}

// UpdateMoments updates the mean, and standard deviation of
// the probability density of the normal distribution depending on the number
// of values given. If only one value is given,
func (n *NormalProposer) UpdateMoments(moments ...float64) {
	n.Mu = moments[0]
	if len(moments) == 2 {
		n.Sigma = moments[1]
	}
}

// Proposers is a list of Proposer types to accomodate proposing new
// values to multiple variables.
type Proposers []Proposer

// Propose gets new random values from each Proposer using their
// Propose() method.
func (p Proposers) Propose() []float64 {
	params := make([]float64, len(p))
	for i, proposer := range p {
		params[i] = proposer.Propose()
	}
	return params
}

// Probs returns the probability of each value based on the probability
// density of its corresponding Proposer by position.
func (p Proposers) Probs(values ...float64) []float64 {
	probs := make([]float64, len(values))
	for i, v := range values {
		probs[i] = p[i].Prob(v)
	}
	return probs
}

// LogProbs returns the log probability of each value based on the probability
// density of its corresponding Proposer by position.
func (p Proposers) LogProbs(values ...float64) []float64 {
	probs := make([]float64, len(values))
	for i, v := range values {
		probs[i] = p[i].LogProb(v)
	}
	return probs
}

// TotalProb returns the total probability of the given values measured based
// on it corresponding Proposer.
func (p Proposers) TotalProb(values ...float64) float64 {
	prob := float64(1)
	for i, v := range values {
		prob *= p[i].Prob(v)
	}
	return prob
}

// TotalLogProb returns the total log probability of the given values measured
// based on it corresponding Proposer.
func (p Proposers) TotalLogProb(values ...float64) float64 {
	prob := float64(0)
	for i, v := range values {
		prob += p[i].LogProb(v)
	}
	return prob
}

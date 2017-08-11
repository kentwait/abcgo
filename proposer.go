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
	UpdateFields(...float64)
}

// NormalProposer is a Proposer based on a normal distribution with a mean
// mu and standard deviation sigma.
type NormalProposer struct {
	prob.Normal
}

// NewNormalProposer is a constructor that returns a pointer to
// a new NormalProposer.
func NewNormalProposer(mu, sigma float64) *NormalProposer {
	p, _ := prob.NewNormal(mu, sigma)
	return &NormalProposer{p}
}

// Moments returns the statistics of the shape of the normal probability
// density.
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

// Propose returns a random value from the initialized normal probability
// density.
func (n *NormalProposer) Propose() float64 {
	return n.Random()
}

// Prob returns the probability of a value in the initialized normal
// probability density.
func (n *NormalProposer) Prob(value float64) float64 {
	return n.Pdf(value)
}

// Probs returns probabilities of one or more given values based on the
// initialized normal probability density.
func (n *NormalProposer) Probs(values ...float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		res[i] = n.Pdf(v)
	}
	return res
}

// LogProb returns the log probability of a value in the initialized normal
// probability density.
func (n *NormalProposer) LogProb(value float64) float64 {
	return math.Log(n.Pdf(value))
}

// LogProbs returns the log probabilities of one or more given values based on
// the initialized normal probability density.
func (n *NormalProposer) LogProbs(values ...float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		res[i] = math.Log(n.Pdf(v))
	}
	return res
}

// UpdateFields updates the mean and standard deviation of
// the probability density of the normal distribution depending on the number
// of values given.
func (n *NormalProposer) UpdateFields(values ...float64) {
	n.Mu = values[0]
	if len(values) == 2 {
		n.Sigma = values[1]
	}
}

// GammaProposer is a Proposer based on a gamma distribution with shape and
// rate parameters.
type GammaProposer struct {
	prob.Gamma
}

// NewGammaProposer is a constructor that returns a pointer to
// a new GammaProposer.
func NewGammaProposer(shape, rate float64) *GammaProposer {
	p, _ := prob.NewGamma(shape, rate)
	return &GammaProposer{p}
}

// Moments returns the statistics of the shape of the gamma
// probability density.
func (n *GammaProposer) Moments(names ...string) []float64 {
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

// Propose returns a random value from the initialized gamma probability density.
func (n *GammaProposer) Propose() float64 {
	return n.Random()
}

// Prob returns the probability of a value in the initialized gamma probability
// density.
func (n *GammaProposer) Prob(value float64) float64 {
	return n.Pdf(value)
}

// Probs returns probabilities of one or more given values based on the
// initialized gamma probability density.
func (n *GammaProposer) Probs(values ...float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		res[i] = n.Pdf(v)
	}
	return res
}

// LogProb returns the log probability of a value in the initialized gamma
// probability density.
func (n *GammaProposer) LogProb(value float64) float64 {
	return math.Log(n.Pdf(value))
}

// LogProbs returns the log probabilities of one or more given values based on
// the initialized gamma probability density.
func (n *GammaProposer) LogProbs(values ...float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		res[i] = math.Log(n.Pdf(v))
	}
	return res
}

// UpdateFields updates the mean, and standard deviation of the probability
// density of the gamma probability density depending on the number
// of values given. If only one value is given,
func (n *GammaProposer) UpdateFields(values ...float64) {
	n.Rate = values[0]
	if len(values) == 2 {
		n.Shape = values[1]
	}
}

// ExponentialProposer is a Proposer based on a exponential distribution with
// lambda parameter
type ExponentialProposer struct {
	prob.Exponential
}

// NewExpontialProposer is a constructor that returns a pointer to
// a new ExponentialProposer.
func NewExpontialProposer(lambda float64) *ExponentialProposer {
	p, _ := prob.NewExponential(lambda)
	return &ExponentialProposer{p}
}

// Moments returns the statistics of the shape of the exponential probability
// density.
func (n *ExponentialProposer) Moments(names ...string) []float64 {
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

// Propose returns a random value from the initialized exponential probability
// density.
func (n *ExponentialProposer) Propose() float64 {
	return n.Random()
}

// Prob returns the probability of a value in the initialized exponential
// probability density.
func (n *ExponentialProposer) Prob(value float64) float64 {
	return n.Pdf(value)
}

// Probs returns probabilities of one or more given values based on the
// initialized exponential probability density.
func (n *ExponentialProposer) Probs(values ...float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		res[i] = n.Pdf(v)
	}
	return res
}

// LogProb returns the log probability of a value in the initialized exponential
// probability density.
func (n *ExponentialProposer) LogProb(value float64) float64 {
	return math.Log(n.Pdf(value))
}

// LogProbs returns the log probabilities of one or more given values based on
// the initialized exponential probability density.
func (n *ExponentialProposer) LogProbs(values ...float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		res[i] = math.Log(n.Pdf(v))
	}
	return res
}

// UpdateFields updates the mean and standard deviation of
// the probability density of the exponential distribution depending on the number
// of values given.
func (n *ExponentialProposer) UpdateFields(values ...float64) {
	n.Lambda = values[0]
}

// LogNormalProposer is a Proposer based on a normal distribution with a mean
// mu and standard deviation sigma.
type LogNormalProposer struct {
	prob.LogNormal
}

// NewLogNormalProposer is a constructor that returns a pointer to
// a new LogNormalProposer.
func NewLogNormalProposer(mu, sigma float64) *LogNormalProposer {
	p, _ := prob.NewLogNormal(mu, sigma)
	return &LogNormalProposer{p}
}

// Moments returns the statistics of the shape of the lognormal probability
// density.
func (n *LogNormalProposer) Moments(names ...string) []float64 {
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

// Propose returns a random value from the initialized lognormal probability
// density.
func (n *LogNormalProposer) Propose() float64 {
	return n.Random()
}

// Prob returns the probability of a value in the initialized lognormal
// probability density.
func (n *LogNormalProposer) Prob(value float64) float64 {
	return n.Pdf(value)
}

// Probs returns probabilities of one or more given values based on the
// initialized lognormal probability density.
func (n *LogNormalProposer) Probs(values ...float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		res[i] = n.Pdf(v)
	}
	return res
}

// LogProb returns the log probability of a value in the initialized lognormal
// probability density.
func (n *LogNormalProposer) LogProb(value float64) float64 {
	return math.Log(n.Pdf(value))
}

// LogProbs returns the log probabilities of one or more given values based on
// the initialized lognormal probability density.
func (n *LogNormalProposer) LogProbs(values ...float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		res[i] = math.Log(n.Pdf(v))
	}
	return res
}

// UpdateFields updates the mean and standard deviation of
// the probability density of the normal distribution depending on the number
// of values given.
func (n *LogNormalProposer) UpdateFields(values ...float64) {
	n.Mu = values[0]
	if len(values) == 2 {
		n.Sigma = values[1]
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

package abcgo

// ParamSampler generates parameter values randomly.
type ParamSampler func(...float64) []float64

// Generator is a data generating function that outputs data based on
// a set of input parameters.
type Generator func(...float64) []float64

// Filter determines whether a set of parameter values are accepted of
// rejected.
type Filter func([]float64, []float64) bool

// RejectionSampler is a simple ABC method that accepts or rejects
// parameter sets based on a similarity metric computed between the observed data and
// the simulated data generated using the parameters.
func RejectionSampler(n int64, y []float64, generateParams ParamSampler, generateData Generator, acceptParams Filter) [][]float64 {
	// n is number of trials

	var acceptedParams [][]float64
	var params, x []float64
	for i := int64(0); i < n; i++ {
		for {
			// sample new parameters theta* from parameter generator
			params = generateParams()

			// generate new data X based on parameters theta
			x = generateData(params...)

			if acceptParams(y, x) == true {
				// store theta if within threshold
				acceptedParams = append(acceptedParams, params)
				break
			}
		}
	}
	return acceptedParams
}

// func MonteCarloSampler(n int, y []float64, epsilon float64) {
// 	// n is number of trials
// 	// epsilon is threshold for discrepancy
// 	// prior Y

// 	// discrepancy between data generated from theta -> X and
// 	// prior data Y
// 	discrepancyF := func(x, y []float64) (d float64) { return }
// 	// parameter sampler
// 	thetaF := func() (res []float64) { return }
// 	// Generative model using theta parameters
// 	generator := func(t []float64) (res []float64) { return }

// 	// accepted parameters
// 	// these generated data x that is within the threshold
// 	var accepted, weights [][]float64
// 	var theta, newTheta x []float64
// 	var discrepancy float64

// 	// Initial round
// 	for {
// 		// sample initial candidate parameters theta* from parameter generator
// 		// based on some distribution
// 		theta = thetaF()

// 		// generate new dataset X based on parameters theta
// 		// such that the number of observations in observed dataset y is the same
// 		// as the number of observations in generated dataset x
// 		x = generator(theta)

// 		// calculate discrepancy between observed dataset y and new dataset x
// 		// using some distance function
// 		discrepancy = discrepancyF(x, y)

// 		// if dataset x is "close enough" to dataset y based on the selected
// 		// metric, break and accept the sampled parameters
// 		if discrepancy > epsilon {
// 			continue
// 		}
// 		break
// 	}
// 	// store theta if within threshold
// 	accepted = append(accepted, theta)

// 	// store weights
// 	for i := 0; i < len(theta); i++ {
// 		weights[0] = append(weights[0], 1/float64(len(theta)))
// 	}

// 	// After first round
// 	for i := 1; i < n; i++ {
// 		for {
// 			// sample candidate parameters theta* from parameter generator
// 			// based on the previous distribution that generated the previously
// 			// accepted theta
// 			theta = thetaF()

// 			// perturb the parameter distribution
// 			// sample parameters from parameter generator using the previously
// 			// accepted parameters as the means of the new distribution
// 			newTheta = thetaF()

// 			// generate new dataset X based on parameters newTheta
// 			// such that the number of observations in observed dataset y is the same
// 			// as the number of observations in generated dataset x
// 			x = generator(newTheta)

// 			// calculate discrepancy between observed dataset y and new dataset x
// 			// using some distance function
// 			discrepancy = discrepancyF(x, y)

// 			// if dataset x is "close enough" to dataset y based on the selected
// 			// metric, break and accept the sampled parameters
// 			if discrepancy > epsilon {
// 				continue
// 			}
// 			break
// 		}
// 	}
// }

package abcgo

func RejectionSampler(n int64, y, epsilon float64) {
	// n is number of trials
	// epsilon is threshold for discrepancy
	// observed data Y

	// discrepancy between data generated from theta -> X and
	// observed data Y
	discrepancyF := func(x, y float64) (d float64) { return }
	// parameter sampler
	thetaF := func() (res float64) { return }
	// Generative model using theta parameters
	generator := func(t float64) (res float64) { return }

	// accepted parameters
	// these generated data x that is within the threshold
	var accepted []float64
	var theta, x float64
	var discrepancy float64
	for i := int64(0); i < n; i++ {
		for {
			// sample new parameters theta* from parameter generator
			theta = thetaF()

			// generate new data X based on parameters theta
			x = generator(theta)

			// calculate discrepancy between observed data y and new data x
			discrepancy = discrepancyF(x, y)

			if discrepancy > epsilon {
				continue
			}
			break
		}
		// store theta if within threshold
		accepted = append(accepted, theta)
	}
}

func MonteCarloSampler(n int, y []float64, epsilon float64) {
	// n is number of trials
	// epsilon is threshold for discrepancy
	// prior Y

	// discrepancy between data generated from theta -> X and
	// prior data Y
	discrepancyF := func(x, y []float64) (d float64) { return }
	// parameter sampler
	thetaF := func() (res []float64) { return }
	// Generative model using theta parameters
	generator := func(t []float64) (res []float64) { return }

	// accepted parameters
	// these generated data x that is within the threshold
	var accepted, weights [][]float64
	var theta, newTheta x []float64
	var discrepancy float64

	// Initial round
	for {
		// sample initial candidate parameters theta* from parameter generator
		// based on some distribution
		theta = thetaF()

		// generate new dataset X based on parameters theta
		// such that the number of observations in observed dataset y is the same
		// as the number of observations in generated dataset x
		x = generator(theta)

		// calculate discrepancy between observed dataset y and new dataset x
		// using some distance function
		discrepancy = discrepancyF(x, y)

		// if dataset x is "close enough" to dataset y based on the selected
		// metric, break and accept the sampled parameters
		if discrepancy > epsilon {
			continue
		}
		break
	}
	// store theta if within threshold
	accepted = append(accepted, theta)

	// store weights
	for i := 0; i < len(theta); i++ {
		weights[0] = append(weights[0], 1/float64(len(theta)))
	}

	// After first round
	for i := 1; i < n; i++ {
		for {
			// sample candidate parameters theta* from parameter generator
			// based on the previous distribution that generated the previously
			// accepted theta
			theta = thetaF()

<<<<<<< HEAD
			// perturb the parameter distribution
			// sample parameters from parameter generator using the previously
			// accepted parameters as the means of the new distribution
			newTheta = thetaF()

			// generate new dataset X based on parameters newTheta
			// such that the number of observations in observed dataset y is the same
			// as the number of observations in generated dataset x
			x = generator(newTheta)
=======
			// generate new dataset X based on parameters theta
			// such that the number of observations in observed dataset y is the same
			// as the number of observations in generated dataset x
			x = generator(theta)
>>>>>>> 501182d06860809cc8e6c2ea04380bc02fe3c0b8

			// calculate discrepancy between observed dataset y and new dataset x
			// using some distance function
			discrepancy = discrepancyF(x, y)

			// if dataset x is "close enough" to dataset y based on the selected
			// metric, break and accept the sampled parameters
			if discrepancy > epsilon {
				continue
			}
			break
		}
	}
}

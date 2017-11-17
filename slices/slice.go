package slices

func Sum1(s [4]float64) (sum float64) {
	s[0] = 100.0
	for _, v := range s {
		sum += v
	}
	return
}

func Sum2(s *[4]float64) (sum float64) {
	s[0] = 100.0
	for _, v := range *s {
		sum += v
	}
	return
}

func Sum3(s []float64) float64 {
	sum := 0.0
	s[0] = 100.0
	for _, v := range s {
		sum += v
	}
	return sum
}

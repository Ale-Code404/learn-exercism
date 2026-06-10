package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	if len(divisors) == 0 {
		return 0
	}

	hash := make(map[int]bool)

	allowed := make(map[int]bool, len(divisors))
	for _, divisor := range divisors {
		if divisor != 0 {
			allowed[divisor] = true
		}
	}

	i := 1

	for {
		for divisor := range allowed {
			multiple := divisor * i

			if multiple >= limit {
				delete(allowed, divisor)
				continue
			}

			hash[multiple] = true
		}

		if len(allowed) == 0 {
			break
		}

		i++
	}

	points := 0
	for key := range hash {
		points += key
	}

	return points
}

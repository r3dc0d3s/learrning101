package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i <= sqrt_approx(nb); i++ {
		if nb%i == 0 {
			return false
		}
		// return true
	}
	return true
}

func sqrt_approx(nb int) int {
	k := 1
	for ; k*k <= nb; k++ {
	}
	return k
}

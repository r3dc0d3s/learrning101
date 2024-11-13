package piscine

import "piscine"

func FindNextPrime(nb int) int {
	for i := nb; i <= 2*nb; i++ {
		if piscine.IsPrime(i) {
			return i
		}
	}
	return -1
}

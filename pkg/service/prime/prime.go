package prime

import (
	"math"
)

func isPrimeNumber(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPrimeNumberV2(n int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func FindLowerNearestPrimeNumber(n int) (int, error) {
	for i := n - 1; i > 1; i-- {
		if isPrimeNumber(i) {
			return i, nil
		}
	}
	return -1, nil
}

func FindLowerNearestPrimeNumberV2(n int) (int, error) {
	for i := n - 1; i > 1; i-- {
		// big.NewInt(int64(i)).ProbablyPrime(0)
		if isPrimeNumberV2(i) {
			return i, nil
		}
	}
	return -1, nil
}

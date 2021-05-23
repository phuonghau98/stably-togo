package prime

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

func FindLowerNearestPrimeNumber(n int) (error, int) {
	for i := n - 1; i > 1; i-- {
		if isPrimeNumber(i) {
			return nil, i
		}
	}
	return nil, -1
}

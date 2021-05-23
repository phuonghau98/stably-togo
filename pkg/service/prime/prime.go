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

func FindLowerNearestPrimeNumber(n int) (int, error) {
	for i := n - 1; i > 1; i-- {
		if isPrimeNumber(i) {
			return i, nil
		}
	}
	return -1, nil
}

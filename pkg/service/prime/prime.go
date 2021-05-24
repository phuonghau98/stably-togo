package prime

import (
	"fmt"
	"math/big"
)

func isPrimeNumber(n *big.Int) bool {
	two := big.NewInt(2)
	one := big.NewInt(1)
	zero := big.NewInt(0)
	if n.Cmp(one) <= 0 {
		return false
	}
	fmt.Println(n)
	for i := new(big.Int).Set(two); i.Cmp(n) < 0; i.Add(i, one) {
		if new(big.Int).Mod(n, i).Cmp(zero) == 0 {
			return false
		}
	}
	return true
}

func isPrimeNumberV2(n *big.Int) bool {
	sqrt := new(big.Int).Sqrt(n)
	two := big.NewInt(2)
	one := big.NewInt(1)
	zero := big.NewInt(0)
	for i := new(big.Int).Set(two); i.Cmp(sqrt) <= 0; i.Add(i, one) {
		fmt.Println("i", i, "n", n, new(big.Int).Mod(n, i))
		if new(big.Int).Mod(n, i).Cmp(zero) == 0 {
			return false
		}
	}
	return n.Cmp(one) > 0
}

func FindLowerNearestPrimeNumber(n *big.Int) (string, error) {
	one := big.NewInt(1)
	for i := n.Sub(n, one); i.Cmp(one) > 0; i.Sub(i, one) {
		if isPrimeNumber(i) {
			return i.String(), nil
		}
	}
	return "-1", nil
}

func FindLowerNearestPrimeNumberV2(n *big.Int) (string, error) {
	one := big.NewInt(1)
	for i := new(big.Int).Sub(n, one); i.Cmp(one) > 0; i.Sub(i, one) {
		// big.NewInt(int64(i)).ProbablyPrime(0)
		if isPrimeNumberV2(i) {
			return i.String(), nil
		}
	}
	return "-1", nil
}

func FindLowerNearestPrimeNumberOptimized(n *big.Int) (string, error) {
	one := big.NewInt(1)
	for i := new(big.Int).Sub(n, one); i.Cmp(one) > 0; i.Sub(i, one) {
		if i.ProbablyPrime(0) {
			return i.String(), nil
		}
	}
	return "-1", nil
}

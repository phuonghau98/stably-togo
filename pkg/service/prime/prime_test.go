package prime

import (
	"fmt"
	"log"
	"math/big"
	"testing"
)

func TestIsPrimeNumber(t *testing.T) {
	testName := func(n *big.Int, isPrime bool) string {
		return fmt.Sprintf("%v is prime number: %t", n, isPrime)
	}

	cases := []struct {
		input    *big.Int
		expected bool
	}{
		{input: big.NewInt(-1), expected: false},
		{input: big.NewInt(2), expected: true},
		{input: big.NewInt(3), expected: true},
		{input: big.NewInt(14), expected: false},
		{input: big.NewInt(60), expected: false},
		{input: big.NewInt(80), expected: false},
		{input: big.NewInt(97), expected: true},
		{input: big.NewInt(99), expected: false},
	}

	for _, c := range cases {
		t.Run(testName(c.input, c.expected), func(t *testing.T) {
			got := isPrimeNumber(c.input)
			want := c.expected

			if got != want {
				log.Fatalf("Wanted: %v, got: %v", want, got)
			}
		})
	}
}

func TestFindLowerNearestPrimeNumber(t *testing.T) {
	testName := func(input *big.Int, expected string) string {
		if expected == "-1" {
			return fmt.Sprintf("There shouldn't be a number that is a prime number and also lower than %v", input)
		}
		return fmt.Sprintf("The highest prime number lower than %v should be: %v", input, expected)
	}

	cases := []struct {
		input    *big.Int
		expected string
	}{
		{input: big.NewInt(2), expected: "-1"},
		{input: big.NewInt(3), expected: "2"},
		{input: big.NewInt(5), expected: "3"},
		{input: big.NewInt(14), expected: "13"},
		{input: big.NewInt(73), expected: "71"},
		{input: big.NewInt(97), expected: "89"},
		{input: big.NewInt(98), expected: "97"},
		{input: big.NewInt(2084), expected: "2083"},
		{input: big.NewInt(8999), expected: "8971"},
		{input: big.NewInt(1299827), expected: "1299821"},
		// {input: 10000003957, expected: 10000003931},
	}

	for _, c := range cases {
		t.Run(testName(c.input, c.expected), func(t *testing.T) {
			got, err := FindLowerNearestPrimeNumber(c.input)
			if err != nil {
				log.Fatalf("There shouldn't be an error, but one caught: %v", err.Error())
			}
			want := c.expected

			if got != want {
				log.Fatalf("Wanted: %v, got: %v", want, got)
			}
		})
	}
}

func TestFindLowerNearestPrimeNumberV2(t *testing.T) {
	testName := func(input *big.Int, expected string) string {
		if expected == "-1" {
			return fmt.Sprintf("There shouldn't be a number that is a prime number and also lower than %v", input)
		}
		return fmt.Sprintf("The highest prime number lower than %v should be: %v", input, expected)
	}

	cases := []struct {
		input    *big.Int
		expected string
	}{
		{input: big.NewInt(2), expected: "-1"},
		{input: big.NewInt(1299827), expected: "1299821"},
		{input: big.NewInt(10000003957), expected: "10000003931"},
		// {input: 100000039572313, expected: "100000039572263"},
		// {input: 1000000395723132323, expected: "1000000395723132283"},
	}

	for _, c := range cases {
		t.Run(testName(c.input, c.expected), func(t *testing.T) {
			got, err := FindLowerNearestPrimeNumberV2(c.input)
			if err != nil {
				log.Fatalf("There shouldn't be an error, but one caught: %v", err.Error())
			}
			want := c.expected

			if got != want {
				log.Fatalf("Wanted: %v, got: %v", want, got)
			}
		})
	}
}

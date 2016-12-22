// pi prints some number of digits of pi and the time taken by the calculation.
package main

import (
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	digits := digitsFromCommandLine(1000)
	scaledPi := fmt.Sprint(π(digits))
	fmt.Printf("3.%s\n", scaledPi[1:])
	fmt.Println("Duration:", time.Since(start))
}

func digitsFromCommandLine(default_digits int64) int64 {
	if len(os.Args) > 2 || (len(os.Args) > 1 &&
		(os.Args[1] == "-h" || os.Args[1] == "--help")) {
		exitWithUsage()
	}
	if len(os.Args) > 1 {
		if places, err := strconv.ParseInt(os.Args[1], 0, 64); err != nil {
			fmt.Printf("%s: invalid number %s\n",
				filepath.Base(os.Args[0]), os.Args[1])
			exitWithUsage()
		} else {
			return places
		}
	}
	return default_digits
}

func exitWithUsage() {
	fmt.Printf("usage: %s <number-of-digits>\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func π(digits int64) *big.Int {
	extraDigits := int64(10)
	extraScale := new(big.Int).Exp(big.NewInt(10), big.NewInt(extraDigits), nil)
	pi := new(big.Int).Mul(big.NewInt(4), arccot(5, digits+extraDigits))
	pi.Sub(pi, arccot(239, digits+extraDigits))
	pi.Mul(pi, big.NewInt(4))
	pi.Div(pi, extraScale)
	return pi
}

func arccot(x int64, digits int64) *big.Int {
	unity := new(big.Int).Exp(big.NewInt(10), big.NewInt(digits), nil)
	coefficient := big.NewInt(1)
	bigX := big.NewInt(x)
	factor := new(big.Int).Div(unity, bigX)

	negativeXSquared := big.NewInt(0).Mul(bigX, bigX)
	negativeXSquared.Neg(negativeXSquared)

	bigTwo := big.NewInt(2)
	bigZero := big.NewInt(0)
	term := big.NewInt(0)
	sum := big.NewInt(0)

	for factor.Cmp(bigZero) != 0 {
		term.Div(factor, coefficient)
		sum.Add(sum, term)
		coefficient.Add(coefficient, bigTwo)
		factor.Div(factor, negativeXSquared)
	}
	return sum
}

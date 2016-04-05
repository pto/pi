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

func digitsFromCommandLine(default_digits int) int {
	if len(os.Args) > 2 || (len(os.Args) > 1 &&
		(os.Args[1] == "-h" || os.Args[1] == "--help")) {
		exitWithUsage()
	}
	if len(os.Args) > 1 {
		if places, err := strconv.Atoi(os.Args[1]); err != nil {
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

func π(digits int) *big.Int {
	factor := big.NewInt(int64(digits))
	ten := big.NewInt(10)
	unity_exponent := big.NewInt(0).Add(factor, ten)
	unity := big.NewInt(0).Exp(ten, unity_exponent, nil)
	left := big.NewInt(0).Mul(big.NewInt(4), arccot(big.NewInt(5), unity))
	right := arccot(big.NewInt(239), unity)
	pi := big.NewInt(0).Mul(big.NewInt(4), big.NewInt(0).Sub(left, right))
	return pi.Div(pi, big.NewInt(0).Exp(ten, ten, nil))
}

func arccot(x, unity *big.Int) *big.Int {
	zero := big.NewInt(0)
	minus_one := big.NewInt(-1)
	x_squared := big.NewInt(0).Mul(x, x)
	coefficient := big.NewInt(1)
	two := big.NewInt(2)
	divisor := big.NewInt(0).Div(unity, x)
	sum := big.NewInt(0)
	factor := big.NewInt(0).Div(divisor, coefficient)
	for factor.Cmp(zero) != 0 {
		sum.Add(sum, factor)
		coefficient.Add(coefficient, two)
		divisor.Mul(divisor, minus_one)
		divisor.Div(divisor, x_squared)
		factor.Div(divisor, coefficient)
	}
	return sum
}

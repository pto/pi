// pi prints some number of digits of pi and the time taken by the calculation.
package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"runtime/pprof"
	"strconv"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write CPU profile to a file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	start := time.Now()
	digits := digitsFromCommandLine(1000)
	scaledPi := fmt.Sprint(π(digits))
	fmt.Printf("3.%s\n", scaledPi[1:])
	fmt.Println("Duration:", time.Since(start))
}

func digitsFromCommandLine(default_digits int) int {
	args := flag.Args()
	if len(args) > 1 || (len(args) == 1 &&
		(args[0] == "-h" || args[0] == "--help")) {
		exitWithUsage()
	}
	if len(args) == 1 {
		if places, err := strconv.Atoi(args[0]); err != nil {
			fmt.Printf("%s: invalid number %s\n",
				filepath.Base(os.Args[0]), args[0])
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
	precision := big.NewInt(int64(digits))
	ten := big.NewInt(10)
	unity_exponent := big.NewInt(0).Add(precision, ten)
	unity := big.NewInt(0).Exp(ten, unity_exponent, nil)
	left := big.NewInt(0).Mul(big.NewInt(4), arccot(big.NewInt(5), unity))
	right := arccot(big.NewInt(239), unity)
	pi := big.NewInt(0).Mul(big.NewInt(4), big.NewInt(0).Sub(left, right))
	return pi.Div(pi, big.NewInt(0).Exp(ten, ten, nil))
}

func arccot(x, unity *big.Int) *big.Int {
	zero := big.NewInt(0)
	minus_x_squared := big.NewInt(0).Mul(x, x)
	minus_x_squared.Neg(minus_x_squared)
	coefficient := big.NewInt(1)
	two := big.NewInt(2)
	term := big.NewInt(0).Div(unity, x)
	sum := big.NewInt(0)
	divisor := big.NewInt(0)
	for term.Cmp(zero) != 0 {
		sum.Add(sum, term)
		term.Mul(term, coefficient)
		coefficient.Add(coefficient, two)
		divisor.Mul(minus_x_squared, coefficient)
		term.Quo(term, divisor)
	}
	return sum
}

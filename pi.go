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
	fmt.Printf("%s.%s\n", scaledPi[0:1], scaledPi[1:])
	fmt.Println("Duration:", time.Since(start))
}

func digitsFromCommandLine(default_digits int64) int64 {
	args := flag.Args()
	if len(args) > 1 || (len(args) == 1 &&
		(args[0] == "-h" || args[0] == "--help")) {
		exitWithUsage()
	}
	if len(args) == 1 {
		if places, err := strconv.ParseInt(args[0], 0, 64); err != nil {
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

func π(digits int64) *big.Int {
	const extraDigits = 10
	extraScale := new(big.Int).Exp(big.NewInt(10), big.NewInt(extraDigits), nil)
	pi := new(big.Int).Mul(big.NewInt(4), arccot(5, digits+extraDigits))
	pi.Sub(pi, arccot(239, digits+extraDigits))
	pi.Mul(pi, big.NewInt(4))
	pi.Div(pi, extraScale)
	return pi
}

func arccot(littleX int64, digits int64) *big.Int {
	unity := new(big.Int).Exp(big.NewInt(10), big.NewInt(digits), nil)
	x := big.NewInt(littleX)
	term := new(big.Int).Div(unity, x)
	coefficient := big.NewInt(1)
	negativeXSquared := new(big.Int).Neg(new(big.Int).Mul(x, x))
	zero := big.NewInt(0)
	two := big.NewInt(2)
	sum := big.NewInt(0)
	divisor := new(big.Int)

	for term.Cmp(zero) != 0 {
		sum.Add(sum, term)
		term.Mul(term, coefficient)
		coefficient.Add(coefficient, two)
		divisor.Mul(negativeXSquared, coefficient)
		term.Quo(term, divisor)
	}
	return sum
}

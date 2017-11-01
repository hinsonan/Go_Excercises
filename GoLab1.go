package main

import (
	"fmt"
	"math"
	"strings"
)

const delta = 1e-6

func Sqrt(x float64) float64 {
	z := x
	n := 0.0
	for math.Abs(n-z) > delta {
		n, z = z, z-(z*z-x)/(2*z)
	}
	return z
}

//excercise 2
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, f := range strings.Fields(s) {
		m[f]++
	}
	return m
}

//excercise 3
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	//excercise 1
	const x = 2
	mine, theirs := Sqrt(x), math.Sqrt(x)
	fmt.Println(mine, theirs, mine-theirs)
	//excercise 2
	fmt.Println(WordCount("this is a string"))
	//excercise 3
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

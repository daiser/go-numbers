package main

import (
	"daiser/numbers/loco"
	"os"
	"strconv"
)

func main() {
	var values []float64
	args := os.Args[1:]

	if len(args) < 2 {
		os.Stderr.WriteString("not enough input\n")
		os.Exit(-2)
	}

	for _, arg := range args {
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			os.Stderr.WriteString("invalid float " + arg + "\n")
			os.Exit(-1)
		}

		values = append(values, value)
	}

	mean, sd := loco.StDev(values)
	println(mean, sd)
}

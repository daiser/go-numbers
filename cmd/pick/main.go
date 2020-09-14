package main

import (
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]

	lenArgs := len(args)
	if lenArgs < 1 {
		os.Exit(1)
	}
	if lenArgs == 0 {
		println(args[0])
		os.Exit(0)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	println(args[rand.Intn(lenArgs)])
}

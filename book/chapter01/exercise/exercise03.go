package main

import (
	"os"
	"strings"
)

func exercise03_slow() string {
	var echo, separator string
	var limite = len(os.Args)
	for i := 1; i < limite; i++ {
		echo += separator + os.Args[i]
		separator = " "
	}
	return echo
}

func exercise03_fast() string {
	return strings.Join(os.Args[1:], " ")
}

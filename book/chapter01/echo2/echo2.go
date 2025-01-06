package main

import (
	"fmt"
	"os"
)

func main() {
	echo, separator := "", ""
	for _, arg := range os.Args[1:] {
		echo += separator + arg
		separator = " "
	}
	fmt.Println(echo)
}

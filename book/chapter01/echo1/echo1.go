package main

import (
	"fmt"
	"os"
)

func main() {
	var echo, separator string
	var limite = len(os.Args)
	for i := 1; i < limite; i++ {
		echo += separator + os.Args[i]
		separator = " "
	}
	fmt.Println(echo)
}

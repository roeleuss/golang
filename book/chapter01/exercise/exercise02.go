package main

import (
	"fmt"
	"os"
	"strconv"
)

func exercise02() {
	for i, arg := range os.Args {
		fmt.Println(strconv.Itoa(i) + ": " + arg)
	}

}

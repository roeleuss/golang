package main

import (
	"fmt"
	"os"
	"strings"
)

func exercise01() {
	fmt.Println(strings.Join(os.Args, " "))
}

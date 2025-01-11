package main

import (
	"bufio"
	"fmt"
	"os"
)

func exercise04() {
	counts := make(map[string]Exercise04)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}

	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%d\t%s\n", n.count, line)
			for i := 0; i < n.count; i++ {
				fmt.Printf("\t%s : %d\n", n.file[i], n.line[i])
			}
		}
	}
}

func countLines(f *os.File, counts map[string]Exercise04) {
	input := bufio.NewScanner(f)
	for line := 1; input.Scan(); line++ {
		data, exist := counts[input.Text()]
		if !exist {
			data = Exercise04{count: 0, file: []string{}, line: []int{}}
		}
		data.count++
		data.file = append(data.file, f.Name())
		data.line = append(data.line, line)
		counts[input.Text()] = data
	}
}

type Exercise04 struct {
	count int
	file  []string
	line  []int
}

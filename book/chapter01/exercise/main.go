package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		input := chooseExercise(reader)
		switch input {
		case 0:
			fmt.Println("Thank you for playing!")
			return
		case 1:
			exercise01()
		case 2:
			exercise02()
		case 3:
			slow := measureExecutionTime(exercise03_slow)
			fast := measureExecutionTime(exercise03_fast)
			fmt.Printf("Execution time (%s): %s\n", "slow", slow)
			fmt.Printf("Execution time (%s): %s\n", "fast", fast)
		case 4:
			exercise04()
		case 5:
			exercise05()
		case 6:
			exercise06()
		default:
			fmt.Println("Invalid exercise number. Please choose a number between 1 and 6.")
		}
	}
}

func chooseExercise(reader *bufio.Reader) int {
	fmt.Print("Choose a exercise number [1-6] to run or type 0 to exit: ")
	input, err := reader.ReadString('\n')
	throw(err, "An error occurred while reading the input.")
	input = strings.TrimSpace(input)
	exerciseNumber, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number between 1 and 6.")
		return chooseExercise(reader)
	}
	return exerciseNumber
}

func measureExecutionTime(f func() string) time.Duration {
	start := time.Now()
	for i := 0; i < 100000000; i++ {
		_ = f()
	}
	return time.Since(start)
}

func throw(err error, message string) {
	if err != nil {
		fmt.Println(message)
		panic(err)
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var fns = map[string]func(int, int) int{
	"+": func(a int, b int) int {
		return a + b
	},
	"-": func(a int, b int) int {
		return a - b
	},
	"*": func(a int, b int) int {
		return a * b
	},
	"/": func(a int, b int) int {
		return a / b
	},
}

func main() {
	fmt.Println("Welcome to Dumb Lisp!")
	for {
		fmt.Print("\nλ ")
		r := bufio.NewReader(os.Stdin)
		input, _ := r.ReadString('\n')

		if input == "(exit)\n" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		var stack nestStack

		for i, r := range input {
			switch r {
			case '(':
				stack.Push(i)

			case ')':
				iOpen, err := stack.Pop()
				if err != nil {
					log.Fatal(err)
				}
				innerSexp := input[iOpen : i+1]
				eval(innerSexp)
			}
		}
	}
}

func eval(sexp string) {
	atoms := strings.Split(strings.Trim(sexp, "()"), " ")
	fnStr := atoms[0]

	arg1Str := atoms[1]
	arg1, err := strconv.Atoi(arg1Str)
	if err != nil {
		log.Fatal(err)
	}

	arg2Str := atoms[2]
	arg2, err := strconv.Atoi(arg2Str)
	if err != nil {
		log.Fatal(err)
	}

	fn := fns[fnStr]
	fmt.Println(fn(arg1, arg2))
}

type nestStack []int

func (s *nestStack) Push(val int) {
	*s = append(*s, val)
}

func (s *nestStack) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, fmt.Errorf("Empty stack")
	}

	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return popped, nil
}

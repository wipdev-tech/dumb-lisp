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
		sexp, _ := r.ReadString('\n')

		if strings.TrimSpace(sexp) == "" {
			continue
		}

		if sexp == "(exit)\n" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		eval(sexp)

	}
}

func eval(sexp string) {
	iOpen := 0
	iClose := 0

	for i, r := range sexp {
		if r == '(' {
			iOpen = i
		} else if r == ')' {
			iClose = i
            break
		}
	}

	innerSexp := sexp[iOpen : iClose+1]
	result := evalInner(innerSexp)

	if iOpen == 0 {
		fmt.Printf("  %v\n", result)
		return
	}

	sexp = strings.Replace(
		sexp,
		innerSexp,
		fmt.Sprintf("%v", result),
		1,
	)
	fmt.Printf("  %s", sexp)
    eval(sexp)
}

func evalInner(sexp string) int {
	atoms := strings.Fields(strings.Trim(sexp, "()"))
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
	return fn(arg1, arg2)
}

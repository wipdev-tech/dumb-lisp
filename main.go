package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    r := bufio.NewReader(os.Stdin)
    fmt.Print("λ ")
    input, _ := r.ReadString('\n')
    println(input)
}


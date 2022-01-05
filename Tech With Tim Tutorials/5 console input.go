package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("type your birth year: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64) // if error, underscore stores the error. we use underscore here cuz we don't care about it now.
	fmt.Printf("you'll be %d years old at the end of 2022", 2022-input)

}

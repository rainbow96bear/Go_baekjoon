package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	A := Input(scanner)[0]
	B := Input(scanner)[0]

	fmt.Fprintln(writer, A+B)
}

func Input(scanner *bufio.Scanner) []int {
	scanner.Scan()

	input := strings.Split(scanner.Text(), " ")

	result := make([]int, len(input))

	for i, v := range input {
		result[i], _  = strconv.Atoi(v)
	}

	return result
}
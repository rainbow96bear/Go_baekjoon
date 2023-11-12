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

	M, _ := strconv.Atoi(Input(scanner))
	
	X := 1000000007
	
	answer := int64(0)
	for i:= 0 ; i<M ; i++ {
		input := strings.Split(Input(scanner), " ")
		N, _ := strconv.Atoi(input[0])
		S, _ := strconv.Atoi(input[1])
		
		answer += (int64(S) * Calc(N, X)) % int64(X)
		answer %= int64(X)
	}
	fmt.Fprintln(writer, answer)
}

func Calc(N, X int) int64 {
	x := X-2
	n := int64(N%X)
	result := int64(1)
	for x > 0 {
		if x%2 == 1 {
			result *= n
			result %= int64(X)
		}
		x /= 2
		n = (n*n) % int64(X)
	}
	return result%int64(X)
}

func Input(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}

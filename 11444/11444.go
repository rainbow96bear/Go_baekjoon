package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	num := Input(scanner)

	result := Process(num)

	fmt.Fprintln(writer, result)
}

func Process(num int) int {
	base := [][]int{{1, 1}, {1, 0}}
	key := [][]int{{1, 0}, {0, 1}}
	check := num
	div := 1000000007
	for check > 0 {
		if check % 2 == 1 {
			key[0] = []int{((key[0][0]*base[0][0])%div + (key[0][1]*base[1][0])%div)%div, ((key[0][0]*base[0][1])%div + (key[0][1]*base[1][1])%div)%div}
			key[1] = []int{((key[1][0]*base[0][0])%div + (key[1][1]*base[1][0])%div)%div, ((key[1][0]*base[0][1])%div + (key[1][1]*base[1][1])%div)%div}
		}
		base = [][]int{{((base[0][0]*base[0][0])%div + (base[0][1]*base[1][0])%div)%div, ((base[0][0]*base[0][1])%div + (base[0][1]*base[1][1])%div)%div}, {((base[1][0]*base[0][0])%div + (base[1][1]*base[1][0])%div)%div, ((base[1][0]*base[0][1])%div + (base[1][1]*base[1][1])%div)%div}}
		check /= 2
	}
	return key[1][0]
}

func Input(scanner *bufio.Scanner) int {
	scanner.Scan()

	input := scanner.Text()

	result, _ := strconv.Atoi(input)

	return result

}
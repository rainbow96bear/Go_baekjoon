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

	N, _ := strconv.Atoi(Input(scanner))

	list := strings.Split(Input(scanner), " ")
	min := 1000000
	max := 0
	for i:=0 ; i<N ; i++ {
		num, _ := strconv.Atoi(list[i])
		if min > num {
			min = num
		}
		if max < num {
			max = num
		}
	}
	fmt.Fprintln(writer, min*max)
}
func Input(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}
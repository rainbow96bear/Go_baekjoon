package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var input string
	fmt.Fscan(reader, &input)

	items := strings.Split(input, "-")
	sum := 0
	for i, v := range items {
		
		nums := strings.Split(v, "+")
		for _, num := range nums {
			if i == 0 {
				intNum, _ := strconv.Atoi(num)
				sum += intNum
			}else {
				intNum, _ := strconv.Atoi(num)
				sum -= intNum
			}
		}
	}
	fmt.Fprintln(writer, sum)
}

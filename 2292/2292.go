package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)
	answer := 1
	check := 1
	for i:=1; check<N ; i++{
		check = check + 6*i
		answer++
	}
	fmt.Fprint(writer, answer)
}
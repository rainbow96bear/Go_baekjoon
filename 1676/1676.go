package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)
	answer := 0
	for i:=1 ; int(math.Pow(5.0, float64(i))) <= N ; i++ {
		answer += N/int(math.Pow(5.0, float64(i)))
	}
	fmt.Fprint(writer, answer)
}

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

	var N, K int
	fmt.Fscan(reader, &N, &K)
	for i:=N ;  ; i++ {
		check := 0
		testNum := i
		for testNum > 0 {
			if testNum%2 == 1 {
				check++
			}
			testNum /= 2
		}
		if check <= K {
			fmt.Fprint(writer, i-N)
			break
		}
	}
}
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

	var T int
	fmt.Fscan(reader, &T)

	for i:=0 ; i<T ; i++ {
		var M, N, x, y int
		fmt.Fscan(reader, &M, &N, &x, &y)
		
		answer := x
		isFind := false
		for answer <= M*N {
			if (answer-y)%N == 0 {
				fmt.Fprintln(writer, answer)
				isFind = true
				break
			}
			answer += M
		}
		if isFind == false {
			fmt.Fprintln(writer, -1)
		}
	}
}
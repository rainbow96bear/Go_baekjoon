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

	houses := make([][3]int,N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &houses[i][0], &houses[i][1], &houses[i][2])
	}
	sum := make([][3]int,N)
	sum[0] = houses[0]
	for i:=1 ; i<N ; i++ {
		for j:=0 ; j<3 ; j++ {
			sum[i][j] = houses[i][j] + min(sum[i-1][(j+1)%3], sum[i-1][(j+2)%3])
		}
	}
	answer := min(sum[N-1][0], sum[N-1][1])
	answer = min(answer, sum[N-1][2])

	fmt.Fprint(writer, answer)
}

func min(a, b int) int {
	if a < b{
		return a
	}
	return b
}
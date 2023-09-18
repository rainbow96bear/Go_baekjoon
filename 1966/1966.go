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

	var testCase int
	fmt.Fscan(reader, &testCase)

	for i:=0 ; i<testCase ; i++ {
		var N, index int
		fmt.Fscan(reader, &N, &index)

		queue := make([]int, N)
		max := 0
		for j:=0 ; j<N ; j++ {
			fmt.Fscan(reader, &queue[j])
			if max < queue[j] {
				max = queue[j]
			}
		}

		answer := 0
		for j:=max ; j>queue[index] ; j-- {
			for k:=0 ; k<len(queue) ; k++ {
				if queue[k]==j {
					if k < index {
						index = index - k -1
					}else {
						index = index + len(queue[k+1:])
					}
					queue = append(queue[k+1:], queue[:k]...)
					answer++
					k = -1
				}
			}
		}
		
		for j:=0 ; j<=index ; j++ {
			if queue[index] == queue[j]{
				answer ++
			}
		}
		fmt.Fprintln(writer, answer)
	}
}
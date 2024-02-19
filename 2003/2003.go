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

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	list := make([]int, N)
	for i, _ := range list {
		fmt.Fscanf(reader, "%d", &list[i])
	}
	fmt.Fscanf(reader, "\n")
	start, end, answer, total := 0, 0, 0, 0

	for start < N {
		if total < M {
			if end == N {
				break
			}
			total += list[end]
			end++
		}else {
			if total == M {
				answer++
			}
			total -= list[start]
			start++
		}
	}
	fmt.Fprintln(writer, answer)
}
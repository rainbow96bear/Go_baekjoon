package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, M int
	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	list := make([]int, M)
	for i:=0 ; i<M ; i++ {
		fmt.Fscanf(reader, "%d\n", &list[i])
	}
	sort.Slice(list, func(i, j int) bool {return list[i]<list[j]})
	num, totalPrice := process(N, M, list)
	fmt.Fprintf(writer, "%d %d\n", num, totalPrice) 
}

func process(N, M int, list []int) (num, totalPrice int) {
	max := 0
	point := 0
	sum := make([]int, M)
	for i := 0; i < M; i++ {
		if M > N {
			if i < (M - N + 1) {
				sum[i] = list[i] * N
			} else {
				sum[i] = list[i] * (M - i)
			}
		} else {
			sum[i] = list[i] * (M - i)
		}
	}

	max = sum[0]
	for i := 1; i < M; i++ {
		if sum[i] > max {
			max = sum[i]
			point = i
		}
	}
	return list[point], max
}
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, S int
	fmt.Fscanf(reader, "%d\n", &N)

	list := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &list[i])
	}
	fmt.Fscanf(reader, "\n")

	fmt.Fscanf(reader, "%d\n", &S)
	result := Sort(N, S, list)
	for _, v := range result {
		fmt.Fprintf(writer, "%d ", v)  // 수정: 문자열 출력을 정수 출력으로 변경
	}
}

func Sort(N, S int, list []int) []int {  // 수정: 문자열 슬라이스를 정수 슬라이스로 변경
	startIndex := 0
	for S > 0 && startIndex < N {
		maxIndex := startIndex
		for j := startIndex; j < N && j <= startIndex+S; j++ {
			if list[maxIndex] < list[j] {
				maxIndex = j
			}
		}
		for j := maxIndex; j > startIndex; j-- {
			list[j], list[j-1] = list[j-1], list[j]
			S--
		}
		startIndex++
	}
	return list
}
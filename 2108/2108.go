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

	var N int
	fmt.Fscan(reader, &N)

	NumList := make([]int, N)
	Check := make(map[int]int)
	total := 0
	for i:=0 ; i<N ; i++{
		fmt.Fscan(reader, &NumList[i])
		Check[NumList[i]]++
		total += NumList[i]
	}

	sort.Ints(NumList)

	// 평균
	mod := float64(total%N)/float64(N)
	if mod < 0 && mod < -0.5{
		fmt.Fprintln(writer, total/N - 1)
	}else if mod >= 0 && mod >= 0.5 {
		fmt.Fprintln(writer, total/N + 1)
	}else {
		fmt.Fprintln(writer, total/N)
	}

	// 중앙값
	fmt.Fprintln(writer, NumList[N/2])
	
	// 최빈값
	max := 0
	maxList := []int{}
	for i:=0 ; i<N ; i++ {
		if max < Check[NumList[i]]{
			maxList = []int{NumList[i]}
			max = Check[NumList[i]]
			continue
		}
		if max == Check[NumList[i]] && maxList[len(maxList)-1] != NumList[i]{
			maxList = append(maxList, NumList[i])
		}
	}
	sort.Ints(maxList)
	if len(maxList) > 1 {
		fmt.Fprintln(writer, maxList[1])
	}else {
		fmt.Fprintln(writer, maxList[0])
	}

	// 범위
	fmt.Fprintln(writer, NumList[len(NumList)-1]-NumList[0])
}
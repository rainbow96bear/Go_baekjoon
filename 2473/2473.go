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
	fmt.Fscanf(reader, "%d\n", &N)

	liquidList := make([]int, N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d", &liquidList[i])
	}
	sort.Slice(liquidList, func(i, j int) bool {return liquidList[i] < liquidList[j]})
	result := makeLiquid(N, liquidList)
	for _, num := range result {
		fmt.Fprintf(writer, "%d ", num)
	}
}

func makeLiquid(N int, liquidList []int) []int {
	bowl := make([]int,3)
	minValue := int64(50000000000001)
	out : for i:=0 ; i<N-2 ; i++ {
		s, e := i+1, N-1
		for s < e {
			resultValue := int64(liquidList[i] + liquidList[s] + liquidList[e])
			if abs(resultValue) < abs(minValue) {
				minValue = resultValue
				bowl[0] = liquidList[i]
				bowl[1] = liquidList[s]
				bowl[2] = liquidList[e]
			}
			if resultValue == 0 {
				break out
			}else if resultValue < 0 {
				s++
			}else {
				e--
			}
		}
	}
	return bowl
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}
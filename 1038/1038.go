package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var list []int
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	result := FindDecreaseNum(N)
	fmt.Fprintln(writer, result)
}

func FindDecreaseNum(N int) int {
	if N<=10 {
		return N
	}
	if N > 1022 {
		return -1
	}
	for i:=0 ; i<10 ; i++ {
		bp(i, 1)
	}
	sort.Slice(list, func(i, j int) bool {return list[i] < list[j]})
	return list[N]
}

func bp(num int, idx int) {
	if idx > 10 {
		return
	}
	list = append(list, num)
	for i:=0 ; i<num%10 ; i++ {
		bp((num*10)+i, idx+1)
	}
}

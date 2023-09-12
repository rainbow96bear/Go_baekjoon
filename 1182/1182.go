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

	var N, S int
	fmt.Fscan(reader, &N, &S)
	
	Nums := make([]int,N)
	for i:=0 ; i<N ; i++ {
		fmt.Fscan(reader, &Nums[i])
	}
	cnt:=0
	cnt = calc(Nums, 0, 0, S, false)
	if S == 0 {
		cnt--
	}


	fmt.Fprintln(writer, cnt)
}

func calc(Nums []int, index, currentSum, S int, isSelected bool) int {
	cnt := 0
	if index == len(Nums) {
		if currentSum == S {
			cnt++
		}
	} else {
		cnt += calc(Nums, index+1, currentSum+Nums[index], S, true)
		cnt += calc(Nums, index+1, currentSum, S, false)
	}
	return cnt
}
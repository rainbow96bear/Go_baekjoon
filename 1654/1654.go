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

	var K, N int
	fmt.Fscan(reader, &K, &N)

	lengthList := make([]int,K)
	for i:=0 ; i<K ; i++ {
		fmt.Fscan(reader, &lengthList[i])
	}
	sort.Ints(lengthList)

	min := 1
	max := lengthList[len(lengthList)-1]
	answer := 0
	for min<=max {
		mid := (min+max)/2
		cnt := 0
		for i:=0 ; i<K ; i++{
			cnt += (lengthList[i]/mid)
		}
		if cnt < N {
			max = mid-1
		}else {
			if answer < mid {
				answer = mid
			}
			min = mid+1
		}
	}
	fmt.Fprintln(writer, answer)
}
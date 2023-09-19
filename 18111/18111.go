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

	var N, M, B int
	fmt.Fscan(reader, &N, &M, &B)

	minHeight := 256
	maxHeight := 0
	field := make([][]int, N)
	for i:=0 ; i<N ; i++ {
		for j:=0 ; j<M ; j++ {
			var input int
			fmt.Fscan(reader, &input)
			field[i] = append(field[i], input)
			if minHeight > input {
				minHeight = input
			} 
			if maxHeight < input{
				maxHeight = input
			}
		}
	}

	heightList := []int{}
	minTime := 0
	for i:=minHeight ; i<=maxHeight ; i++ {
		time := 0
		totalBlock := B
		for j:=0 ; j<N ; j++ {
			for k:=0 ; k<M ; k++ {
				Blocks := i-field[j][k]
				totalBlock -= Blocks
				if Blocks < 0 {
					time += -2*Blocks
				}else {
					time += Blocks
				}
			}
		}
		if totalBlock >= 0{
			if minTime > time  || i == minHeight{
				heightList = []int{i}
				minTime = time
			}else if minTime == time {
				heightList = append(heightList, i)
			}
		}
	}
	sort.Ints(heightList)
	fmt.Fprintf(writer, "%d %d", minTime, heightList[len(heightList)-1])
}
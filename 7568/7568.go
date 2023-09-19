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

	var N int
	fmt.Fscan(reader, &N)

	volumeList := make([][]int, N)
	for i:=0 ; i<N ; i++ {
		var weight, height int
		fmt.Fscan(reader, &weight, &height)
		volumeList[i] = []int{weight, height}
	}

	for i:=0 ; i<N ; i++ {
		rank:=1
		for j:=0 ; j<N ; j++ {
			if i != j {
				isLight := volumeList[i][0] < volumeList[j][0]
				isShort := volumeList[i][1] < volumeList[j][1]
				if isLight && isShort {
					rank++
				}
			}
		}
		fmt.Fprintf(writer, "%d ", rank)
	}
}
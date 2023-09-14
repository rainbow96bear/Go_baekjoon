package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)

	position := make([][]float64, N)
	for i:=0 ; i<N ; i++ {
		var x, y float64
		fmt.Fscan(reader, &x, &y)
		position[i] = []float64{x,y}
	}
	answer := 0.0
	for i:=0 ; i<N-2 ; i++ {
		for j:=i+1 ; j<N-1 ; j++ {
			for k:=j+1 ; k<N ; k++ {
				if rst := calc(position[i], position[j], position[k]) ; answer < rst {
					answer = rst
				}
			}
		}
	}
	fmt.Fprint(writer, answer/2)
}

func calc(pos1, pos2, pos3 []float64) float64 {
	rst1 := pos1[0]*pos2[1] + pos2[0]*pos3[1] + pos3[0]*pos1[1]
	rst2 := pos1[0]*pos3[1] + pos2[0]*pos1[1] + pos3[0]*pos2[1]
	return math.Abs(rst1 - rst2)
}
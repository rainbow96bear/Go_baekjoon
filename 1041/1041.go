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
	fmt.Fscanf(reader, "%d\n", &N)

	var A, B, C, D, E, F int
	fmt.Fscanf(reader, "%d %d %d %d %d %d\n", &A, &B, &C, &D, &E, &F)

	result := CalcMinSum(N, A, B, C, D, E, F)
	fmt.Fprintln(writer, result)
}

func CalcMinSum(N, A, B, C, D, E, F int) (result int64){
	if N == 1 {
		return int64(A+B+C+D+E+F - max(A, B, C, D, E, F))
	}
	minThreeDisplay := min(A+B+C, A+C+E, A+D+E, A+B+D, F+B+C, F+C+E, F+D+E, F+B+D)
	// case ABC ACE ADE ABD
	// case FBC FCE FDE FBD
	
	minTwoDisplay := min(A+B, A+C, A+E, A+D, D+B, B+C, C+E, D+E, F+B, F+C, F+E, F+D)
	// case AB AC AE AD DB BC CE DE FB FC FE FD
	
	minOneDisplay := min(A, B, C, D, E, F)
	// case A B C D E F

	result = int64(minThreeDisplay)*4 + int64(minTwoDisplay)*int64((2*N-3)*4) + int64(minOneDisplay)*int64(5*((N-2)*(N-2))+4*(N-2))
	return result
}

func min(min int, nums ...int) (int) {
	for _, num := range nums {
		if min > num {
			min = num
		}
	} 
	return min
}

func max(max int, nums ...int) (int) {
	for _, num := range nums {
		if max < num {
			max = num
		}
	} 
	return max
}
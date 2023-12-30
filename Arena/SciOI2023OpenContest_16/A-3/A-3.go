package main

import (
	"bufio"
	"fmt"
	"os"
)

type Score struct {
	middleScore int
	lastScore int
}
func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	StudentScoreList := make([]Score, N+1)
	for i:=0 ; i<=N ; i++ {
		var middleScore, lastScore int
		fmt.Fscanf(reader, "%d %d\n", &middleScore, &lastScore)
		StudentScoreList[i] = Score{middleScore:middleScore,lastScore: lastScore }
	}

	// result := FindMinUnBalanceRate(N, StudentScoreList)
	// fmt.Fprintln(writer, result)
}

// func FindMinUnBalanceRate(N int, StudentScoreList []Score) (rate int) {
// 	Group := make([]Score, 3)
// 	for n:=1 ; n<=N ; n++ {
// 		middleScore := StudentScoreList[n].middleScore
// 		lastScore := StudentScoreList[n].lastScore
// 		for i:=0 ; i<3 ; i++ {
// 			if middleScore > Group[i].middleScore && lastScore > Group[i].lastScore {
				
// 			}
// 		}
// 	}
// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
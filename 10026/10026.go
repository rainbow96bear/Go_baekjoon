package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type position struct {
	x int
	y int
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)

	board := make([][]string, N)
	TwoColorBoard := make([][]string, N)
	
	for i:=0 ; i<N ; i++ {
		var input string
		fmt.Fscan(reader, &input)

		TwoColorBoard[i] = make([]string, N)
		board[i] = strings.Split(input, "")
		for j:=0 ; j<N ; j++ {
			if board[i][j] == "R" || board[i][j] == "G" {
				TwoColorBoard[i][j] = "R"
			}else {
				TwoColorBoard[i][j] = "B"
			}
		}
	}


	threeResult := Check(board, N)
	twoResult := Check(TwoColorBoard, N)

	fmt.Fprintln(writer, threeResult, twoResult)

}

func Check(board [][]string, N int) int{
	isVisit := make([][]bool, N)
	for i:=0 ; i<N ; i++ {
		isVisit[i] = make([]bool, N)
	}
	cnt := 0
	for i:=0 ; i<N ; i++ {
		for j:=0 ; j<N ; j++ {
			if isVisit[i][j] == false {
				cnt++
				queue := []position{{i, j}}
				move := [][]int{{0,1}, {0,-1}, {1,0}, {-1,0}}
				isVisit[i][j] = true

				for len(queue)>0 {
					now := queue[0]
					queue = queue[1:]
					for k:=0 ; k<4 ; k++ {
						X := now.x + move[k][0]
						Y := now.y + move[k][1]
						if 0<=X && X<N && 0<=Y && Y<N {
							if isVisit[X][Y] == false && board[X][Y] == board[i][j]{
								queue = append(queue, position{x:X, y:Y})
								isVisit[X][Y] = true
							}
						} 
					}
				}
			}
		}
	}
	return cnt
}
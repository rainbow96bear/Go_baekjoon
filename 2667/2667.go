package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	for i:=0 ; i<N ; i++ {
		var input string
		fmt.Fscan(reader, &input)
		nums := strings.Split(input, "")
		board[i] = nums
	}

	totalGroupNum := 0
	sizeList := []int{}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	
	for i:=0 ; i<N ; i++ {
		for j:=0 ; j<N ; j++ {
			if board[i][j] == "1" {
				totalGroupNum++
				cnt := 0
				start := position{x:i, y:j}
				queue := []position{start}
				board[i][j] = "0"
				for len(queue) > 0 {
					nowPosition := queue[0]
					queue = queue[1:]
					cnt++
					
					
					for i:=0 ; i<4 ; i++ {
						nextX := nowPosition.x+dx[i]
						nextY := nowPosition.y+dy[i]
						xIsOk := 0 <= nextX && nextX <N
						yIsOk := 0 <= nextY && nextY <N
						if xIsOk && yIsOk && board[nextX][nextY] == "1" {
							queue = append(queue, position{x:nextX, y:nextY})
							board[nextX][nextY] = "0" 
						}
					}
				}
				sizeList = append(sizeList, cnt)
			}
		}
	}

	sort.Ints(sizeList)
	fmt.Fprintln(writer, totalGroupNum)
	for i:=0 ; i<len(sizeList) ; i++ {
		fmt.Fprintln(writer, sizeList[i])
	}
}
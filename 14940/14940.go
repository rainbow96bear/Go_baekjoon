package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct{
	x int
	y int
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	board := make([][]int, n)
	isVisit := make([][]bool, n)
	answer := make([][]int, n)

	queue := []position{}
	for i:=0 ; i<n ; i++ {
		board[i] = make([]int, m)
		answer[i] = make([]int, m)
		isVisit[i] = make([]bool, m)
		for j:=0 ; j<m ; j++ {
			fmt.Fscan(reader, &board[i][j])
			if board[i][j] == 2 {
				isVisit[i][j] = true
				queue = append(queue, position{x:i, y:j})
			}
		}
	}
	cnt := 0
	for len(queue) > 0 {
		newQueue := []position{}
		cnt++
		for len(queue) > 0 {
			nowPosition := queue[0]
			queue = queue[1:]
			dx := []int{-1, 1, 0, 0}
			dy := []int{0, 0, -1, 1}
			for i:=0 ; i<4 ; i++{
				x := nowPosition.x+dx[i]
				y := nowPosition.y+dy[i]
				if 0<=x && x<n && 0<=y && y<m && isVisit[x][y] == false {
					isVisit[x][y] = true
					if board[x][y] == 1 {
						newQueue = append(newQueue, position{x:x, y:y})
						answer[x][y] = cnt
					}
				}
			}
		}
		queue = newQueue
	}
	for i:=0 ; i<n ; i++ {
		for j:=0 ; j<m ; j++ {
			if board[i][j] == 1 && isVisit[i][j] == false {
				fmt.Fprintf(writer, "%d ", -1)
			}else {
				fmt.Fprintf(writer, "%d ", answer[i][j])
			}
		}
		fmt.Fprintln(writer)
	}
}
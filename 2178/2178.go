package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type position struct{
	x int
	y int
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, M int
	fmt.Fscan(reader, &N, &M)

	board := make([][]int, N)
	for i:=0 ; i<N ; i++ {
		var row string
		fmt.Fscan(reader, &row)
		input := strings.Split(row, "")
		for j:=0 ; j<M ; j++ {
			num, _ := strconv.Atoi(input[j])
			board[i] = append(board[i], num)
		}
	}

	start := position{x:0, y:0}
	queue := []position{start}
	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:]

		if now.x == N-1 && now.y == M-1 {
			fmt.Fprintln(writer, board[now.x][now.y])
			return
		}

		dx := []int{-1, 1, 0, 0}
		dy := []int{0, 0, -1, 1}

		for i:=0 ; i<4 ; i++ {
			X := now.x+dx[i]
			Y := now.y+dy[i]
			if 0 <= X && X < N && 0 <= Y && Y < M && board[X][Y] == 1 {
				board[X][Y] = board[now.x][now.y]+1
				queue = append(queue, position{x:X, y:Y})
			}
		}
	}
}
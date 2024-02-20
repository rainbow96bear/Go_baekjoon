package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)

	board := make([][]int, N)
	max := 0
	for i := 0; i < N; i++ {
		board[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Fscanf(reader, "%d", &board[i][j])
			if max < board[i][j] {
				max = board[i][j]
			}
		}
		fmt.Fscanf(reader, "\n")
	}

	answer := 1
	for i := 1; i < max; i++ {
		result := findArea(i, N, board)
		if answer < result {
			answer = result
		}
	}

	fmt.Fprintln(writer, answer)
}

func findArea(height, N int, board [][]int) int {
	count := 0

	visited := make([][]bool, N)
	for i := range visited {
		visited[i] = make([]bool, N)
	}

	queue := make([][]int, 0)

	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			if !visited[x][y] && board[x][y] > height {
				count++
				visited[x][y] = true
				queue = append(queue, []int{x, y})
			}

			for len(queue) > 0 {
				now := queue[0]
				queue = queue[1:]
				dx := []int{0, 0, 1, -1}
				dy := []int{1, -1, 0, 0}

				for i := 0; i < 4; i++ {
					nx, ny := now[0]+dx[i], now[1]+dy[i]
					if nx >= 0 && nx < N && ny >= 0 && ny < N && !visited[nx][ny] && board[nx][ny] > height {
						visited[nx][ny] = true
						queue = append(queue, []int{nx, ny})
					}
				}
			}
		}
	}

	return count
}

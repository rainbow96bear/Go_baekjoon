package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r, c, ans int
	board       [][]byte
	isUsed   []bool
	dx        = []int{1, 0, -1, 0}
	dy        = []int{0, 1, 0, -1}
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d %d", &r, &c)

	board = make([][]byte, r)
	isUsed = make([]bool, 26)

	for i := 0; i < r; i++ {
		scanner.Scan()
		line := scanner.Text()
		board[i] = []byte(line)
	}

	isUsed[int(board[0][0]-'A')] = true
	DFS(0, 0, 1)

	fmt.Fprintln(writer, ans)
}

func DFS(x, y, dist int) {
	ans = max(ans, dist)

	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx < 0 || nx >= r || ny < 0 || ny >= c {
			continue
		}

		nextAlpha := int(board[nx][ny] - 'A')
		if !isUsed[nextAlpha] {
			isUsed[nextAlpha] = true
			DFS(nx, ny, dist+1)
			isUsed[nextAlpha] = false
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

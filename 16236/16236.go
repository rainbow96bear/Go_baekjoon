package main

import (
	"container/heap"
	"fmt"
)

type Point struct {
	row, col, distance int
}

type PriorityQueue []Point

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].distance < pq[j].distance }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Point))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func bfs(graph [][]int, start Point, size int) Point {
	n := len(graph)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[start.row][start.col] = true
	queue := PriorityQueue{{start.row, start.col, 0}}
	possibleFishes := make([]Point, 0)

	directions := [4][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

	for len(queue) > 0 {
		current := heap.Pop(&queue).(Point)

		for _, dir := range directions {
			nr, nc := current.row+dir[0], current.col+dir[1]

			if nr >= 0 && nr < n && nc >= 0 && nc < n && !visited[nr][nc] {
				visited[nr][nc] = true

				if graph[nr][nc] == 0 || graph[nr][nc] == size {
					heap.Push(&queue, Point{nr, nc, current.distance + 1})
				} else if graph[nr][nc] > 0 && graph[nr][nc] < size {
					possibleFishes = append(possibleFishes, Point{nr, nc, current.distance + 1})
				}
			}
		}
	}

	if len(possibleFishes) == 0 {
		return Point{-1, -1, -1}
	}

	// 물고기들을 거리, 행, 열 순서로 정렬
	for i := range possibleFishes {
		for j := i + 1; j < len(possibleFishes); j++ {
			if possibleFishes[i].distance > possibleFishes[j].distance ||
				(possibleFishes[i].distance == possibleFishes[j].distance &&
					(possibleFishes[i].row > possibleFishes[j].row ||
						(possibleFishes[i].row == possibleFishes[j].row && possibleFishes[i].col > possibleFishes[j].col))) {
				possibleFishes[i], possibleFishes[j] = possibleFishes[j], possibleFishes[i]
			}
		}
	}

	return possibleFishes[0]
}

func solution(n int, graph [][]int) int {
	size := 2
	exp := 0
	count := 0
	var sharkPos Point

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == 9 {
				sharkPos = Point{i, j, 0}
				graph[i][j] = 0
			}
		}
	}

	for {
		result := bfs(graph, sharkPos, size)

		if result.distance == -1 {
			break
		}

		sharkPos = Point{result.row, result.col, 0}
		graph[result.row][result.col] = 0
		count++
		exp += result.distance

		if count == size {
			size++
			count = 0
		}
	}

	return exp
}

func main() {
	var n int
	fmt.Scan(&n)

	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&graph[i][j])
		}
	}

	answer := solution(n, graph)
	fmt.Println(answer)
}

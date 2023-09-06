package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)
	Point := make([][]int, N)
	for i:=0 ; i<N ; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		Point[i]=[]int{x,y}
	}
	sort.Slice(Point,func(i, j int) bool {
		if Point[i][0] == Point[j][0]{
			return Point[i][1] > Point[j][1]
		}
		return Point[i][0] > Point[j][0]
	})
	for _, v := range Point{
		fmt.Fprintln(writer, v[0], v[1])
	}
}
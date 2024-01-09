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

	var W, H, f, c, x1, y1, x2, y2 int
	fmt.Fscanf(reader, "%d %d %d %d %d %d %d %d\n", &W, &H, &f, &c, &x1, &y1, &x2, &y2)

	result := CalcColorArea(W, H, f, c, x1, y1, x2, y2)
	fmt.Fprintln(writer, result)
}

func CalcColorArea(W, H, f, c, x1, y1, x2, y2 int) int64 {

	leftLength := f
	rightLength := W-f
	leftEmptyLength := 0
	if min(leftLength, x2)-x1 > 0 {
		leftEmptyLength = min(leftLength, x2)-x1
	}
	rightEmptyLength := 0
	if min(rightLength, x2)-x1 > 0 {
		rightEmptyLength = min(rightLength, x2)-x1
	}
	emptyArea := int64(leftEmptyLength+rightEmptyLength)*int64(min(H/(c+1),y2)-y1)

	return int64(W)*int64(H) - emptyArea*int64(c+1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	N, _ := strconv.Atoi(Input(scanner))

	cnt := Process(N)

	fmt.Fprintln(writer, cnt)
}
func Process(N int) int {
	checkCol := make(map[int]bool)
	checkSum := make(map[int]bool)
	checkSub := make(map[int]bool)

	result := Find(0, N, checkCol, checkSum, checkSub)

	return result
}

func Find(row, N int, checkCol, checkSum, checkSub map[int]bool) int {
	if row == N {
		return 1
	}
	cnt := 0
	for i:=0; i<N ; i++ {
		if checkCol[i] == false && checkSum[row+i] == false && checkSub[row-i] == false {
			checkCol[i] = true 
			checkSum[row+i] = true
			checkSub[row-i] = true
			cnt += Find(row+1, N, checkCol, checkSum, checkSub)
			checkCol[i] = false 
			checkSum[row+i] = false
			checkSub[row-i] = false
		}
	}
	return cnt
}

func Input(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}
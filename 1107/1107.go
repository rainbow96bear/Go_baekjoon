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

	var N int
	fmt.Fscan(reader, &N)

	var M int
	fmt.Fscan(reader, &M)

	isBreakButtons := make([]bool, 10)
	for i:=0 ; i<M ; i++ {
		var input int
		fmt.Fscan(reader, &input)
		isBreakButtons[input] = true
	}


	answer := N - 100
	if answer < 0 {
		answer = -answer
	}
	for i:=0 ; i<=999999 ; i++ {
		isPossible := true
		num := i
		cnt := 0
		if num == 0 {
			if isBreakButtons[0] {
				isPossible = false
			}
			cnt++
		}
		for num > 0 {
			if isBreakButtons[num%10] == true {
				isPossible = false
				break
			}
			num /= 10
			cnt++
		}
		if isPossible == true {
			gap := N-i
			if gap < 0 {
				gap = -gap
			}
			if answer > gap + cnt {
				answer = gap + cnt
			}
		}
	}
	fmt.Fprintln(writer, answer)
}
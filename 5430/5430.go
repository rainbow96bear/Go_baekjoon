package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var T int
	fmt.Fscan(reader, &T)

	for i:=0 ; i<T ; i++ {
		var p string
		fmt.Fscan(reader, &p)

		var n int
		fmt.Fscan(reader, &n)

		var inputArray string
		fmt.Fscan(reader, &inputArray)
		inputArray = inputArray[1:len(inputArray)-1]

		numArray := strings.Split(inputArray, ",")

		commandList := strings.Split(p, "")

		direction := true
		left := 0
		right := 0
		for j:=0 ; j<len(commandList) ; j++ {
			if commandList[j] == "R" {
				direction = !direction
			}else {
				if direction {
					left++
				}else {
					right++
				}
			}
		}
		if left+right > n {
			fmt.Fprintln(writer, "error")
			continue
		}
		result := numArray[left:len(numArray)-right]

		answer := ""

		if direction {
			answer = "[" + strings.Join(result, ",") + "]"
		}else {
			for j, k := 0, len(result)-1; j < k; j, k = j+1, k-1 {
				result[j], result[k] = result[k], result[j]
			}
			answer = "[" + strings.Join(result, ",") + "]"
		}

		fmt.Fprintln(writer, answer)
	}
}
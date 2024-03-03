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

	var input string
	fmt.Fscanf(reader, "%s\n", &input)
	list := strings.Split(input, "")
	standard := list[0]
	isSkip := false
	count := 0
	for i:=0 ; i<len(list) ; i++ {
		if standard != list[i] && !isSkip {
			count++
			isSkip = true
		}else {
			if standard == list[i] {
				isSkip = false
			}
	}
	}
	answer := min(count, len(list)-count)
	fmt.Fprintln(writer, answer)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
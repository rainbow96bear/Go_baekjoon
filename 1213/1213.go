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

	var name string
	fmt.Fscan(reader, &name)

	spellingNum := make(map[string]int)
	for _, value := range name {
		spellingNum[string(value)]++
	}

	answerLeft := ""
	answerRight := ""
	centerSpelling := "NOT"
	for spelling:='Z' ; spelling>='A' ; spelling-- {
		for spellingNum[string(spelling)] >= 2 {
			answerLeft = fmt.Sprintf("%s%s", string(spelling), answerLeft)
			answerRight = fmt.Sprintf("%s%s", answerRight, string(spelling))
			spellingNum[string(spelling)] -= 2
		}
		if spellingNum[string(spelling)] == 1 {
			if centerSpelling == "NOT" {
				centerSpelling = string(spelling)
			}else {
				fmt.Fprint(writer, "I'm Sorry Hansoo")
				return
			}
		}
	}
	if centerSpelling != "NOT" {
		fmt.Fprintf(writer,"%s%s%s", answerLeft, centerSpelling, answerRight)
		return
	}
	fmt.Fprintf(writer,"%s%s", answerLeft, answerRight)
}
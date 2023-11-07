package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	writer := bufio.NewWriter(os.Stdout)
	
	defer writer.Flush()
	
	// 오답
	// scanner := bufio.NewScanner(os.Stdin)
	// text := Input(scanner)

	// bomb := Input(scanner)
	// 오답

	// 정답
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	bomb, _ := reader.ReadString('\n')
	bomb = strings.TrimSpace(bomb)
	// 정답

	result := make([]rune, 0)
	for _, v := range text {
		result = append(result, v)
		if len(result)-len(bomb) >= 0 {
            if string(result[len(result)-len(bomb):]) == bomb {
                result = result[:len(result)-len(bomb)]
            }
		}
	}
	if len(result) == 0 {
		fmt.Fprint(writer, "FRULA")
	}else {
		fmt.Fprint(writer,string(result))
	}
}

func Input(scanner *bufio.Scanner) string {
	scanner.Scan()
	input := scanner.Text()
	input = strings.TrimRight(input, "\n")
	input = strings.TrimSpace(input)

	return input
}
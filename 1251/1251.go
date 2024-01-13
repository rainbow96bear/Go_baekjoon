package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	
	var word string
	fmt.Fscanf(reader, "%s\n", &word)

	result := divideWord(word)

	fmt.Fprintln(writer, result)
}


func divideWord(word string) string {
	result := make([]string, 0)
	for i := 1; i < len(word)-1; i++ {
		for j := i + 1; j < len(word); j++ {
			dividedWord := reverse(word[:i]) + reverse(word[i:j]) + reverse(word[j:])
			result = append(result, dividedWord)
		}
	}

	sort.Strings(result)
	return result[0]
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var T int
	fmt.Fscanf(reader, "%d\n", &T)

	for t:=0 ; t<T ; t++ {
		var input string
		fmt.Fscanf(reader, "%s\n", &input)
		result := checkPattern(input)
		fmt.Fprintln(writer, result)
	}
}

func checkPattern(input string) string {
	pattern := "^(100+1+|01)+$"
	match, err := regexp.MatchString(pattern, input)

	if err != nil {
		fmt.Println("Error while matching regex:", err)
		return "N)"
	}

	if match {
		return "YES"
	}
	return "NO"
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	sik := strings.Split(Input(scanner), "")
	sik = RemoveBracket(sik)
	sik = Process(sik, "*", "/")
	sik = Process(sik, "+", "-")

	fmt.Fprintln(writer, strings.Join(sik, ""))
}
func RemoveBracket(sik []string) []string{
	stack := []string{}

	for i:=0 ; i<len(sik) ; i++ {
		stack = append(stack, sik[i])
		if stack[len(stack)-1] == ")" {
			for j:=len(stack)-1 ; j>=0 ; j-- {
				if stack[j] == "(" {
					tempSik := []string{}
					tempSik = append(tempSik, stack[j+1:len(stack)-1]...)
					tempSik = Process(tempSik, "*", "/")
					tempSik = Process(tempSik, "+", "-")
					stack = stack[:j]
					stack = append(stack, strings.Join(tempSik, ""))
					break
				}
			}
		}
	}
	return stack
}

func Process(sik []string, calc1, calc2 string) []string {
	stack := []string{}
	for i:=0 ; i<len(sik) ; i++ {
		if sik[i] ==calc1 || sik[i] == calc2 {
			length:=len(stack)
			newValue := stack[length-1]+sik[i+1]+sik[i]
			stack[length-1] = newValue
			i++
		}else {
			stack = append(stack, sik[i])
		}
	}
	return stack
}

func Input(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}
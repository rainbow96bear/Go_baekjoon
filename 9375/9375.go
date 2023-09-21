package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var TestCase int
	fmt.Fscan(reader, &TestCase)

	for i := 0; i < TestCase; i++ {
		var n int
		fmt.Fscan(reader, &n)
		if n == 0 {
			fmt.Fprintln(writer, 0)
			continue // 다음 테스트 케이스로 이동
		}
		list := make(map[string]int)
		for j := 0; j < n; j++ {
			var wear, kind string
			fmt.Fscan(reader, &wear, &kind)
			list[kind]++
		}
		answer := 1
		for _, count := range list {
			answer *= (count + 1)
		}
		answer-- // 모든 의상을 입지 않은 경우를 빼줍니다.
		fmt.Fprintln(writer, answer)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)

    defer writer.Flush()

    var N int
    fmt.Fscan(reader, &N)

    strGroup := make([]string, N)
    for i := 0; i < N; i++ {
        fmt.Fscan(reader, &strGroup[i])
    }

    sort.Slice(strGroup, func(i, j int) bool {
        return len(strGroup[i]) < len(strGroup[j])
    })

    answer := N

    for i := 0; i < N; i++ {
        for j := i + 1; j < N; j++ {
            if strings.HasPrefix(strGroup[j], strGroup[i]) {
                answer--
                break
            }
        }
    }

    fmt.Fprint(writer, answer)
}
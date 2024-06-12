package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func encode(s string) string {
	n := "Tinkoff"
	tcount := 0
	icount := 0
	ncount := 0
	kcount := 0
	ocount := 0
	fcount := 0
	for _, v := range n {
		for _, i := range s {
			if i == v {
				switch string(i) {
				case "T":
					tcount++
				case "i":
					icount++
				case "n":
					ncount++
				case "k":
					kcount++
				case "o":
					ocount++
				case "f":
					fcount++

				}
			}
		}
	}
	minMatch := min(tcount, icount, ncount, kcount, ocount, fcount/2)
	result, _ := fmt.Printf("Слово Tinkoff встречается %d", minMatch)
	return strconv.Itoa(result)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')

	encode(s)
}

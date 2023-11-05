package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	str := ""
	in := bufio.NewReader(os.Stdin)
	str, _ = in.ReadString('\n')
	splitedStr := strings.Split(str, " ")
	last := splitedStr[len(splitedStr)-1]
	splitedStr[len(splitedStr)-1] = splitedStr[len(splitedStr)-1][0 : len(last)-2] // "someNumber\n" -> "someNumber"
	var setStr = make(map[string]string)
	for i := 0; i < len(splitedStr); i++ {
		setStr[splitedStr[i]] = ""
	}
	fmt.Println(len(setStr))
}

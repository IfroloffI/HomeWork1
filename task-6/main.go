package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func clearStr(str string) []string {
	splitedStr := strings.Split(str, " ")
	last := splitedStr[len(splitedStr)-1]
	splitedStr[len(splitedStr)-1] = splitedStr[len(splitedStr)-1][0 : len(last)-2] // "someNumber\n" -> "someNumber"
	return splitedStr
}

func main() {
	str := ""
	in := bufio.NewReader(os.Stdin)
	str, _ = in.ReadString('\n')
	clearStrVar := clearStr(str)
	var setStr = make(map[string]struct{})
	for i := 0; i < len(clearStrVar); i++ {
		setStr[clearStrVar[i]] = struct{}{}
	}
	fmt.Println(len(setStr))
}

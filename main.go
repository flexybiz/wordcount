package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	text := os.Args[1]
	count := len(strings.Fields(text))
	fmt.Println(count)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	max := point{4, 4}
	min := point{0, 0}
	robot := NewRobot(table{min, max})

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ", robot)
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
	}
}

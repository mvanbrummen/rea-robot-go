package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var interactive = flag.Bool("interactive", false, "Interactive mode")

func init() {
	flag.BoolVar(interactive, "i", false, "Interactive mode")
}

func main() {
	flag.Parse()
	max := point{4, 4}
	min := point{0, 0}
	robot := NewRobot(table{min, max})

	if *interactive {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ")
			text, _ := reader.ReadString('\n')
			parseCommand(text, robot)
			os.Exit(1)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			parseCommand(scanner.Text(), robot)
		}
	}
}

func parseCommand(command string, r *robot) {
	command = strings.Replace(command, "\n", "", -1)
	switch command {
	case "MOVE":
		r.Move()
	case "LEFT":
		r.RotateLeft()
	case "RIGHT":
		r.RotateRight()
	case "REPORT":
		r.Report()
	default:
		if strings.HasPrefix(command, "PLACE ") {
			var (
				x, y   int
				facing string
			)
			fmt.Sscanf(command, "PLACE %d,%d,%s", &x, &y, &facing)
			dir := Parse(facing)
			if dir != -1 {
				r.Place(x, y, dir)
			}
		}
	}
}

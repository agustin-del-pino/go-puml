package main

import (
	"github.com/agustin-del-pino/gopuml/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		println(err)
	}
}

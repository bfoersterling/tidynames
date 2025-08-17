package main

import (
	"fmt"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cli_args := get_cli_args()

	err := cli_args.evalute()

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}

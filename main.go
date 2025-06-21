package main

import (
	"fmt"
	"os"
)

func main() {
	cli_args := get_cli_args()

	err := cli_args.evalute()

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}

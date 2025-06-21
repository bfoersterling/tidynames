package main

import (
	"flag"
	"fmt"
	"os"
)

type cli_args struct {
	dry_run   bool
	recursive string
}

func get_cli_args() (args cli_args) {
	args = cli_args{}

	flag.StringVar(&args.recursive, "r", "", "apply tidynames recursively in the given dir")
	flag.StringVar(&args.recursive, "recursive", "", "apply tidynames recursively in the given dir")
	flag.BoolVar(&args.dry_run, "n", false, "just print changes - do not rename")
	flag.BoolVar(&args.dry_run, "noop", false, "just print changes - do not rename")
	flag.Parse()

	return
}

func (args cli_args) evalute() (err error) {
	rc := replace_config{}

	if len(args.recursive) > 0 {
		err = fmt.Errorf("%s - Recursive tidying is not implemented yet.\n", os.Args[0])
		return
	}

	if flag.NArg() == 0 {
		err = fmt.Errorf("%s - Missing arguments.\n", os.Args[0])
		return
	}

	if args.dry_run {
		fmt.Printf("[dry run]\n")
	}

	rc.tidy_entries(args, flag.Args(), os.Stdout)

	return
}

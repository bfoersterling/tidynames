package main

import (
	"flag"
	"fmt"
	"os"
)

type cli_args struct {
	dry_run   bool
	recursive string
	version   bool
}

func get_cli_args() (args cli_args) {
	args = cli_args{}

	flag.StringVar(&args.recursive, "r", "", "apply tidynames recursively in the given dir")
	flag.StringVar(&args.recursive, "recursive", "", "apply tidynames recursively in the given dir")
	flag.BoolVar(&args.dry_run, "n", false, "just print changes - do not rename")
	flag.BoolVar(&args.dry_run, "noop", false, "just print changes - do not rename")
	flag.BoolVar(&args.version, "v", false, "print the version of tidynames")
	flag.BoolVar(&args.version, "version", false, "print the version of tidynames")
	flag.Parse()

	return
}

func (args cli_args) evalute() (err error) {
	rc := replace_config{
		whitespace: '_',
	}

	if args.version {
		fmt.Printf("%s %s, commit: %s, build at: %s.\n", os.Args[0], version, commit, date)
		os.Exit(0)
	}

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

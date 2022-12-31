package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"

	"github.com/itchyny/maketen-go"
)

const name = "maketen"

const version = "0.1.0"

var revision = "HEAD"

func main() {
	os.Exit((&cli{
		outStream: os.Stdout,
		errStream: os.Stderr,
	}).run(os.Args[1:]))
}

const (
	exitCodeOK = iota
	exitCodeErr
)

type cli struct {
	outStream io.Writer
	errStream io.Writer
}

func (cli *cli) run(args []string) int {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(cli.errStream)
	fs.Usage = func() {
		fs.SetOutput(cli.outStream)
		fmt.Fprintf(cli.outStream, `%[1]s - create 10 from numbers

Version: %s (rev: %s/%s)

Synopsis:
    %% %[1]s 1 2 3 4

Options:
`, name, version, revision, runtime.Version())
		fs.PrintDefaults()
	}
	var showVersion bool
	fs.BoolVar(&showVersion, "v", false, "print version")
	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return exitCodeOK
		}
		return exitCodeErr
	}
	if showVersion {
		fmt.Fprintf(cli.outStream, "%s %s (rev: %s/%s)\n", name, version, revision, runtime.Version())
		return exitCodeOK
	}
	args = fs.Args()
	var err error
	if len(args) < 1 {
		fmt.Fprintf(cli.errStream, "%s: specify numbers\n", name)
		return exitCodeErr
	}
	ns := make([]*maketen.Num, len(args))
	var n *maketen.Num
	for i, arg := range args {
		if n, err = parseInt(arg); err != nil {
			fmt.Fprintf(cli.errStream, "%s: %s\n", name, err)
			return exitCodeErr
		}
		ns[i] = n
	}
	solutions := maketen.Solve(ns...)
	if len(solutions) == 0 {
		fmt.Fprintf(cli.errStream, "%s: no answer\n", name)
		return exitCodeErr
	}
	for _, e := range solutions {
		fmt.Fprintf(cli.outStream, "%+v\n", e)
	}
	return exitCodeOK
}

func parseInt(str string) (*maketen.Num, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %q", str)
	}
	return maketen.NewInt(i), nil
}

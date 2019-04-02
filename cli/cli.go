package cli

import (
	"flag"
	"fmt"
	"io"
	"runtime"
	"strconv"

	"github.com/itchyny/maketen-go"
)

const name = "maketen"

const version = "0.0.0"

var revision = "HEAD"

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
		fmt.Fprintf(cli.outStream, `%[1]s - create 10 from four numbers

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
	var x, y, z, w *maketen.Num
	var err error
	if len(args) < 4 {
		fmt.Fprintln(cli.errStream, "too few arguments")
		return exitCodeErr
	} else if len(args) > 4 {
		fmt.Fprintln(cli.errStream, "too many arguments")
		return exitCodeErr
	}
	if x, err = parseInt(args[0]); err != nil {
		fmt.Fprintf(cli.errStream, "%s\n", err)
		return exitCodeErr
	}
	if y, err = parseInt(args[1]); err != nil {
		fmt.Fprintf(cli.errStream, "%s\n", err)
		return exitCodeErr
	}
	if z, err = parseInt(args[2]); err != nil {
		fmt.Fprintf(cli.errStream, "%s\n", err)
		return exitCodeErr
	}
	if w, err = parseInt(args[3]); err != nil {
		fmt.Fprintf(cli.errStream, "%s\n", err)
		return exitCodeErr
	}
	solutions := maketen.Solve(x, y, z, w)
	if len(solutions) == 0 {
		fmt.Fprintln(cli.errStream, "no answer")
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

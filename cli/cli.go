package cli

import (
	"fmt"
	"io"
	"strconv"

	"github.com/itchyny/maketen-go"
)

const (
	exitCodeOK = iota
	exitCodeErr
)

type cli struct {
	outStream io.Writer
	errStream io.Writer
}

func (cli *cli) run(args []string) int {
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
	for _, e := range maketen.Solve(x, y, z, w) {
		fmt.Fprintf(cli.outStream, "%+v\n", e)
	}
	return exitCodeOK
}

func parseInt(str string) (*maketen.Num, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return nil, err
	}
	return maketen.NewInt(i), nil
}

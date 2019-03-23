package cli

import "os"

// Run maketen.
func Run() int {
	return (&cli{
		outStream: os.Stdout,
		errStream: os.Stderr,
	}).run(os.Args[1:])
}

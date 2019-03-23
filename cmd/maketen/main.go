package main

import (
	"fmt"

	"github.com/itchyny/maketen-go"
)

func main() {
	for _, e := range maketen.Solve(
		maketen.NewInt(1),
		maketen.NewInt(2),
		maketen.NewInt(3),
		maketen.NewInt(4),
	) {
		fmt.Printf("%+v\n", e)
	}
}

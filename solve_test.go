package maketen

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	var results strings.Builder
	var cnt int
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			for z := 0; z < 10; z++ {
				for w := 0; w < 10; w++ {
					got := Solve(NewInt(x), NewInt(y), NewInt(z), NewInt(w))
					for _, e := range got {
						cnt++
						results.WriteString(e.String())
						results.WriteRune('\n')
					}
				}
			}
		}
	}
	if expected := 23741; cnt != expected {
		t.Fatalf("expected %d solutions but got: %d", expected, cnt)
	}

	cmd := exec.Command("bc", "-l")
	cmd.Stdin = strings.NewReader(results.String())
	out := new(bytes.Buffer)
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	scanner := bufio.NewScanner(out)
	var i int
	for scanner.Scan() {
		i++
		line := scanner.Text()
		if line == "10" || strings.HasPrefix(line, "10.000000") ||
			strings.HasPrefix(line, "9.9999999") {
			continue
		}
		t.Fatalf("got %s: %+v", line, strings.Split(results.String(), "\n")[i-1])
	}
	if scanner.Err() != nil {
		t.Fatal(scanner.Err())
	}
}

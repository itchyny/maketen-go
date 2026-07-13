package maketen

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestSolve0(t *testing.T) {
	var count int
	for range Solve() {
		count++
	}
	if expected := 0; count != expected {
		t.Fatalf("expected %d solutions but got: %d", expected, count)
	}
}

func TestSolve1(t *testing.T) {
	var results []string
	var count int
	for x := range 20 {
		for e := range Solve(NewNum(x)) {
			count++
			results = append(results, e.String())
		}
	}
	if expected := 1; count != expected {
		t.Fatalf("expected %d solutions but got: %d", expected, count)
	}
	testResults(t, results)
}

func TestSolve2(t *testing.T) {
	var results []string
	var count int
	for x := range 10 {
		for y := range 10 {
			for e := range Solve(NewNum(x), NewNum(y)) {
				count++
				results = append(results, e.String())
			}
		}
	}
	if expected := 11; count != expected {
		t.Fatalf("expected %d solutions but got: %d", expected, count)
	}
	testResults(t, results)
}

func TestSolve3(t *testing.T) {
	var results []string
	var count int
	for x := range 10 {
		for y := range 10 {
			for z := range 10 {
				for e := range Solve(NewNum(x), NewNum(y), NewNum(z)) {
					count++
					results = append(results, e.String())
				}
			}
		}
	}
	if expected := 412; count != expected {
		t.Fatalf("expected %d solutions but got: %d", expected, count)
	}
	testResults(t, results)
}

func TestSolve4(t *testing.T) {
	var results []string
	var count int
	for x := range 10 {
		for y := range 10 {
			for z := range 10 {
				for w := range 10 {
					for e := range Solve(NewNum(x), NewNum(y), NewNum(z), NewNum(w)) {
						count++
						results = append(results, e.String())
					}
				}
			}
		}
	}
	if expected := 23741; count != expected {
		t.Fatalf("expected %d solutions but got: %d", expected, count)
	}
	testResults(t, results)
}

func testResults(t *testing.T, results []string) {
	cmd := exec.Command("bc", "-l")
	cmd.Stdin = strings.NewReader(strings.Join(results, "\n"))
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
		t.Errorf("got %s: %s", line, results[i-1])
	}
	if scanner.Err() != nil {
		t.Fatal(scanner.Err())
	}
}

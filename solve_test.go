package maketen

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestSolve0(t *testing.T) {
	got := Solve()
	if expected := 0; len(got) != expected {
		t.Fatalf("expected %d solutions but got: %d", expected, len(got))
	}
}

func TestSolve1(t *testing.T) {
	var results []string
	var count int
	for x := 0; x <= 10; x++ {
		got := Solve(NewInt(x))
		for _, e := range got {
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
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			got := Solve(NewInt(x), NewInt(y))
			for _, e := range got {
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
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			for z := 0; z < 10; z++ {
				got := Solve(NewInt(x), NewInt(y), NewInt(z))
				for _, e := range got {
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
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			for z := 0; z < 10; z++ {
				for w := 0; w < 10; w++ {
					got := Solve(NewInt(x), NewInt(y), NewInt(z), NewInt(w))
					for _, e := range got {
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

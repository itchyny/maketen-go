package cli

import (
	"strings"
	"testing"
)

func TestCliRun(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		expected string
		err      string
	}{
		{
			name: "solve 0",
			args: []string{"1", "2", "3", "4"},
			expected: `1 + 2 + 3 + 4
1 * (2 * 3 + 4)
1 * 2 * 3 + 4
`,
		},
		{
			name: "solve 1",
			args: []string{"1", "3", "5", "8"},
			expected: `1 - 3 * (5 - 8)
`,
		},
		{
			name: "solve 2",
			args: []string{"8", "1", "1", "5"},
			expected: `8 / (1 - 1 / 5)
`,
		},
		{
			name: "solve 3",
			args: []string{"1", "1", "9", "9"},
			expected: `(1 + 1 / 9) * 9
`,
		},
		{
			name: "solve 4",
			args: []string{"2", "5"},
			expected: `2 * 5
`,
		},
		{
			name: "solve 5",
			args: []string{"5", "4", "8"},
			expected: `5 / 4 * 8
5 / (4 / 8)
`,
		},
		{
			name: "solve 6",
			args: []string{"4", "5", "6", "7", "8"},
			expected: `(4 * 5 - 6) / 7 + 8
`,
		},
		{
			name: "no answer",
			args: []string{"1", "1", "7", "9"},
			err: `no answer
`,
		},
		{
			name: "specify numbers",
			args: []string{},
			err: `specify numbers
`,
		},
		{
			name: "parse error",
			args: []string{"1", "1", "1", "x"},
			err: `failed to parse "x"
`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var outStream, errStream strings.Builder
			cli := cli{
				outStream: &outStream,
				errStream: &errStream,
			}
			code := cli.run(tc.args)
			if tc.err == "" {
				if code != exitCodeOK {
					t.Errorf("code should be %d but got %d", exitCodeOK, code)
				}
				if got := outStream.String(); got != tc.expected {
					t.Errorf("output should be %q but got %q", tc.expected, got)
				}
			} else {
				if code != exitCodeErr {
					t.Errorf("code should be %d but got %d", exitCodeErr, code)
				}
				if got := errStream.String(); !strings.Contains(got, tc.err) {
					t.Errorf("error output should contain %q but got %q", tc.err, got)
				}
			}
		})
	}
}

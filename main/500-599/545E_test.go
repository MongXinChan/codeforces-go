// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/545/E
// https://codeforces.com/problemset/status/545/problem/E?friends=on
func Test_cf545E(t *testing.T) {
	testCases := [][2]string{
		{
			`3 3
1 2 1
2 3 1
1 3 2
3`,
			`2
1 2`,
		},
		{
			`4 4
1 2 1
2 3 1
3 4 1
4 1 2
4`,
			`4
2 3 4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf545E)
}

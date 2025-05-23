// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/632/D
// https://codeforces.com/problemset/status/632/problem/D?friends=on
func Test_cf632D(t *testing.T) {
	testCases := [][2]string{
		{
			`7 8
6 2 9 2 7 2 3`,
			`6 5
1 2 4 6 7`,
		},
		{
			`6 4
2 2 2 3 3 3`,
			`2 3
1 2 3`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf632D)
}

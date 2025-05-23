// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1989/B
// https://codeforces.com/problemset/status/1989/problem/B?friends=on
func Test_cf1989B(t *testing.T) {
	testCases := [][2]string{
		{
			`5
aba
cb
er
cf
mmm
mmm
contest
test
cde
abcefg`,
			`4
4
3
7
7`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1989B)
}

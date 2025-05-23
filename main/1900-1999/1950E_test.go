// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1950/E
// https://codeforces.com/problemset/status/1950/problem/E?friends=on
func Test_cf1950E(t *testing.T) {
	testCases := [][2]string{
		{
			`5
4
abaa
4
abba
13
slavicgslavic
8
hshahaha
20
stormflamestornflame`,
			`1
4
13
2
10`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1950E)
}

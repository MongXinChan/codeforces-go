// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/622/F
// https://codeforces.com/problemset/status/622/problem/F?friends=on
func Test_cf622F(t *testing.T) {
	testCases := [][2]string{
		{
			`4 1`,
			`10`,
		},
		{
			`4 2`,
			`30`,
		},
		{
			`4 3`,
			`100`,
		},
		{
			`4 0`,
			`4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf622F)
}

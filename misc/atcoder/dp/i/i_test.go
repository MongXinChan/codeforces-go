// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/dp/submit?taskScreenName=dp_i
func Test_run(t *testing.T) {
	t.Log("Current test is [i]")
	testCases := [][2]string{
		{
			`3
0.30 0.60 0.80`,
			`0.612`,
		},
		{
			`1
0.50`,
			`0.5`,
		},
		{
			`5
0.42 0.01 0.42 0.99 0.42`,
			`0.3821815872`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/dp/tasks/dp_i

package warmup

import "testing"

type testCase struct {
	in     []int
	expect int
}

var testCases = []testCase{
	{in: []int{1, 2}, expect: 3},
	{in: []int{3, 2}, expect: 5},
	{in: []int{1, 4}, expect: 5},
	{in: []int{-3, 105}, expect: 102},
}

func TestSolveMeFirst(t *testing.T) {
	for _, aCase := range testCases {
		result := SolveMeFirst(aCase.in[0], aCase.in[1])
		if result != aCase.expect {
			t.Error(
				"For", aCase.in,
				"expected", aCase.expect,
				"got", result,
			)
		}
	}
}

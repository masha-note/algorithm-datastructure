package linklist

import "testing"

func TestReversal(t *testing.T) {
	testcases := []struct {
		data     *Node
		expected *Node
	}{
		{data: linklistFromSlice([]string{}), expected: nil},
		{data: linklistFromSlice([]string{"1"}), expected: linklistFromSlice([]string{"1"})},
		{
			data:     linklistFromSlice([]string{"1", "2"}),
			expected: linklistFromSlice([]string{"2", "1"}),
		},
		{
			data:     linklistFromSlice([]string{"1", "2", "3"}),
			expected: linklistFromSlice([]string{"3", "2", "1"}),
		},
		{
			data:     linklistFromSlice([]string{"1", "2", "3", "4"}),
			expected: linklistFromSlice([]string{"4", "3", "2", "1"}),
		},
	}
	for _, testcase := range testcases {
		if got := reverse(testcase.data); !compareLinklist(got, testcase.expected) {
			t.Errorf("\nexpected:\n%v\ngot:\n%v",
				printLinklist(testcase.expected),
				printLinklist(got))
		}
	}
}

package sensitive

import (
	"testing"
)

func TestTrieTreeFindIn(t *testing.T) {
	trie := NewTrie()

	trie.Add("abc")
	trie.Add("imabc")

	testcases := []struct {
		Text   string
		Expect string
	}{
		{"abc", "abc"},
		{"abc123", "abc"},
		{"123abc", "abc"},
		{"123abc123", "abc"},
	}

	for _, testcase := range testcases {
		if _, got := trie.FindIn(testcase.Text); got != testcase.Expect {
			if got == "" {
				got = "[nothin]"
			}
			t.Fatalf("find %s, got %s, expect %s\n", testcase.Text, got, testcase.Expect)
		}
	}
}

package words

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []string{"Tiago Temporin", " ", "!33241"}
	for _, orig := range testcases {
		rev, err := Reverse(orig)
		if err != nil {
			return
		}

		rev2, err := Reverse(rev)
		if err != nil {
			return
		}

		if orig != rev2 {
			t.Errorf("esperado %q, obtido :%q", orig, rev2)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Tiago Temporin", " ", "!33241"}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err := Reverse(orig)
		if err != nil {
			return
		}

		rev2, err := Reverse(rev)
		if err != nil {
			return
		}

		if orig != rev2 {
			f.Errorf("esperado %q, obtido :%q", orig, rev2)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("a string reversa não é UTF-8: %q", rev)
		}
	})
}

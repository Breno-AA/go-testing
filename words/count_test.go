package words_test

import (
	"go-testing/words"
	"testing"
)

func TestCount(t *testing.T) {
	testcases := []struct {
		Text     string
		Filter   string
		Expected int
	}{
		{Text: "texto de teste", Expected: 3},
		{Text: "aqui vai o seu texto", Filter: "aqui", Expected: 1},
		{Text: "somente mais um  texto", Expected: 4},
		{Text: "", Expected: 0},
		{Text: "esse é um outro texto com filtro", Filter: "aqui", Expected: 0},
	}

	for _, tc := range testcases {
		c := words.Count(tc.Text, tc.Filter)
		if c != tc.Expected {
			t.Fatalf("expected: %d, got %d", tc.Expected, c)
		}
	}
}

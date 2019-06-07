package jij

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			"simple text",
			`foo`,
			`"foo"`,
		},
		{
			"simple json - array",
			`["foo"]`,
			`"[\"foo\"]"`,
		},
		{
			"simple json - object",
			`{"key":"foo"}`,
			`"{\"key\":\"foo\"}"`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Convert(test.input)
			if got != test.want {
				t.Errorf("want %s, got %s", test.want, got)
			}
		})
	}
}

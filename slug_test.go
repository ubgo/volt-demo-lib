package slug

import "testing"

func TestMake(t *testing.T) {
	cases := map[string]string{
		"Hello, World!":    "hello-world",
		"  spaces  ":       "spaces",
		"already-a-slug":   "already-a-slug",
		"Ünïcode dropped":  "n-code-dropped",
		"":                 "",
		"---":              "",
	}
	for in, want := range cases {
		if got := Make(in); got != want {
			t.Errorf("Make(%q) = %q, want %q", in, got, want)
		}
	}
}

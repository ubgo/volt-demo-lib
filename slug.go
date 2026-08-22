// Package slug is the library-repo demo for volt
// (github.com/khanakia/voltkit): an importable package at the repo root,
// bare vX.Y.Z tags, releases that carry notes but no binaries — the artifact
// that matters is the module proxy entry.
package slug

import "strings"

// Make converts s into a lowercase, hyphen-separated slug. Runs of
// non-alphanumeric characters collapse into one hyphen; leading and trailing
// hyphens are trimmed, so no input produces "-x-" shapes.
func Make(s string) string {
	var b strings.Builder
	prevHyphen := true // suppresses a leading hyphen
	for _, r := range strings.ToLower(s) {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
			prevHyphen = false
		case !prevHyphen:
			b.WriteByte('-')
			prevHyphen = true
		}
	}
	return strings.TrimSuffix(b.String(), "-")
}

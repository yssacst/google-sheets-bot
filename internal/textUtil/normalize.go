package textutil

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

var re = regexp.MustCompile(`[^a-z0-9 ]`)

func NormalizeName(name string) string {
	name = strings.ToLower(name)
	name = norm.NFD.String(name)
	name = strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Mn, r) {
			return -1
		}
		return r
	}, name)
	name = re.ReplaceAllString(name, "")
	name = strings.Join(strings.Fields(name), " ")
	return name
}

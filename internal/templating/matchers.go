package templating

import (
	"regexp"
)

var (
	MatchAnyEnv = regexp.MustCompile(`\$\{(\s+)?env(\s+)?:(\s+)?(\w+)(\s+)?}`)
)

func MatchTemplateData(k string) *regexp.Regexp {
	return regexp.MustCompile(`\$\{(\s+)?` + k + `(\s+)?\}`)
}

func MatchEnv(k string) *regexp.Regexp {
	return regexp.MustCompile(`\$\{(\s+)?env(\s+)?:(\s+)?` + k + `(\s+)?\}`)
}

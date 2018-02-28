package tmplname

import (
	"regexp"
	"strings"
)

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

// Convert ...
func Convert(s string) string {
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	//log.Print(a)
	if len(a) > 1 {
		//log.Print(a[0])
		//	log.Print(a[1:])
		return strings.ToLower(a[0] + "/" + strings.Join(a[1:], "_"))
	}
	return strings.ToLower(a[0])
}

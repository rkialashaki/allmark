package pattern

import "regexp"

// IsMatch returns a flag indicating whether the supplied
// text and pattern do match and if yet, the matched text.
func IsMatch(text string, pattern regexp.Regexp) (isMatch bool, matches []string) {
	matches = pattern.FindStringSubmatch(text)
	return matches != nil, matches
}

func isAnagram(s string, t string) bool {
	re1 := strings.Split(s, "")
	re2 := strings.Split(t, "")

	sort.Strings(re1)
	sort.Strings(re2)

	if s == t {
		return true
	}

	return strings.Join(re1, "") == strings.Join(re2, "")
}
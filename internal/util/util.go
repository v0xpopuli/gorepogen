package util

func Contains(a []string, t string) bool {
	for _, e := range a {
		if e == t {
			return true
		}
	}
	return false
}

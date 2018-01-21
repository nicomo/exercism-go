package accumulate

const testVersion = 1

// Accumulate changes a slice of strings according to a given func
func Accumulate(s []string, f func(string) string) []string {
	var result []string
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}

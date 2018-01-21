// Package letter counts the frequency of letters in texts using parallel computation
package letter

type FreqMap map[rune]int

// Frequency counts letter frequencies
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency uses channels
// to count the occurences of letters (see func Frequency)
// in seperate texts
// and aggregate the results
func ConcurrentFrequency(txts []string) FreqMap {

	c := make(chan FreqMap, len(txts))

	for _, txt := range txts {
		go func(v string) {
			c <- Frequency(v)
		}(txt)
	}

	result := make(FreqMap)

	for range txts {
		for k, val := range <-c {
			result[k] += val
		}
	}

	return result
}

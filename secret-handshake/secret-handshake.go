// Package secret is secret :-)
package secret

// list of actions we're going to iterate against
var secrets = []string{"wink", "double blink", "close your eyes", "jump"}

// Handshake takes an int and outputs a slice of strings as secret handshake
func Handshake(i int) []string {

	// strings to be returned
	var handshake = make([]string, 0)

	if i <= 0 {
		return handshake
	}

	for k, v := range secrets {
		if (1<<uint(k))&i > 0 {
			handshake = append(handshake, v)
		}
	}

	if (1<<uint(len(secrets)))&i > 0 {
		reverse(handshake)
	}

	return handshake

}

func reverse(strings []string) {
	for i, j := 0, len(strings)-1; i < j; i, j = i+1, j-1 {
		strings[i], strings[j] = strings[j], strings[i]
	}
}

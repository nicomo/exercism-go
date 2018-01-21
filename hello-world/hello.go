package hello

const testVersion = 2

// get variable, return string
func HelloWorld(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	return "Hello, " + name + "!"

}

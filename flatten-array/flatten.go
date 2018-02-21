// Package flatten takes a nested list and return a single flattened list with all values except nil/null
package flatten

// Flatten takes nested interfaces of ints and returns them as a flat, not nested slice
func Flatten(input interface{}) []interface{} {
	output := make([]interface{}, 0)

	switch input.(type) {
	case []interface{}:
		for _, v := range input.([]interface{}) {
			output = append(output, Flatten(v)...)
		}
	case int:
		output = append(output, input)
	}

	return output
}

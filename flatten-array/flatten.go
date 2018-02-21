// Package flatten takes a nested list and return a single flattened list with all values except nil/null
package flatten

import (
	"reflect"
)

// Flatten takes nested interfaces of ints and returns them as a flat, not nested slice
func Flatten(input interface{}) []interface{} {
	output := make([]interface{}, 0)

	switch reflect.TypeOf(input).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(input)
		for i := 0; i < s.Len(); i++ {
			v := s.Index(i).Interface()
			_, ok := v.(int)
			if !ok {
				interfaceToInt(v, &output)
			} else {
				output = append(output, v)
			}

		}
	}

	return output
}

// interfaceToInt recursively seeks ints in interfaces values
func interfaceToInt(iface interface{}, output *[]interface{}) {
	switch t := iface.(type) {

	case []interface{}:
		for _, val := range t {
			interfaceToInt(val, output)
		}

	case int:
		*output = append(*output, t)
	}
}

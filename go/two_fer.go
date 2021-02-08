// Package twofer implements ShareWith function
package twofer

// ShareWith share the input with the "name" passed to the function.
// if "name" is empty, share it with "you"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}

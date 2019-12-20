package ternary

// Ternary gives you a simple Golang ternary expression
func Ternary(a bool, b, c interface{}) interface{} {
	if a {
		return b
	}

	return c
}

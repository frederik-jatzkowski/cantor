// Package testsuites contains test cases for all interfaces exported from [pkg/github.com/frederik-jatzkowski/cantor].
// These can be used to validate implementations of these interfaces.
package sets

type Constructor[E any, T any] func(elements ...E) T

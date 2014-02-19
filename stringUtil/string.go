package stringUtil

import (
	"bytes"
	"strconv"
)

type TestObj struct {
	A, B int
}

// Reverse function will reverse a string input passed to this function.
// You can pass Unicode characters as well to get the desired output.
func Reverse(s string) string {
	b := []rune(s)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func (s *TestObj) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(strconv.Itoa(s.A))
	buffer.WriteString(strconv.Itoa(s.B))

	return buffer.String()
}

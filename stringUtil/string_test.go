package stringUtil

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"Backward", "drawkcaB"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range tests {
		got := Reverse(c.s)
		if got != c.want {
			t.Error("Reverse(%q) == %q, want %q,", c.s, got, c.want)
		}
	}
}

func TestStringTestObj(t *testing.T) {
	var tests = []struct {
		s    TestObj
		want string
	}{
		{TestObj{1, 2}, "12"},
		{TestObj{-1, -1}, "-1-1"},
		{TestObj{}, "00"},
	}
	for _, c := range tests {
		got := c.s.String()
		if got != c.want {
			t.Error("TestObj String(%q) == %q, want %q,", c.s.String(), got, c.want)
		}
	}
}

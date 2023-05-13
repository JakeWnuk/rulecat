package utils

import "testing"

func TestLenToRule(t *testing.T) {
	tests := []struct {
		str  string
		rule string
		want string
	}{
		{"hello", "]", "] ] ] ] ]"},
		{"world", "[", "[ [ [ [ ["},
	}

	for _, test := range tests {
		got := LenToRule(test.str, test.rule)
		if got != test.want {
			t.Errorf("LenToRule(%q, %q) = %q; want %q", test.str, test.rule, got, test.want)
		}
	}
}

func TestCharToRule(t *testing.T) {
	tests := []struct {
		str  string
		rule string
		want string
	}{
		{"hello", "^", "^h ^e ^l ^l ^o"},
		{"world", "$", "$w $o $r $l $d"},
	}

	for _, test := range tests {
		got := CharToRule(test.str, test.rule)
		if got != test.want {
			t.Errorf("CharToRule(%q, %q) = %q; want %q", test.str, test.rule, got, test.want)
		}
	}
}

func TestCharToIteratingRule(t *testing.T) {
	tests := []struct {
		str   string
		rule  string
		index int
		want  string
	}{
		{"hello", "i", 0, "i0h i1e i2l i3l i4o"},
		{"world", "o", 6, "o6w o7o o8r o9l oAd"},
	}

	for _, test := range tests {
		got := CharToIteratingRule(test.str, test.rule, test.index)
		if got != test.want {
			t.Errorf("CharToIteratingRule(%q, %q, %d) = %q; want %q", test.str, test.rule, test.index, got, test.want)
		}
	}
}

func TestStringToToggle(t *testing.T) {
	tests := []struct {
		str   string
		rule  string
		index int
		want  string
	}{
		{"HelloWorld", "T", 0, "T0 T5"},
		{"HelloWorld", "T", 5, "T5 TA"},
	}

	for _, test := range tests {
		got := StringToToggle(test.str, test.rule, test.index)
		if got != test.want {
			t.Errorf("StringToToggle(%q, %q, %d) = %q; want %q", test.str, test.rule, test.index, got, test.want)
		}
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
	}

	for _, test := range tests {
		got := ReverseString(test.str)
		if got != test.want {
			t.Errorf("ReverseString(%q) = %q; want %q", test.str, got, test.want)
		}
	}
}

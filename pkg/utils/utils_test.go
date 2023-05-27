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

func TestCheckASCIIString(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "All ASCII characters",
			str:  "Hello, World!",
			want: true,
		},
		{
			name: "Contains non-ASCII character",
			str:  "Hello, 世界!",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckASCIIString(tt.str)
			if got != tt.want {
				t.Errorf("CheckASCIIString(%q) = %v, want %v", tt.str, got, tt.want)
			}
		})
	}
}

func TestConvertCharacterMultiByteString(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "All ASCII characters",
			str:  "$H $e $l $l $o $ $W $o $r $l $d $!",
			want: "$H $e $l $l $o $ $W $o $r $l $d $!",
		},
		{
			name: "Contains non-ASCII character",
			str:  "$H $e $l $l $o $  $世 $界 $!",
			want: "$H $e $l $l $o $  $\\xE4 $\\xB8 $\\x96 $\\xE7 $\\x95 $\\x8C $!",
		},
		{
			name: "Contains non-ASCII character with ^",
			str:  "^! ^界 ^世 ^  ^o ^l ^l ^e ^H",
			want: "^! ^\\x8C ^\\x95 ^\\xE7 ^\\x96 ^\\xB8 ^\\xE4 ^  ^o ^l ^l ^e ^H",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertCharacterMultiByteString(tt.str)
			if got != tt.want {
				t.Errorf("ConvertCharacterMultiByteString(%q) = %v, want %v", tt.str, got, tt.want)
			}
		})
	}
}

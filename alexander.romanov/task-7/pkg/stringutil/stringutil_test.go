package stringutil

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"palindrome", "racecar", "racecar"},
		{"normal string", "Hello, World", "dlroW ,olleH"},
		{"with numbers", "12345", "54321"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.in)
			if got != tt.want {
				t.Errorf("Reverse(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"empty string", "", true},
		{"single character", "a", true},
		{"actual palindrome", "racecar", true},
		{"not a palindrome", "hello", false},
		{"unicode palindrome", "世界界世", true},
		{"case sensitive", "Racecar", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindrome(tt.in)
			if got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

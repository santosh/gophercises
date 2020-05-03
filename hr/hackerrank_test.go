package hackerrank

import "testing"

func TestCountWord(t *testing.T) {
	tests := []struct {
		given    string
		expected int
	}{
		{"santoshKumar", 2},
		{"endOfTheWorld", 4},
		{"coronaVirusPandemic", 3},
	}

	for _, tcase := range tests {
		t.Run(tcase.given, func(t *testing.T) {
			output := CountWord(tcase.given)
			if output != tcase.expected {
				t.Errorf("got %q, want %q", output, tcase.expected)
			}
		})
	}
}

func TestCaesarCipher(t *testing.T) {
	tests := []struct {
		name     string
		given    string
		expected string
	}{
		{"test with single word", "santosh", "xfsytxm"},
		{"test with spaces in between", "santosh kumar", "xfsytxm pzrfw"},
		{"test with other special characters", "i-am-the-one", "n-fr-ymj-tsj"},
		{"test for rotating letter", "zebra", "ejgwf"},
	}

	// Using constant key, but tests cases
	// make sure letter rotates
	key := 5
	for _, tcase := range tests {
		t.Run(tcase.name, func(t *testing.T) {
			output := CaesarCipher(tcase.given, int32(key))
			if output != tcase.expected {
				t.Errorf("got %q, want %q", output, tcase.expected)
			}
		})
	}
}

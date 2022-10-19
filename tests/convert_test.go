package mathconv

import (
	"testing"

	"github.com/narutopig/mathconv"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name string
		args int64
		want string
	}{
		{"Zero", 0, "zero"},
		{"Single digit", 8, "eight"},
		{"Double digit", 69, "sixty-nine"},
		{"Large number", 123456, "one hundred twenty-three thousand, four hundred fifty-six"},
		{"Larger number", 727727727, "seven hundred twenty-seven million, seven hundred twenty-seven thousand, seven hundred twenty-seven"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mathconv.Convert(tt.args); got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

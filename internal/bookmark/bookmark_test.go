package bookmark

import (
	"errors"
	"fmt"
	"testing"
)

func TestValidateUrl(t *testing.T) {
	tt := []struct {
		rawUrl   string
		expected error
	}{
		{
			rawUrl:   "https://google.com",
			expected: nil,
		},
	}

	for _, tc := range tt {
		got := validate(tc.rawUrl)

		if !errors.Is(got, tc.expected) {
			fmt.Printf("expected: %v, got: %v\n", tc.expected, got)
			t.Fail()
		}
	}
}

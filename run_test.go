package validate

import (
	"testing"
)

func Test_Run(t *testing.T) {
	testcases := []struct {
		name      string
		rules     []Rule
		expectErr bool
	}{
		{
			name: "Valid",
			rules: []Rule{
				Exists("value", "test"),
			},
			expectErr: false,
		},
		{
			name: "Multiple Valid",
			rules: []Rule{
				Exists("good1", "test"),
				Length("good2", []int{1, 2, 3}, 5),
				Between("good3", 1, 0, 5),
			},
			expectErr: false,
		},
		{
			name: "Multiple Invalid",
			rules: []Rule{
				Exists("bad1", nil),
				Length("bad2", []int{1, 2, 3}, 2),
				Between("bad3", 10, 0, 5),
			},
			expectErr: true,
		},
		{
			name: "Multiple Mixed",
			rules: []Rule{
				Exists("good1", "test"),
				Length("bad2", []int{1, 2, 3}, 2),
				Between("good3", 1, 0, 5),
			},
			expectErr: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := Run(tc.rules...)
			if tc.expectErr && err == nil {
				t.Error("Expected error, got nil")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}

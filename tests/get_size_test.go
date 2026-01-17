package tests

import (
    "testing"
    "code"

    "github.com/stretchr/testify/require"
)


func TestGetPathSize(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		expected  string
		human     bool
		all       bool
		recursive bool
	}{
		{
			name:     "file",
			path:     "../testdata/filename",
			expected: "349B ../testdata/filename",
		},
		{
			name:     "not empty directory without hidden",
			path:     "../testdata/not_empty",
			expected: "1048576B ../testdata/not_empty",
		},
		{
			name:     "not empty directory with recursive",
			path:     "../testdata/not_empty",
			expected: "2097152B ../testdata/not_empty",
			recursive: true,
		},
		{
			name:     "not empty directory with human format",
			path:     "../testdata/not_empty",
			expected: "2.0MB ../testdata/not_empty",
			human:    true,
			recursive: true,
		},
		{
			name:     "not empty directory with human, all and recursive",
			path:     "../testdata/not_empty",
			expected: "5.8MB ../testdata/not_empty",
			human:    true,
			all:      true,
			recursive: true,
		},
		{
			name:     "not empty directory with all and recursive",
			path:     "../testdata/not_empty",
			expected: "6106896B ../testdata/not_empty",
			all:      true,
			recursive: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := code.GetPathSize(tt.path, tt.recursive, tt.human, tt.all)
			require.NoError(t, err)
			require.Equal(t, tt.expected, got)
		})
	}
}


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
			expected: "349B",
		},
		{
			name:     "not empty directory without hidden",
			path:     "../testdata/not_empty",
			expected: "1048576B",
		},
		{
			name:     "not empty directory with recursive",
			path:     "../testdata/not_empty",
			expected: "2097152B",
			recursive: true,
		},
		{
			name:     "not empty directory with human format",
			path:     "../testdata/not_empty",
			expected: "2.0MB",
			human:    true,
			recursive: true,
		},
		{
			name:     "not empty directory with human, all and recursive",
			path:     "../testdata/not_empty",
			expected: "5.8MB",
			human:    true,
			all:      true,
			recursive: true,
		},
		{
			name:     "not empty directory with all and recursive",
			path:     "../testdata/not_empty",
			expected: "6106896B",
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


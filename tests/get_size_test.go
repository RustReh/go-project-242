package tests

import (
	"os"
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
		{
			name:      "nonexistent path",
			path:      "../testdata/nonexistent",
			expected:  "",
			recursive: false,
			human:     false,
			all:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "nonexistent path" {
				got, err := code.GetPathSize(tt.path, tt.recursive, tt.human, tt.all)
				require.Error(t, err)
				require.Equal(t, "", got)
				return
			}

			got, err := code.GetPathSize(tt.path, tt.recursive, tt.human, tt.all)
			require.NoError(t, err)
			require.Equal(t, tt.expected, got)
		})
	}
}


func TestGetPathSizeWithEmptyDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	emptyDir := tmpDir + "/empty"
	err := os.Mkdir(emptyDir, 0755)
	require.NoError(t, err)

	result, err := code.GetPathSize(emptyDir, false, false, false)
	require.NoError(t, err)
	require.Equal(t, "0B", result)
}


func TestGetPathSizeWithSymlink(t *testing.T) {
	tmpDir := t.TempDir()

	targetFile := tmpDir + "/target"
	err := os.WriteFile(targetFile, []byte("test"), 0644)
	require.NoError(t, err)

	symlinkPath := tmpDir + "/link"
	err = os.Symlink(targetFile, symlinkPath)
	require.NoError(t, err)

	result, err := code.GetPathSize(symlinkPath, false, false, false)
	require.NoError(t, err)

	require.NotEmpty(t, result)
	require.Contains(t, result, "B")
}

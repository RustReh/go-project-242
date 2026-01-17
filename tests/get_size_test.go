package tests

import (
    "testing"
    "code"

    "github.com/stretchr/testify/require"
)


func TestGetPathSize(t *testing.T) {
    tests := []struct {
        name     string
        path     string
        expected string
        human bool
    }{
        {
            name:     "not empty directory",
            path:     "../testdata/not_empty",
            expected: "1048576B ../testdata/not_empty",
        },
        {
            name:     "file",
            path:     "../testdata/filename",
            expected: "349B ../testdata/filename",
        },
        {
            name:     "not empty directory with human format",
            path:     "../testdata/not_empty",
            expected: "1.0MB ../testdata/not_empty",
            human: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := code.GetPathSize(tt.path, false, tt.human, false)
            require.NoError(t, err)
            require.Equal(t, tt.expected, got)
        })
    }
}


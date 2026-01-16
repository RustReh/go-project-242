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
    }{
        {
            name:     "not empty directory",
            path:     "../testdata/not_empty",
            expected: "349 ../testdata/not_empty",
        },
        {
            name:     "file",
            path:     "../testdata/filename",
            expected: "349 ../testdata/filename",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := code.GetPathSize(tt.path, false, false, false)
            require.NoError(t, err)
            require.Equal(t, tt.expected, got)
        })
    }
}


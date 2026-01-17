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
        all bool
    }{
        {
            name:     "not empty directory without hidde",
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
        {
            name:     "not empty directory with human and all",
            path:     "../testdata/not_empty",
            expected: "2.9MB ../testdata/not_empty",
            human: true,
            all: true,
        },
        {
            name:     "not empty directory with all",
            path:     "../testdata/not_empty",
            expected: "3053448B ../testdata/not_empty",
            all: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := code.GetPathSize(tt.path, false, tt.human, tt.all)
            require.NoError(t, err)
            require.Equal(t, tt.expected, got)
        })
    }
}


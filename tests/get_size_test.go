package tests


import (
	"testing"
	"code"
)


func TestGetSize(t *testing.T) {
    tests := []struct {
        name     string
        path     string
        expected string
    }{
        {
            name:     "empty directory",
            path:     "../testdata/empty",
            expected: "0 ../testdata/empty",
        },
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
            if err != nil {
                t.Fatalf("GetPathSize failed: %v", err)
            }

            if got != tt.expected {
                t.Errorf("got %q, want %q", got, tt.expected)
            }
        })
    }
}

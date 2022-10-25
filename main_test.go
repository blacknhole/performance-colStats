package main

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name   string
		files  []string
		op     string
		col    int
		exp    string
		expErr error
	}{
		{name: "RunAvg1File",
			files: []string{"./testdata/example.csv"},
			op:    "avg", col: 3, exp: "227.6\n", expErr: nil,
		},
		{name: "RunAvgMultiFiles",
			files: []string{"./testdata/example.csv", "./testdata/example2.csv"},
			op:    "avg", col: 3, exp: "233.84\n", expErr: nil,
		},
		{name: "RunFailRead",
			files: []string{"./testdata/example.csv", "./testdata/fakefile.csv"},
			op:    "avg", col: 3, exp: "", expErr: os.ErrNotExist,
		},
		{name: "RunFailColumn",
			files: []string{"./testdata/example.csv"},
			op:    "avg", col: 0, exp: "", expErr: ErrInvalidColumn,
		},
		{name: "RunFailNoFiles",
			files: []string{},
			op:    "avg", col: 2, exp: "", expErr: ErrNoFiles,
		},
		{name: "RunFailOperation",
			files: []string{"./testdata/example.csv"},
			op:    "invalid", col: 2, exp: "", expErr: ErrInvalidOperation,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := run(tc.files, tc.op, tc.col, &buf)
			if tc.expErr != nil {
				if err == nil {
					t.Error("Expected err. Got nil instead")
				}
				if !errors.Is(err, tc.expErr) {
					t.Errorf("Expected %q, got %q instead", tc.expErr, err)
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error: %q", err)
			}
			if buf.String() != tc.exp {
				t.Errorf("Expected %q, got %q instead", tc.exp, buf.String())
			}
		})
	}
}

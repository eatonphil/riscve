package riscve

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name         string
		input        []byte
		program      *program
		errSubstring string
	}{
		{
			"basic instruction",
			[]byte("add x1, x2, x1"),
			&program{
				[]instruction{
					{
						&opcodes[opcodeIndex("add")],
						[]string{"x1", "x2", "x1"},
					},
				},
				map[string]int{},
			},
			"",
		},
		{
			"no such instruction",
			[]byte("div x1, x2, x1"),
			nil,
			"valid opcode",
		},
		{
			"basic label",
			[]byte("foo:"),
			&program{
				nil,
				map[string]int{"foo": 0},
			},
			"",
		},
		{
			"invalid label",
			[]byte(":"),
			nil,
			"valid label",
		},
		{
			"basic label and instructions",
			[]byte("sub x1, x2, x1\nfoo:"),
			&program{
				[]instruction{
					{
						&opcodes[opcodeIndex("sub")],
						[]string{"x1", "x2", "x1"},
					},
				},
				map[string]int{"foo": 1},
			},
			"",
		},
	}

	for _, test := range tests {
		p, e := Parse(test.input)

		assert.Equal(t, p, test.program, test.name)

		if e == nil && test.errSubstring != "" {
			t.Errorf("Expected error like '%s', got nil in '%s'", test.errSubstring, test.name)
		} else if e != nil && test.errSubstring == "" {
			t.Errorf("Unexpected error '%s' in '%s'", e.Error(), test.name)
		} else if e != nil {
			assert.True(t, strings.Contains(e.Error(), test.errSubstring), test.name)
		}
	}
}

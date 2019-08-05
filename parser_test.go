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
						opcodeRepr("add"),
						[]register{x1, x2, x1},
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
			"valid instruction",
		},
		{
			"register pseudonyms and char cases",
			[]byte("ADD pc, SP, GP"),
			&program{
				[]instruction{
					{
						opcodeRepr("add"),
						[]register{pc, x2, x3},
					},
				},
				map[string]int{},
			},
			"",
		},
		{
			"bad register",
			[]byte("ADD rt"),
			nil,
			"valid register",
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
						opcodeRepr("sub"),
						[]register{x1, x2, x1},
					},
				},
				map[string]int{"foo": 1},
			},
			"",
		},
	}

	for _, test := range tests {
		p, e := Parse(test.input)

		assert.Equal(t, test.program, p, test.name)

		if e == nil && test.errSubstring != "" {
			t.Errorf("Expected error like '%s', got nil in '%s'", test.errSubstring, test.name)
		} else if e != nil && test.errSubstring == "" {
			t.Errorf("Unexpected error '%s' in '%s'", e.Error(), test.name)
		} else if e != nil {
			assert.True(t, strings.Contains(e.Error(), test.errSubstring), test.name)
		}
	}
}

package riscve

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	add1, _ := instructionRepr("add")
	add1.args = []string{"x1", "x2", "x1"}
	add2, _ := instructionRepr("add")
	add2.args = []string{"pc", "sp", "gp"}

	sub, _ := instructionRepr("sub")
	sub.args = []string{"x1", "x2", "x1"}

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
				[]instruction{add1},
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
				[]instruction{add2},
				map[string]int{},
			},
			"",
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
			"basic instruction and label",
			[]byte("sub x1, x2, x1\nfoo:"),
			&program{
				[]instruction{sub},
				map[string]int{"foo": 1},
			},
			"",
		},
		{
			"basic label and instruction",
			[]byte("foo:\n\tsub x1, x2, x1"),
			&program{
				[]instruction{sub},
				map[string]int{"foo": 0},
			},
			"",
		},
	}

	for _, test := range tests {
		p, e := Parse(test.input)

		assert.Equal(t, test.program, p, test.name)

		if e == nil {
			for _, i := range p.instructions {
				assert.True(t, i.valid(), test.name)
			}
		}

		if e == nil && test.errSubstring != "" {
			t.Errorf("Expected error like '%s', got nil in '%s'", test.errSubstring, test.name)
		} else if e != nil && test.errSubstring == "" {
			t.Errorf("Unexpected error '%s' in '%s'", e.Error(), test.name)
		} else if e != nil {
			assert.True(t, strings.Contains(e.Error(), test.errSubstring), test.name)
		}
	}
}

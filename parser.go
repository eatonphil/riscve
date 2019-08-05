package riscve

import "fmt"

type program struct {
	instructions []instruction
	labels       map[string]int
}

type syntax string

const (
	isyntax syntax = "instruction"
	lsyntax syntax = "syntax"
)

func Parse(cs []byte) (*program, error) {
	var partial instruction
	var partialOpcodeSet bool = false
	var p program = program{nil, map[string]int{}}
	var currentToken = ""
	var col = 0
	var row = 1
	var s syntax = isyntax
	var inComment bool = false

	// For cleanliness of parsing, add a trailing newline if there isn't one
	if cs[len(cs)-1] != '\n' {
		cs = append(cs, '\n')
	}

	for _, c := range cs {
		// Comma here is more permissive than should be
		whitespace := c == ' ' || c == '\t' || c == ',' || c == '\n'
		if whitespace {
			if s == isyntax && currentToken != "" {
				if !partialOpcodeSet {
					partialOpcodeSet = true
					var ok = false
					partial, ok = instructionRepr(currentToken)
					if !ok {
						return nil, fmt.Errorf("Expected valid instruction, got '%s' near %d,%d", currentToken, row, col-len(currentToken))
					}
				} else {
					partial.args = append(partial.args, currentToken)
				}
			} else if s == lsyntax {
				p.labels[currentToken] = row - 1
			}

			currentToken = ""

			if c == '\n' {
				if partialOpcodeSet {
					p.instructions = append(p.instructions, partial)
				}

				partial = instruction{}
				partialOpcodeSet = false
				s = isyntax

				inComment = false
				row++
				col = 0
				continue
			}

			col++
			continue
		}

		if inComment {
			col++
			continue
		}

		if c == '#' {
			inComment = true
			col++
			continue
		}

		if c == ':' {
			if currentToken == "" {
				return nil, fmt.Errorf("Expected valid label near %d,%d", row, col-len(currentToken))
			}

			s = lsyntax
			col++
			continue
		}

		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			// Map to lowercase
			if c >= 'A' && c <= 'Z' {
				c = byte(int(c) + ('a' - 'A'))
			}

			currentToken += string(c)
			col++
		}
	}

	return &p, nil
}

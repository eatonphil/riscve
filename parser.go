package riscve

import "fmt"

type instruction struct {
	opcode *opcode
	args   []string
}

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
		whitespace := c == ' ' || c == '\t' || c == ','
		if whitespace {
			if s == isyntax && currentToken != "" {
				if partial.opcode == nil {
					oi := opcodeIndex(currentToken)
					if oi == -1 {
						return nil, fmt.Errorf("Expected valid opcode, got '%s' near %d,%d", currentToken, row, col-len(currentToken))
					}

					partial.opcode = &opcodes[oi]
				} else {
					partial.args = append(partial.args, currentToken)
				}
			}

			currentToken = ""
			col++
			continue
		}

		if c == '\n' {
			if inComment {
				inComment = false
			}

			if s == isyntax {
				if currentToken != "" {
					partial.args = append(partial.args, currentToken)
					currentToken = ""
				}

				p.instructions = append(p.instructions, partial)
				partial = instruction{}
			} else {
				p.labels[currentToken] = row - 1
			}

			s = isyntax
			row++
			col = 0
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

		if (c >= 'A' && c <= 'z') || (c >= '0' && c <= '9') {
			// Map to lowercase
			if c >= 'A' && c <= 'Z' {
				c = byte(int(c) + 26)
			}

			currentToken += string(c)
			col++
		}
	}

	return &p, nil
}

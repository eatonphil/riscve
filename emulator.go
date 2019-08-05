package riscve

import "fmt"

type cpu struct {
	registers map[register]int64
	stack     [8096]uint8
	program   *program
}

func (c *cpu) loadProgram(p *program) error {
	c.program = p

	main := false
	for label, i := range p.labels {
		if label == "main" {
			main = true
			c.registers[pc] = int64(i)
		}
	}

	if !main {
		return fmt.Errorf("Expected main")
	}

	return nil
}

func (c *cpu) emulateProgram() error {
	for _, instr := range c.program.instructions {
		if !instr.valid() {
			return fmt.Errorf("Invalid instruction %+v", instr)
		}

		switch instr.repr {
		case "add":
			rd := instr.rd()
			rs1 := instr.rs(1)
			rs2 := instr.rs(2)
			c.registers[rd] = c.registers[rs1] + c.registers[rs2]
		case "sub":
			rd := instr.rd()
			rs1 := instr.rs(1)
			rs2 := instr.rs(2)
			c.registers[rd] = c.registers[rs1] - c.registers[rs2]
		case "addi":
			rd := instr.rd()
			rs1 := instr.rs(1)
			imm := instr.imm(2)
			c.registers[rd] = c.registers[rs1] + imm
		case "mv":
			rd := instr.rd()
			rs1 := instr.rs(1)
			c.registers[rd] = c.registers[rs1]
		case "li":
			rd := instr.rd()
			imm := instr.imm(1)
			c.registers[rd] = imm
		case "ret":
			fmt.Printf("(exit)%d\n", c.registers[a0])
			return nil
		}

		c.registers[pc]++
	}

	return nil
}

func initializeCpu() *cpu {
	c := cpu{}

	c.registers = map[register]int64{}
	for _, r := range registers {
		c.registers[r] = 0

		// Initialize stack pointer
		if r == sp {
			c.registers[r] = int64(len(c.stack))
		}
	}

	return &c
}

func Emulate(p *program) error {
	c := initializeCpu()
	err := c.loadProgram(p)
	if err != nil {
		return err
	}

	return c.emulateProgram()
}

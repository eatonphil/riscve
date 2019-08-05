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
		switch instr.opcode.repr {
		case "add":

		}
	}

	return nil
}

func initializeCpu() *cpu {
	c := cpu{}

	rs := map[register]int64{}
	for _, r := range registers {
		rs[r] = 0

		// Initialize stack pointer
		if r == sp {
			rs[r] = int64(len(c.stack))
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

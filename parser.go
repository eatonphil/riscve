package riscve

type instruction struct {
	opcode opcode
}

type program struct {
	instructions []instruction
	labels       map[string]int
}

func Parse(f []byte) (*program, error) {
	return nil, nil
}

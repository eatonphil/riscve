package riscve

type itype int

const (
	rType itype = iota
	iType
	sType
	bType
	uType
	jType
)

type opcode struct {
	repr   string
	itype  itype
	opcode int
	func3  int
	func7  int
}

var opcodes = []opcode{
	{
		"lui",
		uType,
		b0110111,
	},
	{
		"auipc",
		uType,
		b0010111,
	},
	{
		"jal",
		jType,
		b1101111,
	},
	{
		"jalr",
		iType,
		b1100111,
	},
	{
		"beq",
		bType,
		b1100011,
	},
	{
		"bne",
		bType,
		b1100011,
		b001,
	},
	{
		"blt",
		bType,
		b1100011,
		b100,
	},
	{
		"blt",
		bType,
		b1100011,
		b100,
	},
	{
		"bge",
		bType,
		b1100011,
		b101,
	},
	{
		"bltu",
		bType,
		b1100011,
		b110,
	},
	{
		"bgeu",
		bType,
		b1100011,
		b111,
	},
	{
		"lb",
		iType,
		b0000011,
	},
	{
		"lh",
		iType,
		b0000011,
		b001,
	},
	{
		"lw",
		iType,
		b0000011,
		b010,
	},
	{
		"lw",
		iType,
		b0000011,
		b010,
	},
}

type register int

const (
	x0 register = iota
	x1, x2, x3, x4, x5, x6, x7, x8,
	x9, x10, x11, x12, x13, x4, x15, x16,
	x17, x18, x19, x20, x21, x22, x23, x24,
	x25, x26, x27, x28, x29, x30, x31
)

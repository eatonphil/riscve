package riscve

import (
	"strconv"
)

type itype int

const (
	rType itype = iota
	iType
	sType
	bType
	uType
	jType
	pType // pseudo
)

type opcode struct {
	repr   string
	itype  itype
	opcode int
	func3  int
	func7  int
	psuedo bool
}

func b(s string) int {
	i, _ := strconv.ParseInt(s, 2, 64)
	return int(i)
}

var opcodes = []opcode{
	{
		"lui",
		uType,
		b("0110111"),
		0,
		0,
		false,
	},
	{
		"auipc",
		uType,
		b("0010111"),
		0,
		0,
		false,
	},
	{
		"jal",
		jType,
		b("1101111"),
		0,
		0,
		false,
	},
	{
		"jalr",
		iType,
		b("1100111"),
		0,
		0,
		false,
	},
	{
		"beq",
		bType,
		b("1100011"),
		0,
		0,
		false,
	},
	{
		"bne",
		bType,
		b("1100011"),
		b("001"),
		0,
		false,
	},
	{
		"blt",
		bType,
		b("1100011"),
		b("100"),
		0,
		false,
	},
	{
		"bge",
		bType,
		b("1100011"),
		b("101"),
		0,
		false,
	},
	{
		"bltu",
		bType,
		b("1100011"),
		b("110"),
		0,
		false,
	},
	{
		"bgeu",
		bType,
		b("1100011"),
		b("111"),
		0,
		false,
	},
	{
		"lb",
		iType,
		b("0000011"),
		0,
		0,
		false,
	},
	{
		"lh",
		iType,
		b("0000011"),
		b("001"),
		0,
		false,
	},
	{
		"lw",
		iType,
		b("0000011"),
		b("010"),
		0,
		false,
	},
	{
		"lbu",
		iType,
		b("0000011"),
		b("100"),
		0,
		false,
	},
	{
		"lhu",
		iType,
		b("0000011"),
		b("101"),
		0,
		false,
	},
	{
		"sb",
		sType,
		b("0100011"),
		0,
		0,
		false,
	},
	{
		"sh",
		sType,
		b("0100011"),
		b("001"),
		0,
		false,
	},
	{
		"sw",
		sType,
		b("0100011"),
		b("010"),
		0,
		false,
	},
	{
		"addi",
		iType,
		b("0010011"),
		0,
		0,
		false,
	},
	{
		"slti",
		iType,
		b("0010011"),
		b("010"),
		0,
		false,
	},
	{
		"sltiu",
		iType,
		b("0010011"),
		b("011"),
		0,
		false,
	},
	{
		"xori",
		iType,
		b("0010011"),
		b("100"),
		0,
		false,
	},
	{
		"ori",
		iType,
		b("0010011"),
		b("110"),
		0,
		false,
	},
	{
		"andi",
		iType,
		b("0010011"),
		b("111"),
		0,
		false,
	},
	{
		"slli",
		rType,
		b("0010011"),
		b("001"),
		0,
		false,
	},
	{
		"srli",
		rType,
		b("0010011"),
		b("101"),
		0,
		false,
	},
	{
		"srai",
		rType,
		b("0010011"),
		b("101"),
		b("0100000"),
		false,
	},
	{
		"add",
		rType,
		b("0110011"),
		0,
		0,
		false,
	},
	{
		"sub",
		rType,
		b("0110011"),
		0,
		b("0100000"),
		false,
	},
	{
		"sll",
		rType,
		b("0110011"),
		b("001"),
		0,
		false,
	},
	{
		"slt",
		rType,
		b("0110011"),
		b("010"),
		0,
		false,
	},
	{
		"sltu",
		rType,
		b("0110011"),
		b("011"),
		0,
		false,
	},
	{
		"xor",
		rType,
		b("0110011"),
		b("100"),
		0,
		false,
	},
	{
		"srl",
		rType,
		b("0110011"),
		b("101"),
		0,
		false,
	},
	{
		"sra",
		rType,
		b("0110011"),
		b("101"),
		b("0100000"),
		false,
	},
	{
		"or",
		rType,
		b("0110011"),
		b("110"),
		0,
		false,
	},
	{
		"and",
		rType,
		b("0110011"),
		b("111"),
		0,
		false,
	},
	{
		"fence",
		iType,
		b("0001111"),
		0,
		0,
		false,
	},
	{
		"ecall",
		iType,
		b("1110011"),
		0,
		0,
		false,
	},
	{
		"ebreak",
		iType,
		b("1110011"),
		0,
		1,
		false,
	},
	{
		"call",
		pType,
		0,
		0,
		0,
		true,
	},
	{
		"ret",
		pType,
		0,
		0,
		0,
		true,
	},
	{
		"mv",
		pType,
		0,
		0,
		0,
		true,
	},
}

func opcodeRepr(repr string) *opcode {
	for _, oc := range opcodes {
		if oc.repr == repr {
			return &oc
		}
	}

	return nil
}

type register int

const (
	x0 register = iota
	x1
	x2
	x3
	x4
	x5
	x6
	x7
	x8
	x9
	x10
	x11
	x12
	x13
	x14
	x15
	x16
	x17
	x18
	x19
	x20
	x21
	x22
	x23
	x24
	x25
	x26
	x27
	x28
	x29
	x30
	x31
)

var registers = []register{
	x0,
	x1,
	x2,
	x3,
	x4,
	x5,
	x6,
	x7,
	x8,
	x9,
	x10,
	x11,
	x12,
	x13,
	x14,
	x15,
	x16,
	x17,
	x18,
	x19,
	x20,
	x21,
	x22,
	x23,
	x24,
	x25,
	x26,
	x27,
	x28,
	x29,
	x30,
	x31,
}

// register psuedoynms
const (
	pc = -1 // program counter

	zero = x0
	ra   = x1 // return address
	sp   = x2 // stack pointer
	gp   = x3 // global pointer
	tp   = x4 // thread pointer

	t0 = x5 // temporary/alternative link register

	t1 = x6 // temporaries
	t2 = x7

	s0 = x8 // saved register
	fp = x8 // frame pointer
	s1 = x9 // saved register

	a0 = x10 // function arguments/return values
	a1 = x11

	a2 = x12 //function arguments
	a3 = x13
	a4 = x14
	a5 = x15
	a6 = x16
	a7 = x17

	s2  = x18 // saved registers
	s3  = x19
	s4  = x20
	s5  = x21
	s6  = x22
	s7  = x23
	s8  = x24
	s9  = x25
	s10 = x26
	s11 = x27

	t3 = x28 // temporaries
	t4 = x29
	t5 = x30
	t6 = x31
)

var registerRepr = map[string]register{
	"x0":  x0,
	"x1":  x1,
	"x2":  x2,
	"x3":  x3,
	"x4":  x4,
	"x5":  x5,
	"x6":  x6,
	"x7":  x7,
	"x8":  x8,
	"x9":  x9,
	"x10": x10,
	"x11": x11,
	"x12": x12,
	"x13": x13,
	"x14": x14,
	"x15": x15,
	"x16": x16,
	"x17": x17,
	"x18": x18,
	"x19": x19,
	"x20": x20,
	"x21": x21,
	"x22": x22,
	"x23": x23,
	"x24": x24,
	"x25": x25,
	"x26": x26,
	"x27": x27,
	"x28": x28,
	"x29": x29,
	"x30": x30,
	"x31": x31,
	"pc":  pc,

	// pseudonyms
	"zero": zero,
	"ra":   ra,
	"sp":   sp,
	"gp":   gp,
	"tp":   tp,
	"t0":   t0,
	"t1":   t1,
	"t2":   t2,
	"s0":   s0,
	"fp":   fp,
	"s1":   s1,
	"a0":   a0,
	"a1":   a1,
	"a2":   a2,
	"a3":   a3,
	"a4":   a4,
	"a5":   a5,
	"a6":   a6,
	"a7":   a7,
	"s2":   s2,
	"s3":   s3,
	"s4":   s4,
	"s5":   s5,
	"s6":   s6,
	"s7":   s7,
	"s8":   s8,
	"s9":   s9,
	"s10":  s10,
	"s11":  s11,
	"t3":   t3,
	"t4":   t4,
	"t5":   t5,
	"t6":   t6,
}

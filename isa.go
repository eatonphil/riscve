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
)

type opcode struct {
	repr   string
	itype  itype
	opcode int
	func3  int
	func7  int
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
	},
	{
		"auipc",
		uType,
		b("0010111"),
		0,
		0,
	},
	{
		"jal",
		jType,
		b("1101111"),
		0,
		0,
	},
	{
		"jalr",
		iType,
		b("1100111"),
		0,
		0,
	},
	{
		"beq",
		bType,
		b("1100011"),
		0,
		0,
	},
	{
		"bne",
		bType,
		b("1100011"),
		b("001"),
		0,
	},
	{
		"blt",
		bType,
		b("1100011"),
		b("100"),
		0,
	},
	{
		"bge",
		bType,
		b("1100011"),
		b("101"),
		0,
	},
	{
		"bltu",
		bType,
		b("1100011"),
		b("110"),
		0,
	},
	{
		"bgeu",
		bType,
		b("1100011"),
		b("111"),
		0,
	},
	{
		"lb",
		iType,
		b("0000011"),
		0,
		0,
	},
	{
		"lh",
		iType,
		b("0000011"),
		b("001"),
		0,
	},
	{
		"lw",
		iType,
		b("0000011"),
		b("010"),
		0,
	},
	{
		"lbu",
		iType,
		b("0000011"),
		b("100"),
		0,
	},
	{
		"lhu",
		iType,
		b("0000011"),
		b("101"),
		0,
	},
	{
		"sb",
		sType,
		b("0100011"),
		0,
		0,
	},
	{
		"sh",
		sType,
		b("0100011"),
		b("001"),
		0,
	},
	{
		"sw",
		sType,
		b("0100011"),
		b("010"),
		0,
	},
	{
		"addi",
		iType,
		b("0010011"),
		0,
		0,
	},
	{
		"slti",
		iType,
		b("0010011"),
		b("010"),
		0,
	},
	{
		"sltiu",
		iType,
		b("0010011"),
		b("011"),
		0,
	},
	{
		"xori",
		iType,
		b("0010011"),
		b("100"),
		0,
	},
	{
		"ori",
		iType,
		b("0010011"),
		b("110"),
		0,
	},
	{
		"andi",
		iType,
		b("0010011"),
		b("111"),
		0,
	},
	{
		"slli",
		rType,
		b("0010011"),
		b("001"),
		0,
	},
	{
		"srli",
		rType,
		b("0010011"),
		b("101"),
		0,
	},
	{
		"srai",
		rType,
		b("0010011"),
		b("101"),
		b("0100000"),
	},
	{
		"add",
		rType,
		b("0110011"),
		0,
		0,
	},
	{
		"sub",
		rType,
		b("0110011"),
		0,
		b("0100000"),
	},
	{
		"sll",
		rType,
		b("0110011"),
		b("001"),
		0,
	},
	{
		"slt",
		rType,
		b("0110011"),
		b("010"),
		0,
	},
	{
		"sltu",
		rType,
		b("0110011"),
		b("011"),
		0,
	},
	{
		"xor",
		rType,
		b("0110011"),
		b("100"),
		0,
	},
	{
		"srl",
		rType,
		b("0110011"),
		b("101"),
		0,
	},
	{
		"sra",
		rType,
		b("0110011"),
		b("101"),
		b("0100000"),
	},
	{
		"or",
		rType,
		b("0110011"),
		b("110"),
		0,
	},
	{
		"and",
		rType,
		b("0110011"),
		b("111"),
		0,
	},
	{
		"fence",
		iType,
		b("0001111"),
		0,
		0,
	},
	{
		"ecall",
		iType,
		b("1110011"),
		0,
		0,
	},
	{
		"ebreak",
		iType,
		b("1110011"),
		0,
		1,
	},
}

func opcodeIndex(repr string) int {
	for i, oc := range opcodes {
		if oc.repr == repr {
			return i
		}
	}

	return -1
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

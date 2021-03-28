package instruction

import . "nes6502/cpu/addressing"

//InstructionSet
type Instruction struct{
	Bytes       int
	Cycle       int
	Description string
	MicroInstruction
	AddressMode
}

var Opcodes map[byte]Instruction = map[byte]Instruction{
	0x69: Instruction{
		2, 2, "ADC #oper",ADC, Immediate,
	},0x65: Instruction{
		2, 3, "ADC oper",ADC, Zeropage,
	},0x75: Instruction{
		2, 4, "ADC oper,X",ADC, ZeropageX,
	},0x6D: Instruction{
		3, 4, "ADC oper",ADC, Absolute,
	},0x7D: Instruction{
		3, 4, "ADC oper,X",ADC, AbsoluteX,
	},0x79: Instruction{
		3, 4, "ADC oper,Y",ADC, AbsoluteY,
	},0x61: Instruction{
		2, 6, "ADC (oper,X)",ADC, IndirectX,
	},0x71: Instruction{
		2, 5, "ADC (oper),Y",ADC, IndirectY,
	},0x29: Instruction{
		2, 2, "AND #oper",AND, Immediate,
	},0x25: Instruction{
		2, 3, "AND oper",AND, Zeropage,
	},0x35: Instruction{
		2, 4, "AND oper,X",AND, ZeropageX,
	},0x2D: Instruction{
		3, 4, "AND oper",AND, Absolute,
	},0x3D: Instruction{
		3, 4, "AND oper,X",AND, AbsoluteX,
	},0x39: Instruction{
		3, 4, "AND oper,Y",AND, AbsoluteY,
	},0x21: Instruction{
		2, 6, "AND (oper,X)",AND, IndirectX,
	},0x31: Instruction{
		2, 5, "AND (oper),Y",AND, IndirectY,
	},0x0A: Instruction{
		1, 2, "ASL A",ASL, Accumulator,
	},0x06: Instruction{
		2, 5, "ASL oper",ASL, Zeropage,
	},0x16: Instruction{
		2, 6, "ASL oper,X",ASL, ZeropageX,
	},0x0E: Instruction{
		3, 6, "ASL oper",ASL, Absolute,
	},0x1E: Instruction{
		3, 7, "ASL oper,X",ASL, AbsoluteX,
	},0x90: Instruction{
		2, 2, "BCC oper",BCC, Relative,
	},0xB0: Instruction{
		2, 2, "BCS oper",BCS, Relative,
	},0xF0: Instruction{
		2, 2, "BEQ oper",BEQ, Relative,
	},0x24: Instruction{
		2, 3, "BIT oper",BIT, Zeropage,
	},0x2C: Instruction{
		3, 4, "BIT oper",BIT, Absolute,
	},0x30: Instruction{
		2, 2, "BMI oper",BMI, Relative,
	},0xD0: Instruction{
		2, 2, "BNE oper",BNE, Relative,
	},0x10: Instruction{
		2, 2, "BPL oper",BPL, Relative,
	},0x00: Instruction{
		1, 7, "BRK",BRK, Implied,
	},0x50: Instruction{
		2, 2, "BVC oper",BVC, Relative,
	},0x70: Instruction{
		2, 2, "BVC oper",BVS, Relative,
	},0x18: Instruction{
		1, 2, "CLC",CLC, Implied,
	},0xD8: Instruction{
		1, 2, "CLD",CLD, Implied,
	},0x58: Instruction{
		1, 2, "CLI",CLI, Implied,
	},0xB8: Instruction{
		1, 2, "CLV",CLV, Implied,
	},0xC9: Instruction{
		2, 2, "CMP #oper",CMP, Immediate,
	},0xC5: Instruction{
		2, 3, "CMP oper",CMP, Zeropage,
	},0xD5: Instruction{
		2, 4, "CMP oper,X",CMP, ZeropageX,
	},0xCD: Instruction{
		3, 4, "CMP oper",CMP, Absolute,
	},0xDD: Instruction{
		3, 4, "CMP oper,X",CMP, AbsoluteX,
	},0xD9: Instruction{
		3, 4, "CMP oper,Y",CMP, AbsoluteY,
	},0xC1: Instruction{
		2, 6, "CMP (oper,X)",CMP, IndirectX,
	},0xD1: Instruction{
		2, 5, "CMP (oper),Y",CMP, IndirectY,
	},0xE0: Instruction{
		2, 2, "CPX #oper",CPX, Immediate,
	},0xE4: Instruction{
		2, 3, "CPX oper",CPX, Zeropage,
	},0xEC: Instruction{
		3, 4, "CPX oper",CPX, Absolute,
	},0xC0: Instruction{
		2, 2, "CPY #oper",CPY, Immediate,
	},0xC4: Instruction{
		2, 3, "CPY oper",CPY, Zeropage,
	},0xCC: Instruction{
		3, 4, "CPY oper",CPY, Absolute,
	},0xC6: Instruction{
		2, 5, "DEC oper",DEC, Zeropage,
	},0xD6: Instruction{
		2, 6, "DEC oper,X",DEC, ZeropageX,
	},0xCE: Instruction{
		3, 6, "DEC oper",DEC, Absolute,
	},0xDE: Instruction{
		3, 7, "DEC oper,X",DEC, AbsoluteX,
	},0xCA: Instruction{
		1, 2, "DEC",DEX, Implied,
	},0x88: Instruction{
		1, 2, "DEC",DEY, Implied,
	},0x49: Instruction{
		2, 2, "EOR #oper",EOR, Immediate,
	},0x45: Instruction{
		2, 3, "EOR oper",EOR, Zeropage,
	},0x55: Instruction{
		2, 4, "EOR oper,X",EOR, ZeropageX,
	},0x4D: Instruction{
		3, 4, "EOR oper",EOR, Absolute,
	},0x5D: Instruction{
		3, 4, "EOR oper,X",EOR, AbsoluteX,
	},0x59: Instruction{
		3, 4, "EOR oper,Y",EOR, AbsoluteY,
	},0x41: Instruction{
		2, 6, "EOR (oper,X)",EOR, IndirectX,
	},0x51: Instruction{
		2, 5, "EOR (oper),Y",EOR, IndirectY,
	},0xE6: Instruction{
		2, 5, "INC oper",INC, Zeropage,
	},0xF6: Instruction{
		2, 6, "INC oper,X",INC, ZeropageX,
	},0xEE: Instruction{
		3, 6, "INC oper",INC, Absolute,
	},0xFE: Instruction{
		3, 7, "INC oper,X",INC, AbsoluteX,
	},0xE8: Instruction{
		1, 2, "INX",INX, Implied,
	},0xC8: Instruction{
		1, 2, "INY",INY, Implied,
	},0x4C: Instruction{
		3, 3, "JMP oper",JMP, Absolute,
	},0x6C: Instruction{
		3, 5, "JMP (oper)",JMP, Indirect,
	},0x20: Instruction{
		3, 6, "JSR oper",JSR, Absolute,
	},0xA9: Instruction{
		2, 2, "LDA #oper",LDA, Immediate,
	},0xA5: Instruction{
		2, 3, "LDA oper",LDA, Zeropage,
	},0xB5: Instruction{
		2, 4, "LDA oper,X",LDA, ZeropageX,
	},0xAD: Instruction{
		3, 4, "LDA oper",LDA, Absolute,
	},0xBD: Instruction{
		3, 4, "LDA oper,X",LDA, AbsoluteX,
	},0xB9: Instruction{
		3, 4, "LDA oper,Y",LDA, AbsoluteY,
	},0xA1: Instruction{
		2, 6, "LDA (oper,X)",LDA, IndirectX,
	},0xB1: Instruction{
		2, 5, "LDA (oper),Y",LDA, IndirectY,
	},0xA2: Instruction{
		2, 2, "LDX #oper",LDX, Immediate,
	},0xA6: Instruction{
		2, 3, "LDX oper",LDX, Zeropage,
	},0xB6: Instruction{
		2, 4, "LDX oper,Y",LDX, ZeropageY,
	},0xAE: Instruction{
		3, 4, "LDX oper",LDX, Absolute,
	},0xBE: Instruction{
		3, 4, "LDX oper,Y",LDX, AbsoluteY,
	},0xA0: Instruction{
		2, 2, "LDY #oper",LDY, Immediate,
	},0xA4: Instruction{
		2, 3, "LDY oper",LDY, Zeropage,
	},0xB4: Instruction{
		2, 4, "LDY oper,X",LDY, ZeropageX,
	},0xAC: Instruction{
		3, 4, "LDY oper",LDY, Absolute,
	},0xBC: Instruction{
		3, 4, "LDY oper,X",LDY, AbsoluteX,
	},0x4A: Instruction{
		1, 2, "LSR A",LSR, Accumulator,
	},0x46: Instruction{
		2, 5, "LSR oper",LSR, Zeropage,
	},0x56: Instruction{
		2, 6, "LSR oper,X",LSR, ZeropageX,
	},0x4E: Instruction{
		3, 6, "LSR oper",LSR, Absolute,
	},0x5E: Instruction{
		3, 7, "LSR oper,X",LSR, AbsoluteX,
	},0xEA: Instruction{
		1, 2, "NOP",NOP, Implied,
	},0x09: Instruction{
		2, 2, "ORA #oper",ORA, Immediate,
	},0x05: Instruction{
		2, 3, "ORA oper",ORA, Zeropage,
	},0x15: Instruction{
		2, 4, "ORA oper,X",ORA, ZeropageX,
	},0x0D: Instruction{
		3, 4, "ORA oper",ORA, Absolute,
	},0x1D: Instruction{
		3, 4, "ORA oper,X",ORA, AbsoluteX,
	},0x19: Instruction{
		3, 4, "ORA oper,Y",ORA, AbsoluteY,
	},0x01: Instruction{
		2, 6, "ORA (oper,X)",ORA, IndirectX,
	},0x11: Instruction{
		2, 5, "ORA (oper),Y",ORA, IndirectY,
	},0x48: Instruction{
		1, 3, "PHA",PHA, Implied,
	},0x08: Instruction{
		1, 3, "PHP",PHP, Implied,
	},0x68: Instruction{
		1, 4, "PLA",PLA, Implied,
	},0x28: Instruction{
		1, 4, "PLP",PLP, Implied,
	},0x2A: Instruction{
		1, 2, "ROL A",ROL, Accumulator,
	},0x26: Instruction{
		2, 5, "ROL oper",ROL, Zeropage,
	},0x36: Instruction{
		2, 6, "ROL oper,X",ROL, ZeropageX,
	},0x2E: Instruction{
		3, 6, "ROL oper",ROL, Absolute,
	},0x3E: Instruction{
		3, 7, "ROL oper,X",ROL, AbsoluteX,
	},0x6A: Instruction{
		1, 2, "ROR A",ROR, Accumulator,
	},0x66: Instruction{
		2, 5, "ROR oper",ROR, Zeropage,
	},0x76: Instruction{
		2, 6, "ROR oper,X",ROR, ZeropageX,
	},0x6E: Instruction{
		3, 6, "ROR oper",ROR, Absolute,
	},0x7E: Instruction{
		3, 7, "ROR oper,X",ROR, AbsoluteX,
	},0x40: Instruction{
		1, 6, "RTI",RTI, Implied,
	},0x60: Instruction{
		1, 6, "RTS",RTS, Implied,
	},0xE9: Instruction{
		2, 2, "SBC #oper",SBC, Immediate,
	},0xE5: Instruction{
		2, 3, "SBC oper",SBC, Zeropage,
	},0xF5: Instruction{
		2, 4, "SBC oper,X",SBC, ZeropageX,
	},0xED: Instruction{
		3, 4, "SBC oper",SBC, Absolute,
	},0xFD: Instruction{
		3, 4, "SBC oper,X",SBC, AbsoluteX,
	},0xF9: Instruction{
		3, 4, "SBC oper,Y",SBC, AbsoluteY,
	},0xE1: Instruction{
		2, 6, "SBC (oper,X)",SBC, IndirectX,
	},0xF1: Instruction{
		2, 5, "SBC (oper),Y",SBC, IndirectY,
	},0x38: Instruction{
		1, 2, "SEC",SEC, Implied,
	},0xF8: Instruction{
		1, 2, "SED",SED, Implied,
	},0x78: Instruction{
		1, 2, "SEI",SEI, Implied,
	},0x85: Instruction{
		2, 3, "STA oper",STA, Zeropage,
	},0x95: Instruction{
		2, 4, "STA oper,X",STA, ZeropageX,
	},0x8D: Instruction{
		3, 4, "STA oper",STA, Absolute,
	},0x9D: Instruction{
		3, 5, "STA oper,X",STA, AbsoluteX,
	},0x99: Instruction{
		3, 5, "STA oper,Y",STA, AbsoluteY,
	},0x81: Instruction{
		2, 6, "STA (oper,X)",STA, IndirectX,
	},0x91: Instruction{
		2, 6, "STA (oper),Y",STA, IndirectY,
	},0x86: Instruction{
		2, 3, "STX oper",STX, Zeropage,
	},0x96: Instruction{
		2, 4, "STX oper,Y",STX, ZeropageY,
	},0x8E: Instruction{
		3, 4, "STX oper",STX, Absolute,
	},0x84: Instruction{
		2, 3, "STY oper",STY, Zeropage,
	},0x94: Instruction{
		2, 4, "STY oper,X",STY, ZeropageX,
	},0x8C: Instruction{
		3, 4, "STY oper",STY, Absolute,
	},0xAA: Instruction{
		1, 2, "TAX",TAX, Implied,
	},0xA8: Instruction{
		1, 2, "TAY",TAY, Implied,
	},0xBA: Instruction{
		1, 2, "TSX",TSX, Implied,
	},0x8A: Instruction{
		1, 2, "TXA",TXA, Implied,
	},0x9A: Instruction{
		1, 2, "TXS",TXS, Implied,
	},0x98: Instruction{
		1, 2, "TYA",TYA, Implied,
	},
}

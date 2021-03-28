package instruction

import (
	. "nes6502/bus"
	. "nes6502/cpu/alu"
	. "nes6502/cpu/register"
)

type Executable interface {
	Execute()
}

type MicroInstruction struct {
	Name        string
	Description string
	Execute     func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU)
}

func jump(pc uint16, offset byte) uint16 {
	return uint16(int16(pc) + int16(offset))
}

var (
	ADC = MicroInstruction{
		"ADC",
		"Add Memory to Accumulator with Carry",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.AC = alu.Add(register.AC, resolved[0])
		},
	}
	AND = MicroInstruction{
		"AND",
		"AND Memory with Accumulator",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.AC = alu.And(register.AC, resolved[0])
		},
	}
	ASL = MicroInstruction{
		"ASL",
		"Shift Left One Bit (Memory or Accumulator)",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.AC = alu.ShiftLeft(resolved[0])
		},
	}
	BCC = MicroInstruction{
		"BCC",
		"Branch on Carry Clear",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			if !register.GetStatus(SR_FLAG_Carry) {
				register.PC = jump(register.PC, resolved[0])
			}
		},
	}
	BCS = MicroInstruction{
		"BCS",
		"Branch on Carry Set",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			if register.GetStatus(SR_FLAG_Carry) {
				register.PC = jump(register.PC, resolved[0])
			}
		},
	}
	BEQ = MicroInstruction{
		"BEQ",
		"Branch on Result Zero",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			if register.GetStatus(SR_FLAG_Zero) {
				register.PC = jump(register.PC, resolved[0])
			}
		},
	}
	BIT = MicroInstruction{
		"BIT",
		"Test Bits in Memory with Accumulator",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	BMI = MicroInstruction{
		"BMI",
		"Branch on Result Minus",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			if register.GetStatus(SR_FLAG_Negative) {
				register.PC = jump(register.PC, resolved[0])
			}
		},
	}
	BNE = MicroInstruction{
		"BNE",
		"Branch on Result not Zero",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			if !register.GetStatus(SR_FLAG_Zero) {
				register.PC = jump(register.PC, resolved[0])
			}
		},
	}
	BPL = MicroInstruction{
		"BPL",
		"Branch on Result Plus",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			if register.GetStatus(SR_FLAG_Negative) {
				register.PC = jump(register.PC, resolved[0])
			}
		},
	}
	BRK = MicroInstruction{
		"BRK",
		"Force Break",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.SetStatus(SR_FLAG_Interrupt, true)
			//Stack<=PC
			vector := <-bus.ReadWord(0xFFFE)
			register.PC = uint16(vector[0]<<8 + vector[1])
		},
	}
	BVC = MicroInstruction{
		"BVC",
		"Branch on Overflow Clear",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			if !register.GetStatus(SR_FLAG_Overflow) {
				register.PC = jump(register.PC, resolved[0])
			}
		},
	}
	BVS = MicroInstruction{
		"BVS",
		"Branch on Overflow Set",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			if register.GetStatus(SR_FLAG_Overflow) {
				register.PC = jump(register.PC, resolved[0])
			}
		},
	}
	CLC = MicroInstruction{
		"CLC",
		"Clear Carry Flag",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.SetStatus(SR_FLAG_Carry, false)
		},
	}
	CLD = MicroInstruction{
		"CLD",
		"Clear Decimal Mode",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.SetStatus(SR_FLAG_Decimal, false)
		},
	}
	CLI = MicroInstruction{
		"CLI",
		"Clear Interrupt Disable Bit",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.SetStatus(SR_FLAG_Interrupt, false)
		},
	}
	CLV = MicroInstruction{
		"CLV",
		"Clear Overflow Flag",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.SetStatus(SR_FLAG_Overflow, false)
		},
	}
	CMP = MicroInstruction{
		"CMP",
		"Compare Memory with Accumulator",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	CPX = MicroInstruction{
		"CPX",
		"Compare Memory and Index X",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	CPY = MicroInstruction{
		"CPY",
		"Compare Memory and Index Y",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	DEC = MicroInstruction{
		"DEC",
		"Decrement Memory by One",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
		},
	}
	DEX = MicroInstruction{
		"DEX",
		"Decrement Index X by One",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	DEY = MicroInstruction{
		"DEY",
		"Decrement Index Y by One",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	EOR = MicroInstruction{
		"EOR",
		"Exclusive-OR Memory with Accumulator",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	INC = MicroInstruction{
		"INC",
		"Increment Memory by One",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	INX = MicroInstruction{
		"INX",
		"Increment Index X by One",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.X += 1
		},
	}
	INY = MicroInstruction{
		"INY",
		"Increment Index Y by One",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.Y += 1
		},
	}
	JMP = MicroInstruction{
		"JMP",
		"Jump to New Location",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.PC = address
		},
	}
	JSR = MicroInstruction{
		"JSR",
		"Jump to New Location Saving Return Address",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			//stack<=pc
			register.PC = address
		},
	}
	LDA = MicroInstruction{
		"LDA",
		"Load Accumulator with Memory",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	LDX = MicroInstruction{
		"LDX",
		"Load Index X with Memory",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	LDY = MicroInstruction{
		"LDY",
		"Load Index Y with Memory",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	LSR = MicroInstruction{
		"LSR",
		"Shift One Bit Right (Memory or Accumulator)",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.AC = alu.ShiftRight(resolved[0])
		},
	}
	NOP = MicroInstruction{
		"NOP",
		"No Operation",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}

	ORA = MicroInstruction{
		"ORA",
		"OR Memory with Accumulator",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	PHA = MicroInstruction{
		"PHA",
		"Push Accumulator on Stack",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			bus.WriteByte(uint16(register.SP)+0x1000, make([]byte, register.AC))
		},
	}
	PHP = MicroInstruction{
		"PHP",
		"Push Processor Status on Stack",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.SP++
			bus.WriteByte(uint16(register.SP)+0x1000, make([]byte, register.SR))
		},
	}
	PLA = MicroInstruction{
		"PLA",
		"Pull Accumulator from Stack",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	PLP = MicroInstruction{
		"PLP",
		"Pull Processor Status from Stack",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	ROL = MicroInstruction{
		"ROL",
		"Rotate One Bit Left (Memory or Accumulator)",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	ROR = MicroInstruction{
		"ROR",
		"Rotate One Bit Right (Memory or Accumulator)",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	RTI = MicroInstruction{
		"RTI",
		"Return from Interrupt",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			bs := <-bus.ReadWord(0x1000 + uint16(register.SP))
			register.PC = uint16(bs[0]<<8 + bs[1])
		},
	}
	RTS = MicroInstruction{
		"RTS",
		"Return from Subroutine",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	SBC = MicroInstruction{
		"SBC",
		"Subtract Memory from Accumulator with Borrow",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	SEC = MicroInstruction{
		"SEC",
		"Set Carry Flag",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.SetStatus(SR_FLAG_Carry, true)
		},
	}
	SED = MicroInstruction{
		"SED",
		"Set Decimal Flag",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.SetStatus(SR_FLAG_Decimal, true)
		},
	}
	SEI = MicroInstruction{
		"SEI",
		"Set Interrupt Disable Status",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.SetStatus(SR_FLAG_Interrupt, true)
		},
	}
	STA = MicroInstruction{
		"STA",
		"Store Accumulator in Memory",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	STX = MicroInstruction{
		"STX",
		"Store Index X in Memory",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	STY = MicroInstruction{
		"STY",
		"Sore Index Y in Memory",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	TAX = MicroInstruction{
		"TAX",
		"Transfer Accumulator to Index X",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.X = register.AC
		},
	}
	TAY = MicroInstruction{
		"TAY",
		"Transfer Accumulator to Index Y",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.Y = register.AC
		},
	}
	TSX = MicroInstruction{
		"TSX",
		"Transfer Stack Pointer to Index X",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.X = register.SP
		},
	}
	TXA = MicroInstruction{
		"TXA",
		"Transfer Index X to Accumulator",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.AC = register.X
		},
	}
	TXS = MicroInstruction{
		"TXS",
		"Transfer Index X to Stack Register",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.SP = register.X
		},
	}
	TYA = MicroInstruction{
		"TYA",
		"Transfer Index Y to Accumulator",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.AC = register.Y
		},
	}
)

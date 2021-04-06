package instruction

import (
	. "nes6502/bus"
	. "nes6502/cpu/alu"
	. "nes6502/cpu/register"
	. "nes6502/misc"
)

type Executable interface {
	Execute()
}

type MicroInstruction struct {
	Name        string
	Description string
	Execute     func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU)
}

func pushPC(register *Register, bus *Bus) {
	push(register, bus, byte(register.PC>>8))
	push(register, bus, byte(register.PC))
}

func pullPC(register *Register, bus *Bus) uint16 {

	bs1 := pull(register, bus)
	bs0 := pull(register, bus)
	register.PC = uint16(bs0)<<8 + uint16(bs1)
	return register.PC
}

func push(register *Register, bus *Bus, data byte) {
	bus.WriteByte(uint16(register.SP)+0x0100, []byte{data})
	Console.Trace("Stack Push %04X=%02X\n", uint16(register.SP)+0x0100, data)
	register.SP--
}
func pull(register *Register, bus *Bus) byte {
	register.SP++
	top := (<-bus.ReadByte(uint16(register.SP) + 0x0100))[0]
	Console.Trace("Stack Pull %04X=%02X\n", uint16(register.SP)+0x0100, top)
	return top
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
				register.PC = address
			}
		},
	}
	BCS = MicroInstruction{
		"BCS",
		"Branch on Carry Set",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			if register.GetStatus(SR_FLAG_Carry) {
				register.PC = address
			}
		},
	}
	BEQ = MicroInstruction{
		"BEQ",
		"Branch on Result Zero",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			if register.GetStatus(SR_FLAG_Zero) {
				register.PC = address
			}
		},
	}
	BIT = MicroInstruction{
		"BIT",
		"Test Bits in Memory with Accumulator",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.SetStatus(SR_FLAG_Overflow, operand[0]>>6&0x1 == 1)
			register.SetStatus(SR_FLAG_Negative, operand[0]>>7&0x1 == 1)
			register.SetStatus(SR_FLAG_Zero, register.AC&operand[0] == 0)
		},
	}
	BMI = MicroInstruction{
		"BMI",
		"Branch on Result Minus",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			if register.GetStatus(SR_FLAG_Negative) {
				register.PC = address
			}
		},
	}
	BNE = MicroInstruction{
		"BNE",
		"Branch on Result not Zero",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			if !register.GetStatus(SR_FLAG_Zero) {
				register.PC = address
			}
		},
	}
	BPL = MicroInstruction{
		"BPL",
		"Branch on Result Plus",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			if register.GetStatus(SR_FLAG_Negative) {
				register.PC = address
			}
		},
	}
	BRK = MicroInstruction{
		"BRK",
		"Force Break",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {

			if !register.GetStatus(SR_FLAG_Interrupt) {
				//push PC+2, push SR

				push(register, bus, register.SR)
				pushPC(register, bus)

				vector := <-bus.ReadWord(0xFFFE)
				register.PC = uint16(vector[0])<<8 + uint16(vector[1])

				register.SetStatus(SR_FLAG_Interrupt, true)
			}

		},
	}
	BVC = MicroInstruction{
		"BVC",
		"Branch on Overflow Clear",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			if !register.GetStatus(SR_FLAG_Overflow) {
				register.PC = address
			}
		},
	}
	BVS = MicroInstruction{
		"BVS",
		"Branch on Overflow Set",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			if register.GetStatus(SR_FLAG_Overflow) {
				register.PC = address
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
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			alu.Sub(resolved[0], register.AC)
		},
	}
	CPX = MicroInstruction{
		"CPX",
		"Compare Memory and Index X",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			alu.Sub(resolved[0], register.X)
		},
	}
	CPY = MicroInstruction{
		"CPY",
		"Compare Memory and Index Y",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			alu.Sub(resolved[0], register.Y)
		},
	}
	DEC = MicroInstruction{
		"DEC",
		"Decrement Memory by One",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			r := alu.Sub((<-bus.ReadByte(address))[0], 1)
			bus.WriteByte(address, []byte{r})
		},
	}
	DEX = MicroInstruction{
		"DEX",
		"Decrement Index X by One",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.X = alu.Sub(register.X, 1)
		},
	}
	DEY = MicroInstruction{
		"DEY",
		"Decrement Index Y by One",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.Y = alu.Sub(register.Y, 1)
		},
	}
	EOR = MicroInstruction{
		"EOR",
		"Exclusive-OR Memory with Accumulator",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {},
	}
	INC = MicroInstruction{
		"INC",
		"Increment Memory by One",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			r := alu.Add(resolved[0], 1)
			bus.WriteByte(address, []byte{r})

		},
	}
	INX = MicroInstruction{
		"INX",
		"Increment Index X by One",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.X = alu.Add(register.X, 1)
		},
	}
	INY = MicroInstruction{
		"INY",
		"Increment Index Y by One",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.Y = alu.Add(register.Y, 1)
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
			pushPC(register, bus)
			register.PC = address
		},
	}
	LDA = MicroInstruction{
		"LDA",
		"Load Accumulator with Memory",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.AC = resolved[0]
			alu.NegativeOut(int8(register.AC))
			alu.ZeroOut(int8(register.AC))
		},
	}
	LDX = MicroInstruction{
		"LDX",
		"Load Index X with Memory",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.X = resolved[0]
			alu.NegativeOut(int8(register.X))
			alu.ZeroOut(int8(register.X))
		},
	}
	LDY = MicroInstruction{
		"LDY",
		"Load Index Y with Memory",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.Y = resolved[0]
			alu.NegativeOut(int8(register.Y))
			alu.ZeroOut(int8(register.Y))
		},
	}
	LSR = MicroInstruction{
		"LSR",
		"Shift One Bit Right (Memory or Accumulator)",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			r := alu.ShiftRight(resolved[0])

			if address != 0 {
				bus.WriteByte(address, []byte{r})
			} else {
				register.AC = r
			}
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
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.AC = alu.Or(register.AC, resolved[0])
		},
	}
	PHA = MicroInstruction{
		"PHA",
		"Push Accumulator on Stack",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			push(register, bus, register.AC)
		},
	}
	PHP = MicroInstruction{
		"PHP",
		"Push Processor Status on Stack",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			push(register, bus, register.SR)
		},
	}
	PLA = MicroInstruction{
		"PLA",
		"Pull Accumulator from Stack",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.AC = pull(register, bus)
			alu.NegativeOut(int8(register.AC))
			alu.ZeroOut(int8(register.AC))
		},
	}
	PLP = MicroInstruction{
		"PLP",
		"Pull Processor Status from Stack",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.SR = pull(register, bus)
		},
	}
	ROL = MicroInstruction{
		"ROL",
		"Rotate One Bit Left (Memory or Accumulator)",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			r := alu.RotateLeft(resolved[0])
			if address != 0 {
				bus.WriteByte(address, []byte{r})
			} else {
				register.AC = r
			}
		},
	}
	ROR = MicroInstruction{
		"ROR",
		"Rotate One Bit Right (Memory or Accumulator)",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			r := alu.RotateRight(resolved[0])
			if address != 0 {
				bus.WriteByte(address, []byte{r})
			} else {
				register.AC = r
			}
		},
	}
	RTI = MicroInstruction{
		"RTI",
		"Return from Interrupt",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.SP = pull(register, bus)
			pullPC(register, bus)

		},
	}
	RTS = MicroInstruction{
		"RTS",
		"Return from Subroutine",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			//Pull PC
			pullPC(register, bus)
		},
	}
	SBC = MicroInstruction{
		"SBC",
		"Subtract Memory from Accumulator with Borrow",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.AC = alu.Sub(register.AC, resolved[0])
		},
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
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			bus.WriteByte(address, []byte{register.AC})
		},
	}
	STX = MicroInstruction{
		"STX",
		"Store Index X in Memory",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			bus.WriteByte(address, []byte{register.X})
		},
	}
	STY = MicroInstruction{
		"STY",
		"Sore Index Y in Memory",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			bus.WriteByte(address, []byte{register.Y})
		},
	}
	TAX = MicroInstruction{
		"TAX",
		"Transfer Accumulator to Index X",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.X = register.AC
			alu.NegativeOut(int8(register.X))
			alu.ZeroOut(int8(register.X))
		},
	}
	TAY = MicroInstruction{
		"TAY",
		"Transfer Accumulator to Index Y",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.Y = register.AC
			alu.NegativeOut(int8(register.Y))
			alu.ZeroOut(int8(register.Y))
		},
	}
	TSX = MicroInstruction{
		"TSX",
		"Transfer Stack Pointer to Index X",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.X = register.SP
			alu.NegativeOut(int8(register.X))
			alu.ZeroOut(int8(register.X))
		},
	}
	TXA = MicroInstruction{
		"TXA",
		"Transfer Index X to Accumulator",
		func(operand []byte, address uint16, resolved []byte, bus *Bus, register *Register, alu *ALU) {
			register.AC = register.X
			alu.NegativeOut(int8(register.AC))
			alu.ZeroOut(int8(register.AC))
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
			alu.NegativeOut(int8(register.AC))
			alu.ZeroOut(int8(register.AC))
		},
	}
)

package register

import "fmt"

const (
	SR_FLAG_Negative  = 1 << 7
	SR_FLAG_Overflow  = 1 << 6
	SR_FLAG_Ignored   = 1 << 5
	SR_FLAG_Break     = 1 << 4
	SR_FLAG_Decimal   = 1 << 3
	SR_FLAG_Interrupt = 1 << 2
	SR_FLAG_Zero      = 1 << 1
	SR_FLAG_Carry     = 1
)

type Register struct {
	//Registers:
	PC uint16 //BigEndian
	AC uint8
	X  uint8
	Y  uint8
	SR uint8
	SP uint8
}

func (register *Register) Reset() {
	register.PC = 0x0800 //program counter	(16 bit)
	register.AC = 0x00   //accumulator	(8 bit)
	register.X = 0x00    //X register	(8 bit)
	register.Y = 0x00    //Y register	(8 bit)
	register.SR = 0x00   //status register [NV-BDIZC]	(8 bit)
	register.SP = 0x00   //stack pointer	(8 bit)
}

func (register *Register) PrintStatus() {
	r := fmt.Sprintf("%08b", register.SR)
	fmt.Printf("status: %s\n", r)
}
func (register *Register) GetStatus(flag byte) bool {
	return register.SR&flag == flag
}
func (register *Register) SetStatus(flag byte, value bool) {

	/**
	  SR Flags (bit 7 to bit 0):

	  N	Negative
	  V	Overflow
	  -	ignored
	  B	Break
	  D	Decimal (use BCD for arithmetics)
	  I	Interrupt (IRQ disable)
	  Z	Zero
	  C	Carry
	*/
	if value {
		register.SR = register.SR | flag
	} else {
		register.SR = register.SR & (0xFF ^ flag)
	}
}

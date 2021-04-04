package cpu

import (
	"fmt"
	. "nes6502/bus"
	. "nes6502/cpu/alu"
	. "nes6502/cpu/instruction"
	. "nes6502/cpu/register"
)

type InstructionCycle struct {
	phase           chan int
	instructionCode byte
	instruction     Instruction
}

const (
	PhaseFetch       = 0
	PhaseTranslation = 1
	PhaseExecution   = 2
)

func (instructionCycle *InstructionCycle) print() {
	println("Cycle [%i]", instructionCycle.instructionCode)
}

type CPU struct {
	register *Register

	bus *Bus

	//Processor Stack:
	stack            []uint8
	cycle            int
	opcodes          map[byte]Instruction
	instructionCycle InstructionCycle
	alu              *ALU
}

func (cpu *CPU) changePhase(phase int) {
	cpu.instructionCycle.phase <- phase
}

func (cpu *CPU) runInstructionCycle(cycle chan int) {

	for phase := range cpu.instructionCycle.phase {
		switch phase {
		case PhaseFetch:
			{
				cpu.instructionCycle.instructionCode = cpu.fetch(cycle)
				go cpu.changePhase(PhaseTranslation)
			}
		case PhaseTranslation:
			{
				cpu.instructionCycle.instruction = cpu.translate(cpu.instructionCycle.instructionCode)
				go cpu.changePhase(PhaseExecution)
			}
		case PhaseExecution:
			{
				cpu.execute(cpu.instructionCycle.instruction, cycle)
				go cpu.changePhase(PhaseFetch)
			}
		}
	}
}

/**
opr executing phase
*/
func (cpu *CPU) execute(instruction Instruction, cycle chan int) {

	operandLen := instruction.Bytes - 1
	//fmt.Printf("operand length :%d\t", operandLen)
	operand := <-cpu.bus.Read(cpu.register.PC-uint16(instruction.Bytes)+1, operandLen)

	data, address := instruction.Resolve(operand, *cpu.bus, *cpu.register)
	fmt.Printf("%v\t0x%04X,%04X \n", operand, address, data)

	instruction.Execute(operand, address, []byte{data}, cpu.bus, cpu.register, cpu.alu)

	//cpu.register.PrintStatus()
	var tick = 0
	for tick < instruction.Cycle {
		tick++
		cpu.cycle = <-cycle
	}

	//fmt.Printf("execute cycles: %d, Total cycle: %d \n", instruction.Cycle, cpu.cycle)
}

func (cpu *CPU) fetch(cycle chan int) byte {
	//fetch next instruction from ram
	//instruction := <-cpu.instructionCycle.instruction
	cpu.cycle = <-cycle

	//fmt.Printf("Fetch PC: %X\t", cpu.register.PC)
	data := <-cpu.bus.ReadByte(cpu.register.PC)

	//fmt.Printf("OPCODE is : %X\t", data)
	return data[0]
}

func (cpu *CPU) translate(opCode byte) Instruction {
	instruction := cpu.opcodes[opCode]
	fmt.Printf("%04X\t%s\t", cpu.register.PC, instruction.Name)
	cpu.register.PC += uint16(instruction.Bytes)
	return instruction
}

func (cpu *CPU) Init() {

	//load instruction set
	cpu.opcodes = Opcodes
	// reset registers
	cpu.register = &Register{}
	cpu.alu = &ALU{}
	cpu.alu.Init(cpu.register)
	// define ram

	cpu.instructionCycle.phase = make(chan int)
	go cpu.changePhase(0)

	// Init stack
	// connect buses between ram and cpu

	//OAM
	//PPU
	//DMA

}

func (cpu *CPU) AcceptClockPulse(pules chan int) {

	cycle := make(chan int)

	go cpu.runInstructionCycle(cycle)
	for {
		cyclesPerPules := <-pules
		for i := 0; i < cyclesPerPules; i++ {
			cycle <- 1
		}
	}
}

func (cpu *CPU) Connect(bus *Bus) {
	cpu.bus = bus
}

func (cpu *CPU) Reset() {
	resetVector := <-cpu.bus.ReadWord(0xFFFC)
	fmt.Printf("Reset Vector: %X\n", uint16(resetVector[0])<<8+uint16(resetVector[1]))
	cpu.register.PC = uint16(resetVector[0])<<8 + uint16(resetVector[1])
}

package nes

import (
	. "nes6502/clock"
	. "nes6502/cpu"
	. "nes6502/ppu"
	"os"
)

type NES struct{
	//init Clock
	clock  *Clock
	cpu *CPU
	ppu *PPU
}

const BaseClockDivision = 3

func (nes *NES) PowerOn() {

	//init Clock
	clock := &Clock{}
	clock.SetFrequency(5.37, MHZ)

	cpuCycle := make(chan int, 5360520)
	clock.FrequencyDivision(cpuCycle, BaseClockDivision)

	ppuCycle := make(chan int, 5360520)
	clock.FrequencyDivision(ppuCycle, BaseClockDivision * 3)

	cpu := &CPU{}
	cpu.Init()

	ppu := &PPU{}
	ppu.Init()

	cpu.AcceptClockPulse(cpuCycle)
	ppu.AcceptClockPulse(ppuCycle)

	//SYNC CPU & CPU
	go clock.StartTick()

	println("startTick")


}

func (nes *NES) Reset() {

}

func (nes *NES) PowerOff() {
	os.Exit(0)
}



package nes

import (
	. "nes6502/clock"
	. "nes6502/cpu"
	. "nes6502/ppu"
	"os"
)

type NES struct {
	//init Clock
	clock     *Clock
	cpu       *CPU
	ppu       *PPU
	powerFlag bool
}

const (
	BaseClockDivision = 3
	PulseBuffer       = 5360520
)

func (nes *NES) PowerOn() {

	if nes.powerFlag {
		return
	}
	//init Clock
	clock := &Clock{}
	clock.SetFrequency(5.37, MHZ)

	cpuCycle := make(chan int, PulseBuffer)
	clock.FrequencyDivision(cpuCycle, BaseClockDivision)

	ppuCycle := make(chan int, PulseBuffer)
	clock.FrequencyDivision(ppuCycle, BaseClockDivision*3)

	cpu := &CPU{}
	cpu.Init()

	ppu := &PPU{}
	ppu.Init()

	cpu.AcceptClockPulse(cpuCycle)
	ppu.AcceptClockPulse(ppuCycle)

	//SYNC CPU & CPU
	go clock.StartTick()

	println("startTick")

	//TODO wait for power off signal
	//TODO wait for reset signal

}

func (nes *NES) Reset() {
	if nes.powerFlag {
		//RESET interrupt
	}
}

func (nes *NES) PowerOff() {
	os.Exit(0)
}

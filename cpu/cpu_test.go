package cpu

import (
	"io/ioutil"
	. "nes6502/bus"
	. "nes6502/clock"
	. "nes6502/io"
	. "nes6502/ppu"
	. "nes6502/ram"
	. "nes6502/rom"
	"testing"
)

func TestPowerOn(t *testing.T) {

	clock := &Clock{}
	clock.SetFrequency(5.37, MHZ)

	cpuCycle := make(chan int, 5360520)
	clock.FrequencyDivision(cpuCycle, 3)

	u := CPU{}
	u.Init()

	go u.AcceptClockPulse(cpuCycle)

	//fs,_ := ioutil.ReadFile("../rom/bin/start")
	//println(err.Error())

	r := &Resolver{}
	/**
	01-abs_x_wrap
	-------------
	Verifies that $FFFF wraps around to 0 for STA abs,X and LDA abs,X.
	*/
	romFile, _ := ioutil.ReadFile("../test/cpu/01-abs_x_wrap.nes")
	rom := r.Resolve(romFile)

	ram := &Ram{}

	//ram.Dump(fs)

	bus := &Bus{}

	ppuRegisters := &PPURegister{}
	ppuRegisters.Allocate(8)

	sram := &SRam{}
	dma := &DMA{}
	ioRegisters := NewIORegisters()
	bus.Init(&RamMapper{Ram: ram, PpuRegister: ppuRegisters, IoRegister: &ioRegisters, Rom: rom, SRam: sram, Dma: dma})
	u.Connect(bus)

	u.Reset()

	clock.StartTick()

	t.Logf("TestPowerOn")
}

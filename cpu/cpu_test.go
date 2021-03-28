package cpu

import (
	"io/ioutil"
	. "nes6502/bus"
	. "nes6502/clock"
	"nes6502/io"
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
	romfile,_ := ioutil.ReadFile("../rom/1.Branch_Basics.nes")
	rom := r.Resolve(romfile)//&Rom{}

	ram := &Ram{}

	//ram.Dump(fs)

	bus := &Bus{}

	ppuRegisters := &PPURegister{}
	ppuRegisters.Allocate(8)


	sram := &SRam{}
	dma := &DMA{}
	ioRegisters := &io.IORegisters{}
	bus.Init(&RamMapper{ram, ppuRegisters,ioRegisters,rom,sram, dma})
	u.Connect(bus)

	clock.StartTick()

	t.Logf("TestPowerOn")
}


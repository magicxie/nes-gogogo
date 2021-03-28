package bus

import (
	"fmt"
	. "nes6502/io"
	. "nes6502/ppu"
	. "nes6502/ppu/vram"
	. "nes6502/ram"
	. "nes6502/rom"
	"testing"
)

func TestBusRead(t *testing.T) {

	ram := &Ram{}
	ddd := []byte{0x4C, 0xF5, 0xC5, 0x60, 0x78, 0xD8, 0xA2, 0xFF, 0x9A, 0xAD, 0x02, 0x20, 0x10, 0xFB, 0xAD}
	for i, dd := range ddd {
		ram.WriteByte(uint16(i), dd)
	}

	bus := &Bus{}
	rom := &Rom{}
	sram := &SRam{}
	ppuRegisters := &PPURegister{}
	ppuRegisters.Allocate()

	bus.Init(&RamMapper{ram, ppuRegisters, rom, sram})

	for i := range ddd {
		d := <-bus.Read(uint16(i), 1)
		fmt.Printf("-> %X", d)
	}

}

func TestBusWrite(t *testing.T) {

	ram := &Ram{}
	ddd := []byte{0x4C, 0xF5, 0xC5, 0x60, 0x78, 0xD8, 0xA2, 0xFF, 0x9A, 0xAD, 0x02, 0x20, 0x10, 0xFB, 0xAD}

	bus := &Bus{}

	rom := &Rom{}
	sram := &SRam{}
	sram.Allocate(0x2000)

	ppuRegisters := &PPURegister{}
	ppuRegisters.Allocate(8)

	ioRegisters := &IORegister{}
	ioRegisters.Allocate(20)

	oam := &Oam{}
	oam.Allocate(256)

	dma := &DMA{ram, oam}
	bus.Init(&RamMapper{ram, ppuRegisters, ioRegisters,rom, sram, dma})

	for i := range ddd {
		bus.WriteWord(uint16(i*2), []byte{0x56, 0x78})
		fmt.Printf("%X -> %X\n", uint16(i*2), ram.ReadBytes(uint16(i*2), 2))
	}
	bus.WriteByte(0x2000, []byte{0xCF})
	fmt.Printf("%b\n", 0xCF)
	fmt.Printf("NameTable %d\n", ppuRegisters.NameTable())
	fmt.Printf("OAMMode %d\n", ppuRegisters.OAMMode())
	fmt.Printf("SpritePatternAddress %d\n", ppuRegisters.SpritePatternAddress())
	fmt.Printf("BackgroundPatternAddress %d\n", ppuRegisters.BackgroundPatternAddress())
	fmt.Printf("SpriteHeight %d\n", ppuRegisters.SpriteHeight())
	fmt.Printf("EnableNMI %v\n", ppuRegisters.EnableNMI())

	ppu := &PPU{}
	ppu.Oam = oam

	vbus := &Bus{}
	vbus.Init(&VRamMapper{&PatternTables{},
		&NameTables{},
		&Palette{},
		&Palette{}})

}

package rom

import (
	"fmt"
)

const (
	VERTICAL_MIRROR   = true
	HORIZONTAL_MIRROR = false
)

type Trainer struct {
	data []byte
}

type Program struct {
	data []byte
}

func (prg *Program) ReadBytes(address uint16, bytes int) []byte {
	return prg.data[address : address+uint16(bytes)]
}

type Character struct {
	data []byte
}

func (chr *Character) Dump() []byte {
	return chr.data
}

type Header struct {
	Magic        string
	PGMMirrors   byte
	CHRMirrors   byte
	MirrorType   bool
	HasBattery   bool
	HasTrainer   bool
	ScanMode     bool
	Mapper       byte
	VSUnisystem  bool
	PlayChoice10 bool
	NESVersion   byte
	RamPGMSize   byte
	TVMODE0      string
	TVMODE1      string
	RamPGM       bool
}

type Rom struct {
	Header    *Header
	Trainer   *Trainer
	Program   *Program
	Character *Character
}

func (rom *Rom) Summary() {

	fmt.Printf("PGM Mirrors %d * 16KB\n", rom.Header.PGMMirrors)
	fmt.Printf("CHR Mirrors %d * 8KB\n", rom.Header.CHRMirrors)

	mirrorType := "horizontal"
	if rom.Header.MirrorType {
		mirrorType = "vertical"
	}
	fmt.Printf("MirrorType %s\n", mirrorType)
	fmt.Printf("HasBattery %v\n", rom.Header.HasBattery)
	fmt.Printf("HasTrainer %v\n", rom.Header.HasTrainer)
	fmt.Printf("ScanMode %v\n", rom.Header.ScanMode)
	fmt.Printf("HasTrainer %v\n", rom.Header.HasTrainer)
	fmt.Printf("Mapper %08X\n", rom.Header.Mapper)
	fmt.Printf("VSUnisystem %v\n", rom.Header.VSUnisystem)
	fmt.Printf("PlayChoice10 %v\n", rom.Header.PlayChoice10)
	fmt.Printf("NESVersion %v\n", rom.Header.NESVersion)
	fmt.Printf("RamPGMSize %v * 8K\n", rom.Header.RamPGMSize)
	fmt.Printf("TVMODE0 %s\n", rom.Header.TVMODE0)
	fmt.Printf("TVMODE1 %s\n", rom.Header.TVMODE1)
	fmt.Printf("RamPGM %v\n", rom.Header.RamPGM)

	fmt.Printf("Trainer size:%v\n", len(rom.Trainer.data))
	fmt.Printf("PGM size:%v\n", len(rom.Program.data))
	fmt.Printf("CHR size:%v\n\n", len(rom.Character.data))

}

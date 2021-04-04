package rom

import (
	"errors"
	"fmt"
)

type Resolver struct {
}

var INVALID_NES_FORMAT_ERROR = errors.New("invalid nes header")

func (resolver *Resolver) Resolve(romRawData []byte) (rom *Rom) {

	rom = &Rom{}

	romSize := len(romRawData)
	fmt.Printf("ROM size %fK\n", float32(romSize)/1024)

	_, rom.Header = resolver.resolveHeader(romRawData[0:16])
	rom.Trainer = &Trainer{}

	trainerSize := 512
	pos := 16
	if rom.Header.HasTrainer {
		rom.Trainer.data = romRawData[pos : pos+trainerSize]
		pos += trainerSize
	}

	program := &Program{}
	var pgmSize int = int(rom.Header.PGMMirrors) * 8192 * 2
	program.data = romRawData[pos : pos+pgmSize]
	pos += pgmSize
	rom.Program = program

	character := &Character{}
	var chrSize int = int(rom.Header.CHRMirrors) * 8192
	character.data = romRawData[pos : pos+chrSize]
	pos += chrSize
	rom.Character = character

	if pos < romSize {
		fmt.Printf("Has additional rom is %d to %d\n", pos, romSize)
	} else {
		fmt.Printf("ROM END\n")
	}
	return rom
}

func (resolver *Resolver) resolveHeader(headerData []byte) (err error, header *Header) {
	header = &Header{}

	nes := string(headerData[0:4])
	println("FILE HEAD", nes)

	if nes != MAGIC {
		err = INVALID_NES_FORMAT_ERROR
		return err, nil
	}

	header.Magic = nes
	header.PGMMirrors = headerData[4]
	header.CHRMirrors = headerData[5]

	mirrorFlags := headerData[6]
	header.MirrorType = mirrorFlags>>7 == 1
	header.HasBattery = mirrorFlags<<1>>7 == 1
	header.HasTrainer = mirrorFlags<<2>>7 == 1
	header.VSUnisystem = mirrorFlags<<3>>7 == 1
	header.Mapper = mirrorFlags & 0x0F

	systemFlags := headerData[7]
	header.ScanMode = systemFlags>>7 == 1
	header.PlayChoice10 = systemFlags<<1>>7 == 1
	header.NESVersion = systemFlags << 3 >> 7
	header.Mapper += (mirrorFlags & 0x0F) << 4

	header.RamPGMSize = headerData[8]
	tvMode0Flags := headerData[9]

	if tvMode0Flags>>7 == 0 {
		header.TVMODE0 = "NTSC"
	} else {
		header.TVMODE0 = "PAL"
	}

	tvMode1Flags := (headerData[10] & 192)

	if tvMode1Flags == 0 {
		header.TVMODE1 = "NTSC"
	} else if tvMode1Flags == 1 || tvMode1Flags == 3 {
		header.TVMODE1 = "PAL"
	} else {
		header.TVMODE1 = "PAL/NTSC"
	}

	header.RamPGM = tvMode1Flags<<4>>6 == 0

	return nil, header
}

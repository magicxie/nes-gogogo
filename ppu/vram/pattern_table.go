package vram

import "nes6502/io"

type PatternTable struct{

}

type PatternTables struct{
	io.Memory
	PatternTables *[]PatternTable
}

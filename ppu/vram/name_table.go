package vram

import "nes6502/io"

type AttributeTable struct{

}


type PhysicNameTable struct{

}

type NameTable struct{

	PhysicNameTable *PhysicNameTable
	AttributeTable *AttributeTable
}

type NameTables struct{
	io.Memory
	NameTables *[]NameTable
}

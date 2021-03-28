package ram

import (
	"testing"
)

func TestLittleEndian(t *testing.T) {
	ram := &Ram{}
	var data uint16 = 0xAFBD
	ram.WriteUint16(0, data)
	var write = ram.ReadUint16(0)
	if  write != data{
		t.Fatalf("readUint16 error, expect:%X, actual:%X", data, write)
	}

}

package rom

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	. "nes6502/cpu/instruction"
	"testing"
)

func Test_Rom(t *testing.T) {

	fs, _ := ioutil.ReadFile("../rom/bin/adcz")

	var bs = 0
	for i := 0; i < len(fs); i += bs {

		ins := Opcodes[fs[i]]

		bs = ins.Bytes
		if bs > 1{
			oprands := fs[i+1 : i+ins.Bytes]
			var v uint16 = uint16(oprands[0])
			if len(oprands) > 1 {
				v = binary.LittleEndian.Uint16(oprands)
			}
			fmt.Printf("%04x %s:%d 0x%04X\n", i, ins.Description, ins.Bytes-1, v)
			if ins.Name == "JMP"{
				fmt.Printf("jmp to %x", int(v-0x800))
				bs=0
				i = int(v-0x800)
			}
		}else{
			fmt.Printf("%04x %s:%d\n", i, ins.Name, ins.Bytes)
			if bs ==0 {
				i+=1
			}
		}

		//i += ins.Bytes
	}
}

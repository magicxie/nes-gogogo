package alu

/**
 * !!!Caution!!!: Binary Coded Decimal mode is not support on NES
 */
import (
	"fmt"
	"strconv"
)

type BCD struct{
}
func (bcd *BCD) EncodeDecimal(dec int) uint64{
	s := strconv.Itoa(dec)
	bits := len(s) * 4
	ui, _ := strconv.ParseUint(s, 16, bits)
	return ui
}

func (bcd *BCD) DecodeBcd(dec int) int{
	ui := fmt.Sprintf("%x", dec)
	s,_ := strconv.Atoi(ui)
	return s
}


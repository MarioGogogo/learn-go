package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var i8 int8
	var ui16 uint16
	var ui64 uint64

	fmt.Println("i8", unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	fmt.Println("ui16", unsafe.Sizeof(ui16), 0, math.MaxUint16)
	fmt.Println("ui64", unsafe.Sizeof(ui64), 0, uint64(math.MaxUint64))

	var ui uint
	ui = uint(math.MaxUint64) //再 +1 会导致 overflow 溢出

	fmt.Println("max", unsafe.Sizeof(ui), 0, ui)

	//var imax int
	//imax = int(math.MaxInt64)  //再 +1 会导致 overflow 溢出
	//imin1 = int(math.MinInt64) //再 -1 会导致 overflow 溢出

	var b int = 077
	fmt.Println("b", b)
	var b1 int = 0xff
	fmt.Println("b1", b1)
}

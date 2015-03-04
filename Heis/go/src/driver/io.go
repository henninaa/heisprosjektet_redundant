package driver // where "driver" is the folder that contains io.go, io.c, io.h, channels.go, channels.c and driver.go
/*
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
*/
import "C"

func io_init() bool {
	return bool(int(C.io_init()) != -1)
}

func io_set_bit(channel int){
	C.io_set_bit(C.int(channel)
} 

func io_clear_bit(channel int){
	C.io_clear_bit(C.int(channel))
}

func io_write_bit(channel int, value int){
	C.io_write_bit(C.int(channel),C.int(value))
}

func io_read_bit(channel int) bool {
	return bool(int(C.io_read_bit(C.int(channel))) != 0)
}

func Io_read_analog(channel int) int {
	return int(C.io_read_analog(C.int(channel)))
}
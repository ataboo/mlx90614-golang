package main

// #include "bcm2835.h"
// #include <stdlib.h>
import "C"
import (
	"errors"
	"log"
	"fmt"
	"unsafe"
)

func main() {
	if err := connect(); err != nil {
		log.Fatal("Failed to connect")
	}


}

func connect() error {
	defer C.bcm2835_i2c_end()

	if C.bcm2835_init() == 0 {
		return errors.New("Init: failed")
	}

	C.bcm2835_i2c_end()
	C.bcm2835_i2c_begin()
	C.bcm2835_i2c_setSlaveAddress(0x5a)

	buf := ""
	pBuf := unsafe.Pointer(C.CString(buf))
	defer C.free(pBuf)
	reg := string([]byte{0x07})
	pReg := unsafe.Pointer(C.CString(reg))
	defer C.free(pReg)
	C.bcm2835_i2c_read_register_rs((*C.char)(pReg), (*C.char)(pBuf), 2)

	fmt.Sprintf("\nRead: %+v", buf)

	return nil
}
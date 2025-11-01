package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var name uint8;
	fmt.Println(unsafe.Sizeof(name));
}
package main

import _ "unsafe"

//go:linkname printf _printf
func printf(format *int8, __llgo_va_list ...any)

var hello = [...]int8{'H', 'e', 'l', 'l', 'o', '\n', 0}

func main() {
	printf(&hello[0])
}
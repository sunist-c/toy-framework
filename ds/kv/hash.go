package kv

import "unsafe"

func hash[Key any](key Key) uint {
	gc := &key
	start := uintptr(unsafe.Pointer(gc))
	offset := unsafe.Sizeof(key)
	sizeOfByte := unsafe.Sizeof(byte(0))

	hashSum := uint(0)
	for ptr := start; ptr < start+offset; ptr += sizeOfByte {
		b := *(*byte)(unsafe.Pointer(ptr))
		hashSum += uint(b)
		hashSum = uint(b) + (hashSum << 6) + (hashSum << 16) - hashSum
	}
	return hashSum
}

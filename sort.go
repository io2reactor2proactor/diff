package main

import (
	"hash/crc32"
)

type Key []Item
func (a Key) Len() int           { return len(a) }
func (a Key) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Key) Less(i, j int) bool { return a[i].vir < a[j].vir }

const seed uint32 = 31 //31 131 1313 13131 131313 etc..
func old_bkdrHash(str string) uint32 {
    var h uint32
    for _, c := range str {
        h = h * seed + uint32(c)
    }
    return h
}

func bkdrHash(key string) uint32 {
    if len(key) < 64 {
        var scratch [64]byte 
        copy(scratch[:], key)
        return crc32.ChecksumIEEE(scratch[:len(key)])
    }
    return crc32.ChecksumIEEE([]byte(key))
}

package main

import (
    "bytes"
    "fmt"

    "github.com/zhuangsirui/binpacker"
)

type T struct {
    A uint16
    B string
    C []byte
}

func main() {
    field1 := uint16(1)
    field2 := "Hello World"
    field3 := []byte("Hello World")
    buffer := new(bytes.Buffer)
    binpacker.NewPacker(buffer).
        PushUint16(field1).
        PushUint16(uint16(len(field2))).PushString(field2).
        PushUint16(uint16(len(field3))).PushBytes(field3)

    t := new(T)
    unpacker := binpacker.NewUnpacker(buffer)
    unpacker.FetchUint16(&t.A).StringWithUint16Perfix(&t.B).BytesWithUint16Perfix(&t.C)
    fmt.Println(t," byte : ",string(t.C))
}

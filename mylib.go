package mylib

import (
	"bytes"
	"encoding/binary"
	"io"
)

func ReadUint16LE(r io.Reader) uint16 {
	var buf bytes.Buffer
	io.CopyN(&buf, r, 2)
	return binary.LittleEndian.Uint16(buf.Bytes())
}

func ReadUint32LE(r io.Reader) uint32 {
	var buf bytes.Buffer
	io.CopyN(&buf, r, 4)
	return binary.LittleEndian.Uint32(buf.Bytes())
}

func ReadUint64LE(r io.Reader) uint64 {
	var buf bytes.Buffer
	io.CopyN(&buf, r, 8)
	return binary.LittleEndian.Uint64(buf.Bytes())
}

func ReadUint16BE(r io.Reader) uint16 {
	var buf bytes.Buffer
	io.CopyN(&buf, r, 2)
	return binary.BigEndian.Uint16(buf.Bytes())
}

func ReadUint32BE(r io.Reader) uint32 {
	var buf bytes.Buffer
	io.CopyN(&buf, r, 4)
	return binary.BigEndian.Uint32(buf.Bytes())
}

func ReadUint64BE(r io.Reader) uint64 {
	var buf bytes.Buffer
	io.CopyN(&buf, r, 8)
	return binary.BigEndian.Uint64(buf.Bytes())
}

func Padding(d *bytes.Buffer, base uint32) uint32 {
	len := d.Len()
	p := base - (uint32(len) % base)
	if p != base {
		for i := 0; i < int(p); i++ {
			binary.Write(d, binary.LittleEndian, uint8(0))
		}
	} else {
		p = 0
	}
	return uint32(p)
}

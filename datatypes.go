package stargopher

import (
	"encoding/binary"
)

type WorldCoord [2]int64

type NetTile struct {
	Unknown     int16
	Unknown2    uint8
	Unknown3    uint8
	Unknown4    int16
	Unknown5    uint8
	Unknown6    int16
	Unknown7    uint8
	Unknown8    uint8
	Unknown9    int16
	Unknown10   uint8
	Unknown11   uint8
	Unknown12   uint8
	Unknown13   uint8
	Unknown14   uint8
	LiquidLevel uint8
	Gravity     SVLQ
}

//SVLQ is the variable length quantity type. Google for the schema info.
type SVLQ []byte

//VLQ is the variable length quantity type. Google for the schema info.
type VLQ []byte

//Decode returns an int64 from a SVLQ
func (v SVLQ) Decode() int64 {
	x, _ := binary.Varint(v)
	return x
}

//Decode returns an uint64 from a VLQ
func (v VLQ) Decode() uint64 {
	x, _ := binary.Uvarint(v)
	return x
}

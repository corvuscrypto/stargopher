package stargopher

//SVLQ is the variable length quantity type. Google for the schema info.
type SVLQ []byte

//VLQ is the variable length quantity type. Google for the schema info.
type VLQ []byte

//WorldCoord contains the X and Y coordinates of a world in Starbound
type WorldCoord [2]int64

//NetTile seems to hold information about game tiles
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

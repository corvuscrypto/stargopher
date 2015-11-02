package stargopher

//ProtocolVersion is the first packet sent. It contains the server version.
type ProtocolVersion struct {
	Version uint32
}

//ConnectionResponse tells the client whether their connection
//attempt is successful or if they have been rejected.
//It is the final packet sent in the handshake process.
type ConnectionResponse struct {
	Success             bool
	ClientID            VLQ
	RejectionReason     string
	CelestialInfoExists bool
	OrbitalLevels       int32
	ChunkSize           int32
	XYMax               int32
	ZMin                int32
	ZMax                int32
	NumSectors          VLQ
	SectorID            string
	SectorSeed          uint64
	SectorPrefix        string
	Parameters          interface{}
	SectorConfig        interface{}
}

//DisconnectResponse is used to notify the client of a disconnect.
type DisconnectResponse struct {
	Unknown uint8
}

//HandshakeChallenge provides a salt and round count for password
//verification. It is followed by Handshake Response.
type HandshakeChallenge struct {
	ClaimMessage string
	Salt         string
	HashCount    int32
}

//ChatReceivedPacket is sent to a client with a chat message.
type ChatReceivedPacket struct {
	UserId         int
	UserNameLength int
	UserName       string
	MessageLength  int
	Message        string
}

//UniverseTimeUpdate is sent from the server to update the current time.
type UniverseTimeUpdate struct {
	Time SVLQ
}

//CelestialResponse has yet to be fully understood.
type CelestialResponse struct {
	//Unknown
}

//ClientConnect  is sent in the handshake process immediately after the Protocol
//Version. It contains all relevant data about the connecting player.
type ClientConnect struct {
	AssetDigest []uint8
	Claim       interface{}
	UUIDFlag    bool
	UUID        [16]uint8
	PlayerName  string
	Species     string
	Shipworld   []uint8
	Account     string
}

//ClientDisconnect is sent when the client disconnects.
type ClientDisconnect struct {
	Unknown uint8
}

//HandshakeResponse is the response to the Handshake Challenge.
type HandshakeResponse struct {
	ClaimResponse string
	PasswordHash  string
}

//WarpCommand is sent when the player warps/is warped to a planet or ship.
type WarpCommand struct {
	WarpType    uint32
	WorldCoords WorldCoord
	PlayerName  string
}

//ChatSent is sent from the client whenever a message is sent in the chat window.
type ChatSent struct {
	Message string
	Channel uint8
}

//CelestialRequest has yet to be fully understood.
type CelestialRequest struct {
	//Unknown
}

//ClientContextUpdate has yet to be fully understood.
type ClientContextUpdate struct {
	Data []uint8
}

//WorldStart is sent to the client when a world thread has been started
//on the server.
type WorldStart struct {
	Planet          interface{}
	WorldStructure  interface{}
	Sky             []uint8
	ServerWeather   []uint8
	SpawnX          float64
	SpawnY          float64
	WorldProperties map[string]interface{}
	ClientID        uint32
	Local           bool
}

//WorldStop is called when a world thread is stopped.
type WorldStop struct {
	Status string
}

//TileArrayUpdate is called when an array of tiles has its properties updated.
type TileArrayUpdate struct {
	TileX  SVLQ
	TileY  SVLQ
	Width  VLQ
	Height VLQ
	Tiles  [][]NetTile
}

//TileUpdate is called when a tile is updated.
type TileUpdate struct {
	TileX SVLQ
	TileY SVLQ
	Tile  NetTile
}

//TileLiquidUpdate is sent when the liquid on a tile has changed position.
type TileLiquidUpdate struct {
	TileX       SVLQ
	TileY       SVLQ
	LiquidLevel uint8
	LiquidType  uint8
}

//TileDamageUpdate is sent when a tile is damaged.
type TileDamageUpdate struct {
	TileX                  int32
	TileY                  int32
	Unknown                uint8
	DamagePercentage       float64
	DamageEffectPercentage float64
	SourceX                float64
	SourceY                float64
	DamageType             uint8
}

//TileModificationFailure is sent when a tile list cannot successfully be modified.
type TileModificationFailure struct {
	//unknown
}

//GiveItem attempts to give an item to a player. If the player's
//inventory is full, it will drop on the ground next to them.
type GiveItem struct {
	ItemName       string
	Count          VLQ
	ItemProperties interface{}
}

//ContainerSwapResult is sent whenever two items are
// swapped in an open container.
type ContainerSwapResult struct {
	//unknown
}

//EnvironmentUpdate is sent on an environment update.
type EnvironmentUpdate struct {
	//unknown
}

//EntityInteractResult contains the results of an entity interaction.
type EntityInteractResult struct {
	ClientID uint32
	EntityID int32
	Results  interface{}
}

//ModifyTileList contains a list of tiles and modifications to them.
type ModifyTileList struct {
	//Unknown
}

//DamageTile updates a tile's damage
type DamageTile struct {
	//Unknown
}

//DamageTileGroup updates an entire tile group's damage.
type DamageTileGroup struct {
	//Unknown
}

//RequestDrop requests an item drop from the ground.
type RequestDrop struct {
	EntityID SVLQ
}

//SpawnEntity requests that the server spawn an entity.
type SpawnEntity struct {
	//Unknown
}

//EntityInteract is sent when a client attempts to interact with an entity.
type EntityInteract struct {
	//Unknown
}

//ConnectWire connects a wire.
type ConnectWire struct {
	//Unknown
}

//DisconnectAllWires disconnects all wires.
type DisconnectAllWires struct {
	//Unknown
}

//OpenContainer opens a container.
type OpenContainer struct {
	EntityID SVLQ
}

//CloseContainer closes a container.
type CloseContainer struct {
	EntityID SVLQ
}

//ContainerSwap swaps an item in a container.
type ContainerSwap struct {
	//Unknown
}

//ContainerItemApply applies an item to another item in a container.
type ContainerItemApply struct {
	//Unknown
}

//ContainerStartCrafting initiates crafting on an item in a container
//(Used in pixel compressors and the like?)
type ContainerStartCrafting struct {
	EntityID SVLQ
}

//ContainerStopCrafting packet stops crafting on an item in a container
type ContainerStopCrafting struct {
	EntityID SVLQ
}

//BurnContainer burns a container.
type BurnContainer struct {
	EntityID SVLQ
}

//ClearContainer clears a container.
type ClearContainer struct {
	EntityID SVLQ
}

//WorldClientStateUpdate contains a world client state update
type WorldClientStateUpdate struct {
	Delta []uint8
}

//EntityCreate creates an entity.
type EntityCreate struct {
	EntityType uint8
	StoreData  []uint8
	EntityID   SVLQ
}

//EntityUpdate updates an entity's properties.
type EntityUpdate struct {
	EntityID SVLQ
	Delta    []uint8
}

//EntityDestroy destroys an entity.
type EntityDestroy struct {
	EntityID SVLQ
	Death    bool
}

//DamageNotification notifies the receiver of damage received.
type DamageNotification struct {
	CausingEntityID    SVLQ
	TargetEntityID     SVLQ
	PosX               SVLQ //*
	PosY               SVLQ //*
	Damage             SVLQ //*
	DamageType         uint8
	DamageSourceType   string
	TargetMaterialType string
	HitResultType      uint8

	//* denotes that values must be divided by 100 before manipulation
}

//StatusEffectRequest requests a status effect from the server.
type StatusEffectRequest struct {
	Unknown          SVLQ
	StatusEffectName string
	Unknown2         interface{}
	Multiplier       float64
}

//UpdateWorldProperties updates world properties.
type UpdateWorldProperties struct {
	NumPairs      VLQ
	PropertyName  string
	PropertyValue interface{}
}

//Heartbeat  is periodically sent to inform the other party that
// the other end is still connected.
type Heartbeat struct {
	CurrentStep VLQ
}

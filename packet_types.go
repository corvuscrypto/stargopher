package stargopher

//ProtocolVersion is the first packet sent. It contains the server version.
type protocolVersion struct {
	Version uint32
}

//ConnectionResponse tells the client whether their connection
//attempt is successful or if they have been rejected.
//It is the final packet sent in the handshake process.
type connectionResponse struct {
	ID                  PacketType
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
type disconnectResponse struct {
	ID      PacketType
	Unknown uint8
}

//HandshakeChallenge provides a salt and round count for password
//verification. It is followed by Handshake Response.
type handshakeChallenge struct {
	ID           PacketType
	ClaimMessage string
	Salt         string
	HashCount    int32
}

//ChatReceivedPacket is sent to a client with a chat message.
type chatReceived struct {
	ID             PacketType
	Channel        []byte
	UserID         int
	UserNameLength int
	UserName       string
	MessageLength  int
	Message        string
}

//UniverseTimeUpdate is sent from the server to update the current time.
type universeTimeUpdate struct {
	ID   PacketType
	Time SVLQ
}

//CelestialResponse has yet to be fully understood.
type celestialResponse struct {
	ID PacketType
	//Unknown
}

//ClientConnect  is sent in the handshake process immediately after the Protocol
//Version. It contains all relevant data about the connecting player.
type clientConnect struct {
	ID          PacketType
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
type clientDisconnect struct {
	ID      PacketType
	Unknown uint8
}

//HandshakeResponse is the response to the Handshake Challenge.
type handshakeResponse struct {
	ID            PacketType
	ClaimResponse string
	PasswordHash  string
}

//WarpCommand is sent when the player warps/is warped to a planet or ship.
type warpCommand struct {
	ID          PacketType
	WarpType    uint32
	WorldCoords WorldCoord
	PlayerName  string
}

//ChatSent is sent from the client whenever a message is sent in the chat window.
type chatSent struct {
	ID      PacketType
	Message string
	Channel uint8
}

//CelestialRequest has yet to be fully understood.
type celestialRequest struct {
	ID PacketType
	//Unknown
}

//ClientContextUpdate has yet to be fully understood.
type clientContextUpdate struct {
	ID   PacketType
	Data []uint8
}

//WorldStart is sent to the client when a world thread has been started
//on the server.
type worldStart struct {
	ID              PacketType
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
type worldStop struct {
	ID     PacketType
	Status string
}

//TileArrayUpdate is called when an array of tiles has its properties updated.
type tileArrayUpdate struct {
	ID     PacketType
	TileX  SVLQ
	TileY  SVLQ
	Width  VLQ
	Height VLQ
	Tiles  [][]NetTile
}

//TileUpdate is called when a tile is updated.
type tileUpdate struct {
	ID    PacketType
	TileX SVLQ
	TileY SVLQ
	Tile  NetTile
}

//TileLiquidUpdate is sent when the liquid on a tile has changed position.
type tileLiquidUpdate struct {
	ID          PacketType
	TileX       SVLQ
	TileY       SVLQ
	LiquidLevel uint8
	LiquidType  uint8
}

//TileDamageUpdate is sent when a tile is damaged.
type tileDamageUpdate struct {
	ID                     PacketType
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
type tileModificationFailure struct {
	ID PacketType
	//unknown
}

//GiveItem attempts to give an item to a player. If the player's
//inventory is full, it will drop on the ground next to them.
type giveItem struct {
	ID             PacketType
	ItemName       string
	Count          VLQ
	ItemProperties interface{}
}

//ContainerSwapResult is sent whenever two items are
// swapped in an open container.
type containerSwapResult struct {
	ID PacketType
	//unknown
}

//EnvironmentUpdate is sent on an environment update.
type environmentUpdate struct {
	ID PacketType
	//unknown
}

//EntityInteractResult contains the results of an entity interaction.
type entityInteractResult struct {
	ID       PacketType
	ClientID uint32
	EntityID int32
	Results  interface{}
}

//ModifyTileList contains a list of tiles and modifications to them.
type modifyTileList struct {
	ID PacketType
	//Unknown
}

//DamageTile updates a tile's damage
type damageTile struct {
	ID PacketType
	//Unknown
}

//DamageTileGroup updates an entire tile group's damage.
type damageTileGroup struct {
	ID PacketType
	//Unknown
}

//RequestDrop requests an item drop from the ground.
type requestDrop struct {
	ID       PacketType
	EntityID SVLQ
}

//SpawnEntity requests that the server spawn an entity.
type spawnEntity struct {
	ID PacketType
	//Unknown
}

//EntityInteract is sent when a client attempts to interact with an entity.
type entityInteract struct {
	ID PacketType
	//Unknown
}

//ConnectWire connects a wire.
type connectWire struct {
	ID PacketType
	//Unknown
}

//DisconnectAllWires disconnects all wires.
type disconnectAllWires struct {
	ID PacketType
	//Unknown
}

//OpenContainer opens a container.
type openContainer struct {
	ID       PacketType
	EntityID SVLQ
}

//CloseContainer closes a container.
type closeContainer struct {
	ID       PacketType
	EntityID SVLQ
}

//ContainerSwap swaps an item in a container.
type containerSwap struct {
	ID PacketType
	//Unknown
}

//ContainerItemApply applies an item to another item in a container.
type containerItemApply struct {
	ID PacketType
	//Unknown
}

//ContainerStartCrafting initiates crafting on an item in a container
//(Used in pixel compressors and the like?)
type containerStartCrafting struct {
	ID       PacketType
	EntityID SVLQ
}

//ContainerStopCrafting packet stops crafting on an item in a container
type containerStopCrafting struct {
	ID       PacketType
	EntityID SVLQ
}

//BurnContainer burns a container.
type burnContainer struct {
	ID       PacketType
	EntityID SVLQ
}

//ClearContainer clears a container.
type clearContainer struct {
	ID       PacketType
	EntityID SVLQ
}

//WorldClientStateUpdate contains a world client state update
type worldClientStateUpdate struct {
	ID    PacketType
	Delta []uint8
}

//EntityCreate creates an entity.
type entityCreate struct {
	ID         PacketType
	EntityType uint8
	StoreData  []uint8
	EntityID   SVLQ
}

//EntityUpdate updates an entity's properties.
type entityUpdate struct {
	ID       PacketType
	EntityID SVLQ
	Delta    []uint8
}

//EntityDestroy destroys an entity.
type entityDestroy struct {
	ID       PacketType
	EntityID SVLQ
	Death    bool
}

//DamageNotification notifies the receiver of damage received.
type damageNotification struct {
	ID                 PacketType
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
type statusEffectRequest struct {
	ID               PacketType
	Unknown          SVLQ
	StatusEffectName string
	Unknown2         interface{}
	Multiplier       float64
}

//UpdateWorldProperties updates world properties.
type updateWorldProperties struct {
	ID            PacketType
	NumPairs      VLQ
	PropertyName  string
	PropertyValue interface{}
}

//Heartbeat  is periodically sent to inform the other party that
// the other end is still connected.
type heartbeat struct {
	ID          PacketType
	CurrentStep VLQ
}

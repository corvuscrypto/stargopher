package stargopher

type PacketType int

const (
	ProtocolVersion PacketType = iota
	ServerDisconnect
	ConnectSuccess
	ConnectFailure
	HandshakeChallenge
	ChatReceived
	TimeUpdate
	CelestialResponse
	PlayerWarpResult
	ClientConnect
	ClientDisconnectRequest
	HandshakeResponse
	PlayerWarp
	FlyShip
	ChatSent
	CelestialRequest
	ClientContextUpdate
	WorldStart
	WorldStop
	CentralStructureUpdate
	TileArrayUpdate
	TileUpdate
	TileLiquidUpdate
	TileDamageUpdate
	TileModificationFailure
	GiveItem
	ContainerSwapResult
	EnvironmentUpdate
	EntityInteractResult
	UpdateTileProtection
	ModifyTileList
	DamageTileGroup
	CollectLiquid
	RequestDrop
	SpawnEntity
	EntityInteract
	ConnectWire
	DisconnectAllWires
	OpenContainer
	CloseContainer
	ContainerSwap
	ItemApplyInContainer
	ContainerStartCrafting
	ContainerStopCrafting
	BurnContainer
	ClearContainer
	WorldClientStateUpdate
	EntityCreate
	EntityUpdate
	EntityDestroy
	HitRequest
	DamageRequest
	DamageNotification
	EntityMessage
	EntityMessageResponse
	UpdateWorldProperties
	HeartbeatUpdate
)

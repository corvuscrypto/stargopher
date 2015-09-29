package stargopher

type packetType int

const (
	protocolVersion packetType = iota
	serverDisconnect
	connectSuccess
	connectFailure
	handshakeChallenge
	chatReceived
	timeUpdate
	celestialResponse
	playerWarpResult
	clientConnect
	clientDisconnectRequest
	handshakeResponse
	playerWarp
	flyShip
	chatSent
	celestialRequest
	clientContextUpdate
	worldStart
	worldStop
	centralStructureUpdate
	tileArrayUpdate
	tileUpdate
	tileLiquidUpdate
	tileDamageUpdate
	tileModificationFailure
	giveItem
	containerSwapResult
	environmentUpdate
	entityInteractResult
	updateTileProtection
	modifyTileList
	damageTileGroup
	collectLiquid
	requestDrop
	spawnEntity
	entityInteract
	connectWire
	disconnectAllWires
	openContainer
	closeContainer
	containerSwap
	itemApplyInContainer
	containerStartCrafting
	containerStopCrafting
	burnContainer
	clearContainer
	worldClientStateUpdate
	entityCreate
	entityUpdate
	entityDestroy
	hitRequest
	damageRequest
	damageNotification
	entityMessage
	entityMessageResponse
	updateWorldProperties
	heartbeatUpdate
)

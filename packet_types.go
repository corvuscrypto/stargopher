package stargopher

import "reflect"

type basePacket struct {
	ID            PacketType
	PayloadLength int64
}

//ProtocolVersion is the first packet sent. It contains the server version.
type protocolVersion struct {
	basePacket
	Version uint32
}

//DisconnectResponse is used to notify the client of a disconnect.
type serverDisconnect struct {
	basePacket
	Unknown uint8
}

//ConnectionResponse tells the client whether their connection
//attempt is successful or if they have been rejected.
//It is the final packet sent in the handshake process.
type connectSuccess struct {
	basePacket
	Unknown []byte
}

//ConnectionResponse tells the client whether their connection
//attempt is successful or if they have been rejected.
//It is the final packet sent in the handshake process.
type connectFailure struct {
	basePacket
	Unknown []byte
}

//HandshakeChallenge provides a salt and round count for password
//verification. It is followed by Handshake Response.
type handshakeChallenge struct {
	basePacket
	Unknown []byte
}

//ChatReceivedPacket is sent to a client with a chat message.
type chatReceived struct {
	basePacket
	Channel        []byte `length:"5"`
	UserID         uint8
	UserNameLength VLQ
	UserName       string `lengthPrefix:"true"`
	MessageLength  VLQ
	Message        string `lengthPrefix:"true"`
}

//UniverseTimeUpdate is sent from the server to update the current time.
type timeUpdate struct {
	basePacket
	Time int32
}

//CelestialResponse has yet to be fully understood.
type celestialResponse struct {
	basePacket
	Unknown []byte
}

type playerWarpResult struct {
	basePacket
	Unknown []byte
}

//ClientConnect  is sent in the handshake process immediately after the Protocol
//Version. It contains all relevant data about the connecting player.
type clientConnect struct {
	basePacket
	Data []byte
}

//ClientDisconnect is sent when the client disconnects.
type clientDisconnectRequest struct {
	basePacket
	Unknown uint8
}

//HandshakeResponse is the response to the Handshake Challenge.
type handshakeResponse struct {
	basePacket
	ClaimResponse string
	PasswordHash  string
}

//WarpCommand is sent when the player warps/is warped to a planet or ship.
type playerWarp struct {
	basePacket
	WarpType    uint32
	WorldCoords WorldCoord
	PlayerName  string
}

type flyShip struct {
	basePacket
}

//ChatSent is sent from the client whenever a message is sent in the chat window.
type chatSent struct {
	basePacket
	MessageLength uint8
	Message       string `lengthPrefix:"True"`
	Channel       uint8
}

//CelestialRequest has yet to be fully understood.
type celestialRequest struct {
	basePacket
	Unknown []byte
}

//ClientContextUpdate has yet to be fully understood.
type clientContextUpdate struct {
	basePacket
	Data []byte
}

//WorldStart is sent to the client when a world thread has been started
//on the server.
type worldStart struct {
	basePacket
	Data []byte
}

//WorldStop is called when a world thread is stopped.
type worldStop struct {
	basePacket
	Status string
}

//WorldStop is called when a world thread is stopped.
type centralStructureUpdate struct {
	basePacket
	Unknown []byte
}

//TileArrayUpdate is called when an array of tiles has its properties updated.
type tileArrayUpdate struct {
	basePacket
	TileX  int64
	TileY  int64
	Width  int64
	Height int64
	Tiles  []byte
}

//TileUpdate is called when a tile is updated.
type tileUpdate struct {
	basePacket
	TileX int64
	TileY int64
	Tile  []byte
}

//TileLiquidUpdate is sent when the liquid on a tile has changed position.
type tileLiquidUpdate struct {
	basePacket
	TileX       int64
	TileY       int64
	LiquidLevel uint8
	LiquidType  uint8
}

//TileDamageUpdate is sent when a tile is damaged.
type tileDamageUpdate struct {
	basePacket
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
	basePacket
	Unknown []byte
}

//GiveItem attempts to give an item to a player. If the player's
//inventory is full, it will drop on the ground next to them.
type giveItem struct {
	basePacket
	ItemNameLength VLQ
	ItemName       string
	Count          VLQ
	ItemProperties []byte
}

//ContainerSwapResult is sent whenever two items are
// swapped in an open container.
type containerSwapResult struct {
	basePacket
	Unknown []byte
}

//EnvironmentUpdate is sent on an environment update.
type environmentUpdate struct {
	basePacket
	Data []byte
}

//EntityInteractResult contains the results of an entity interaction.
type entityInteractResult struct {
	basePacket
	ClientID uint32
	EntityID int32
	Results  []byte
}

type updateTileProtection struct {
	basePacket
	Data []byte
}

//ModifyTileList contains a list of tiles and modifications to them.
type modifyTileList struct {
	basePacket
	Data []byte
}

//DamageTileGroup updates an entire tile group's damage.
type damageTileGroup struct {
	basePacket
	Data []byte
}

type collectLiquid struct {
	basePacket
	Data []byte
}

//RequestDrop requests an item drop from the ground.
type requestDrop struct {
	basePacket
	Data []byte
}

//SpawnEntity requests that the server spawn an entity.
type spawnEntity struct {
	basePacket
	Data []byte
}

//EntityInteract is sent when a client attempts to interact with an entity.
type entityInteract struct {
	basePacket
	Data []byte
}

//ConnectWire connects a wire.
type connectWire struct {
	basePacket
	Data []byte
}

//DisconnectAllWires disconnects all wires.
type disconnectAllWires struct {
	basePacket
	Data []byte
}

//OpenContainer opens a container.
type openContainer struct {
	basePacket
	EntityID VLQ
}

//CloseContainer closes a container.
type closeContainer struct {
	basePacket
	EntityID VLQ
}

//ContainerSwap swaps an item in a container.
type containerSwap struct {
	basePacket
	Data []byte
}

//ContainerItemApply applies an item to another item in a container.
type itemApplyInContainer struct {
	basePacket
	Data []byte
}

//ContainerStartCrafting initiates crafting on an item in a container
//(Used in pixel compressors and the like?)
type containerStartCrafting struct {
	basePacket
	EntityID int64
}

//ContainerStopCrafting packet stops crafting on an item in a container
type containerStopCrafting struct {
	basePacket
	EntityID int64
}

//BurnContainer burns a container.
type burnContainer struct {
	basePacket
	EntityID int64
}

//ClearContainer clears a container.
type clearContainer struct {
	basePacket
	EntityID int64
}

//WorldClientStateUpdate contains a world client state update
type worldClientStateUpdate struct {
	basePacket
	Delta []uint8
}

//EntityCreate creates an entity.
type entityCreate struct {
	basePacket
	Data []byte
}

//EntityUpdate updates an entity's properties.
type entityUpdate struct {
	basePacket
	EntityID int64
	Delta    []uint8
}

//EntityDestroy destroys an entity.
type entityDestroy struct {
	basePacket
	Data []byte
}

type hitRequest struct {
	basePacket
	Data []byte
}

type damageRequest struct {
	basePacket
	Data []byte
}

//DamageNotification notifies the receiver of damage received.
type damageNotification struct {
	basePacket
	CausingEntityID    int64
	TargetEntityID     int64
	PosX               int64 //*
	PosY               int64 //*
	Damage             int64 //*
	DamageType         uint8
	DamageSourceType   string
	TargetMaterialType string
	HitResultType      uint8

	//* denotes that values must be divided by 100 before manipulation
}

type entityMessage struct {
	basePacket
	Data []byte
}

type entityMessageResponse struct {
	basePacket
	Data []byte
}

//UpdateWorldProperties updates world properties.
type updateWorldProperties struct {
	basePacket
	NumPairs      int64
	PropertyName  string
	PropertyValue []byte
}

//Heartbeat  is periodically sent to inform the other party that
// the other end is still connected.
type heartbeatUpdate struct {
	basePacket
	CurrentStep VLQ
}

var packetRegistry = []reflect.Type{
	reflect.TypeOf(protocolVersion{}),
	reflect.TypeOf(serverDisconnect{}),
	reflect.TypeOf(connectSuccess{}),
	reflect.TypeOf(connectFailure{}),
	reflect.TypeOf(handshakeChallenge{}),
	reflect.TypeOf(chatReceived{}),
	reflect.TypeOf(timeUpdate{}),
	reflect.TypeOf(celestialResponse{}),
	reflect.TypeOf(playerWarpResult{}),
	reflect.TypeOf(clientConnect{}),
	reflect.TypeOf(clientDisconnectRequest{}),
	reflect.TypeOf(handshakeResponse{}),
	reflect.TypeOf(playerWarp{}),
	reflect.TypeOf(flyShip{}),
	reflect.TypeOf(chatSent{}),
	reflect.TypeOf(celestialRequest{}),
	reflect.TypeOf(clientContextUpdate{}),
	reflect.TypeOf(worldStart{}),
	reflect.TypeOf(worldStop{}),
	reflect.TypeOf(centralStructureUpdate{}),
	reflect.TypeOf(tileArrayUpdate{}),
	reflect.TypeOf(tileUpdate{}),
	reflect.TypeOf(tileLiquidUpdate{}),
	reflect.TypeOf(tileDamageUpdate{}),
	reflect.TypeOf(tileModificationFailure{}),
	reflect.TypeOf(giveItem{}),
	reflect.TypeOf(containerSwapResult{}),
	reflect.TypeOf(environmentUpdate{}),
	reflect.TypeOf(entityInteractResult{}),
	reflect.TypeOf(updateTileProtection{}),
	reflect.TypeOf(modifyTileList{}),
	reflect.TypeOf(damageTileGroup{}),
	reflect.TypeOf(collectLiquid{}),
	reflect.TypeOf(requestDrop{}),
	reflect.TypeOf(spawnEntity{}),
	reflect.TypeOf(entityInteract{}),
	reflect.TypeOf(connectWire{}),
	reflect.TypeOf(disconnectAllWires{}),
	reflect.TypeOf(openContainer{}),
	reflect.TypeOf(closeContainer{}),
	reflect.TypeOf(containerSwap{}),
	reflect.TypeOf(itemApplyInContainer{}),
	reflect.TypeOf(containerStartCrafting{}),
	reflect.TypeOf(containerStopCrafting{}),
	reflect.TypeOf(burnContainer{}),
	reflect.TypeOf(clearContainer{}),
	reflect.TypeOf(worldClientStateUpdate{}),
	reflect.TypeOf(entityCreate{}),
	reflect.TypeOf(entityUpdate{}),
	reflect.TypeOf(entityDestroy{}),
	reflect.TypeOf(hitRequest{}),
	reflect.TypeOf(damageRequest{}),
	reflect.TypeOf(damageNotification{}),
	reflect.TypeOf(entityMessage{}),
	reflect.TypeOf(entityMessageResponse{}),
	reflect.TypeOf(updateWorldProperties{}),
	reflect.TypeOf(heartbeatUpdate{}),
}

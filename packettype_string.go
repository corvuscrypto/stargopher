// generated by stringer -type=PacketType; DO NOT EDIT

package stargopher

import "fmt"

const _PacketType_name = "protocolVersionserverDisconnectconnectSuccessconnectFailurehandshakeChallengechatReceivedtimeUpdatecelestialResponseplayerWarpResultclientConnectclientDisconnectRequesthandshakeResponseplayerWarpflyShipchatSentcelestialRequestclientContextUpdateworldStartworldStopcentralStructureUpdatetileArrayUpdatetileUpdatetileLiquidUpdatetileDamageUpdatetileModificationFailuregiveItemcontainerSwapResultenvironmentUpdateentityInteractResultupdateTileProtectionmodifyTileListdamageTileGroupcollectLiquidrequestDropspawnEntityentityInteractconnectWiredisconnectAllWiresopenContainercloseContainercontainerSwapitemApplyInContainercontainerStartCraftingcontainerStopCraftingburnContainerclearContainerworldClientStateUpdateentityCreateentityUpdateentityDestroyhitRequestdamageRequestdamageNotificationentityMessageentityMessageResponseupdateWorldPropertiesheartbeatUpdate"

var _PacketType_index = [...]uint16{0, 15, 31, 45, 59, 77, 89, 99, 116, 132, 145, 168, 185, 195, 202, 210, 226, 245, 255, 264, 286, 301, 311, 327, 343, 366, 374, 393, 410, 430, 450, 464, 479, 492, 503, 514, 528, 539, 557, 570, 584, 597, 617, 639, 660, 673, 687, 709, 721, 733, 746, 756, 769, 787, 800, 821, 842, 857}

func (i PacketType) String() string {
	if i < 0 || i >= PacketType(len(_PacketType_index)-1) {
		return fmt.Sprintf("PacketType(%d)", i)
	}
	return _PacketType_name[_PacketType_index[i]:_PacketType_index[i+1]]
}

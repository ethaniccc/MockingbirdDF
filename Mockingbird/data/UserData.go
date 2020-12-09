package data

import (
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/go-gl/mathgl/mgl64"
	"sync"
)

type UserData struct {
	MoveData MovementData
	BasicData BasicData
}

type MovementData struct {
	Location, LastLocation                                                             mgl64.Vec3
	MoveDelta, LastMoveDelta                                                           mgl64.Vec3
	OnGround                                                                           bool
	OnGroundTicks, OffGroundTicks                                                      uint64
	BlockAbove, BlockBelow                                                             world.Block
	LastMotion                                                                         mgl64.Vec3
	Yaw, Pitch, LastYaw, LastPitch, YawDelta, PitchDelta, LastYawDelta, LastPitchDelta float64
	AppendingTeleport                                                                  bool
	TeleportPos                                                                        mgl64.Vec3
	// TODO: Add more data
}

type BasicData struct {
	loggedIn bool
}

var dataList = sync.Map{}

func createData(player *player.Player){
	dataList.Store(player, UserData{})
}

func GetData(player *player.Player) (*UserData, bool){
	v, enabled := dataList.Load(player)
	if !enabled{
		return nil, false
	}
	return v.(*UserData), true
}
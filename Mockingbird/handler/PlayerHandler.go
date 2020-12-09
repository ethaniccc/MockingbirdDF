package handler

import (
	"github.com/df-mc/dragonfly/dragonfly/event"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/go-gl/mathgl/mgl64"
	"math"
	"server/Mockingbird/data"
	"server/Mockingbird/detections"
	"sync"
)

type PlayerHandler struct {
	player.NopHandler
	detections detections.DetectionList
	p *player.Player
}

var handlers = sync.Map{}

func NewPlayerHandler(player *player.Player) *PlayerHandler{
	handler := &PlayerHandler{p: player, detections: detections.DetectionList{}}
	handlers.Store(player, handler)
	return handler
}

func getHandler(player *player.Player) (*PlayerHandler, bool){
	v, enabled := handlers.Load(player)
	if !enabled{
		return nil, false
	}
	return v.(*PlayerHandler), true
}

func (handler PlayerHandler) HandleMove(_ *event.Context, pos mgl64.Vec3, _, _ float64){
	var p = handler.p
	var userData, hasData = data.GetData(p)
	if !hasData{
		return
	} else {
		userData.MoveData.LastLocation = userData.MoveData.Location
		userData.MoveData.Location = p.Position()
		userData.MoveData.LastMoveDelta = userData.MoveData.MoveDelta
		userData.MoveData.MoveDelta = userData.MoveData.Location.Sub(userData.MoveData.LastLocation)
		userData.MoveData.OnGround = math.Mod(mgl64.Round(userData.MoveData.Location.Y(), 4), 0.015625) == 0
		if userData.MoveData.OnGround {
			userData.MoveData.OnGroundTicks++
			userData.MoveData.OffGroundTicks = 0
		} else {
			userData.MoveData.OffGroundTicks++
			userData.MoveData.OnGroundTicks = 0
		}
		userData.MoveData.LastYaw = userData.MoveData.Yaw
		userData.MoveData.LastPitch = userData.MoveData.Pitch
		userData.MoveData.Yaw = p.Yaw()
		userData.MoveData.Pitch = p.Pitch()
		userData.MoveData.LastYawDelta = userData.MoveData.YawDelta
		userData.MoveData.LastPitchDelta = userData.MoveData.PitchDelta
		userData.MoveData.YawDelta = math.Abs(userData.MoveData.Yaw - userData.MoveData.LastYaw)
		userData.MoveData.PitchDelta = math.Abs(userData.MoveData.Pitch - userData.MoveData.LastPitch)
	}
}
package detections

import (
	"github.com/df-mc/dragonfly/dragonfly/player"
	"math"
	"server/Mockingbird/data"
)

type FlyA struct {
	Detection
	DetectionBase
}

func (d FlyA) check(player *player.Player){
	var userData, hasData = data.GetData(player)
	if !hasData{
		return
	} else {
		if userData.MoveData.OffGroundTicks >= 10 {
			var currentYDelta = userData.MoveData.MoveDelta.Y()
			var lastYDelta = userData.MoveData.LastMoveDelta.Y()
			var prediction = (lastYDelta - 0.08) * 0.980000019073486
			var equalness = math.Abs(currentYDelta - prediction)
			if equalness > 0.0015 {
				d.preVL++
				if d.preVL >= 3 {
					d.fail(player)
				}
			} else {
				d.preVL *= 0.75
			}
		}
	}
}
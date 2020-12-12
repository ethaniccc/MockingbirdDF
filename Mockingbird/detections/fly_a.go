package detections

import (
	"fmt"
	"math"
	"server/Mockingbird/data"
)

type FlyA struct {
	DetectionBase
}

func (d *FlyA) Check(data *data.UserData){
	if data.MoveData.OffGroundTicks >= 10 {
		currentYDelta, lastYDelta := data.MoveData.MoveDelta.Y(), data.MoveData.LastMoveDelta.Y()
		prediction := (lastYDelta - 0.08) * 0.980000019073486
		equalness := math.Abs(currentYDelta - prediction)
		fmt.Printf("current=" + fmt.Sprintf("%f", currentYDelta) + " last=" + fmt.Sprintf("%f", lastYDelta) + " equalness=" + fmt.Sprintf("%f", equalness) + "\n")
		if equalness > 0.0015 {
			d.preVL = d.preVL + 1
			fmt.Printf("preVL=" + fmt.Sprintf("%f", d.preVL) + "\n")
			if d.preVL >= 3 {
				d.fail(data.Player)
			}
		} else {
			d.preVL = d.preVL * 0.75
		}
	}
}
package detections

import (
	"fmt"
	"github.com/df-mc/dragonfly/dragonfly/player"
)

type DetectionBase struct {

	name string
	isExperimental bool
	description string
	maxVL uint64
	currentVL uint64
	preVL float64

}

type Detection interface {
	check(player player.Player)
}

func (d DetectionBase) fail(player *player.Player){
	d.currentVL++
	if d.currentVL >= d.maxVL{
		fmt.Printf("punished player " + player.Name() + " for " + d.name)
		player.Disconnect("Unfair advantage detected (" + d.name + ")")
	}
	player.Message("You were detected for " + d.name + " @ " + string(d.currentVL))
}
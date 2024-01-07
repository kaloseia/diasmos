package plugin

import (
	ti "time"

	"github.com/kaloseia/diasmos/domain/time"
	"github.com/kaloseia/diasmos/domain/uuid"
)

func NewInstallation(plugin Plugin) Installation {
	return Installation{
		UUID:      uuid.New(),
		Plugin:    plugin,
		StartedAt: time.Now(),
	}
}

type Installation struct {
	UUID      string
	Plugin    Plugin
	StartedAt ti.Time

	//context       context.Context
	//contextCancel context.CancelFunc
}

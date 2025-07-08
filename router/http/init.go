package http

import (
	"fmt"
	"github.com/emenesism/Decentralized-voting-backend/config"
)

func Init() {
	r := NewRouter()
	r.Run(fmt.Sprintf("%s:%d", config.AppConfig.Host, config.AppConfig.Port))
}

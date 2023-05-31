package global

import (
	"ums_backend/config"

	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
)

var (
	GVA_CONFIG              config.Server
	GVA_Concurrency_Control = &singleflight.Group{}
	GVA_LOG                 *zap.Logger
)

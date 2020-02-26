package providersvc

import (
	"github.com/gin-gonic/gin"
	"github.com/mateuszmierzwinski/libmicro/config"
)

// APIProvider interface describes provider structure and methods
type APIProvider interface {

	// GetInfo returns information about existing Provider implementation
	GetInfo() *APIInfo

	// Register registers provider in provided router engine with config provided
	Register(router *gin.Engine, configProvider config.Provider)
}

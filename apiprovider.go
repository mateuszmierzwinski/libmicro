package libmicro

import "github.com/gin-gonic/gin"

type APIInfo struct {
	ProviderName string
}

type ApiProvider interface {
	GetInfo() *APIInfo
	Register(router gin.Router)
}

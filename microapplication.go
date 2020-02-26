package libmicro

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mateuszmierzwinski/libmicro/config"
	"github.com/mateuszmierzwinski/libmicro/providersvc"
	"log"
	"net/http"
)

// MicroApplication is application context that stores all providers and middleware
type MicroApplication struct {
	providers  []providersvc.APIProvider
	middleware []gin.HandlerFunc
}

// RegisterMiddleware adds middlewares to list for further load on application execution
func (microApp *MicroApplication) RegisterMiddleware(middleware ...gin.HandlerFunc) *MicroApplication {
	microApp.middleware = append(microApp.middleware, middleware...)
	return microApp
}

// RegisterProviders registers API provider
func (microApp *MicroApplication) RegisterProviders(providers ...providersvc.APIProvider) *MicroApplication {
	microApp.providers = append(microApp.providers, providers...)
	return microApp
}

// Execute function starts application with provided configuration
func (microApp *MicroApplication) Execute(configProvider config.Provider) error {
	if configProvider == nil {
		return fmt.Errorf("configProvider is nil")
	}

	router := gin.New()
	router.Use(gin.Recovery())

	if gin.Mode() != gin.ReleaseMode {
		router.Use(gin.Logger())
	}

	if q := len(microApp.middleware); q > 0 {
		router.Use(microApp.middleware...)
	}

	var providersActive []string

	for _, provider := range microApp.providers {
		provider.Register(router, configProvider)
		providersActive = append(providersActive, provider.GetInfo().ProviderName)
		log.Printf("Registered provider: %s", provider.GetInfo().ProviderName)
	}

	log.Fatal(
		http.ListenAndServe(
			configProvider.GetConfigWithDefaultValue("MICROCFG_BIND_ADDR", ":8080"),
			router,
		))

	return nil
}

// NewMicroApplication simplifies building of framework application context
func NewMicroApplication() *MicroApplication {
	return &MicroApplication{}
}

# LibMicro framework

LibMicro is a framework that is based on Gin (https://github.com/gin-gonic/).

## How to use?

To use LibMicro you need to setup some configuration first.

### 1. Add new module to go.mod

```
module libmicroAwesomeProject

go 1.13

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/mateuszmierzwinski/libmicro v0.0.1
)
```

Now you can pull dependencies with use of:

```go
go mod tidy
```

### 2. Include new imports in your main.go file

```go
import (
	"github.com/mateuszmierzwinski/libmicro"
	"github.com/mateuszmierzwinski/libmicro/config/cmdconfigprovider"
)
```

### 3. Define new provider in your project as separated package

```go
package serviceprovider

import (
	"github.com/gin-gonic/gin"
	"github.com/mateuszmierzwinski/libmicro/config"
	"github.com/mateuszmierzwinski/libmicro/providersvc"
	"log"
)

type serviceProvider struct {

}

func (s serviceProvider) GetInfo() *providersvc.APIInfo {
	return &providersvc.APIInfo{ProviderName: "Simple Provider"}
}

func (s serviceProvider) Register(router *gin.Engine, configProvider config.ConfigProvider) {
  // here register your path binding to a method
}

func New() providersvc.APIProvider {
	return &serviceProvider{}
}
```

### 4. Initialize application context and inject service providers

```go
import (
  ...
	"libmicroAwesomeProject/serviceprovider"
)

func main() {
	(&libmicro.MicroApplication{}).
		RegisterProviders(
			serviceprovider.New(),
		).
		Execute(cmdconfigprovider.New())
}
```

## How to read configuration?

Default implementation of LibMicro has embedded set of configuration providers. You can choose from:

- cmdconfigprovider - CMD Configuration Provider package handling parameters provided by command line
- envconfigprovider - Environment variables based Configuration Provider handling parameters provided by ENV variables of operating system or container
- yamlconfigprovider - Configuration provider handling parameters given by yaml config file

Each of config providers works differently but from application level it gives same interface:

```go
// GetConfigByName returns configuration value by provided key or empty string if does not exist
	GetConfigByName(string) string

	// GetConfigWithDefaultValue returns configuration value by provided key or default value if does not exist
	GetConfigWithDefaultValue(string, string) string

	// OverrideWithValue allows to change programmatically configuration by key
	OverrideWithValue(string, string)
```

## QA section

### Q: Why another framework/library?

A: We would say - for fun - but it's not only that. We made some set of micro services in our past and pattern was mostly the same. We were thinking about extracting this repetitive actions into separated library so we would speed things up a lot by developing only real functionality. Each time we should not care about HTTP server, authorization. We would like to skip directly to write business-case code. This library is somehow allowing us to do so.

### Q: Why this project is based on Gin?

A: Arbitrarily we've tested multiple ways to handle REST micro services. We've been using Gorilla MUX in combination with HTTP Server provided by Golang. We've also been using gRPC. We've selected Gin because from one perspective REST is still most common in modern webservices development. gRPC requires contract and special handling. MUX package by Gorilla had been slower than Gin. For us those were main criteria to select Gin as our HTTP provider.

### Q: Is this library stable?

A: You can consider it as a stable when you use released version. Each released version had been tested extensively against set of applications to prevent backward compatibility. Unreleased versions (eg. master branch) are not stable and you should use always released version.

### Q: Why it's MIT licensed?

A: Because this software is for good of all. If you like to use it please do. There is no hidden restrictions over this code (check MIT license). We're developing it for ourselves but if you would like to use it - please do. If you would like to contribute - you're welcome to join our project.

### Q: Do you consider further extensions of this library?

A: Yes, we have some ideas how we can extend this library. We would like to add more stuff working out of the box, like for example JWT/JWK security, pre-made providers, CLI to generate some stuff for us (and you). Our list of features is great and we will release it when we will have more time to focus on them but it will be done.

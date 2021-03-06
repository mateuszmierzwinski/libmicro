package envconfigprovider

import (
	"github.com/mateuszmierzwinski/libmicro/config"
	"os"
	"strings"
)

// EnvConfigProvider is default config provider that is able to read configuration from environment
type EnvConfigProvider struct {
	config map[string]string
}

// GetConfigByName returns configuration value by provided key or empty string if does not exist
func (c *EnvConfigProvider) GetConfigByName(configName string) string {
	if v, ok := c.config[configName]; ok {
		return v
	}

	return ""
}

// GetConfigWithDefaultValue returns configuration value by provided key or default value if does not exist
func (c *EnvConfigProvider) GetConfigWithDefaultValue(configName string, defaultValue string) string {
	if v, ok := c.config[configName]; ok {
		if len(v) == 0 {
			return defaultValue
		}

		return v
	}

	return defaultValue
}

// OverrideWithValue allows to change programatically configuration by key
func (c *EnvConfigProvider) OverrideWithValue(configName string, valueToSet string) {
	c.config[configName] = valueToSet
}

func (c *EnvConfigProvider) initModule(envVariables []string) {
	c.config = make(map[string]string)
	for _, env := range envVariables {
		if len(env) > 0 {
			vars := strings.Split(env, "=")

			switch len(vars) {
			case 2:
				c.config[vars[0]] = vars[1]
				break
			case 1:
				c.config[vars[0]] = ""
				break
			default:
				break
			}
		}
	}
}

// New constructs new provider and initializes this module on the fly
func New() config.Provider {
	c := &EnvConfigProvider{}
	c.initModule(os.Environ())
	return c
}

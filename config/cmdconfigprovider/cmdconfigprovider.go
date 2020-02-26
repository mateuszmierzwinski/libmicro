package cmdconfigprovider

import (
	"github.com/mateuszmierzwinski/libmicro/config"
	"os"
	"strings"
)

// CmdConfigProvider is default config provider that is able to read configuration from command line
type CmdConfigProvider struct {
	config map[string]string
}

// GetConfigByName returns configuration value by provided key or empty string if does not exist
func (c *CmdConfigProvider) GetConfigByName(configName string) string {
	if v, ok := c.config[configName]; ok {
		return v
	}

	return ""
}

// GetConfigWithDefaultValue returns configuration value by provided key or default value if does not exist
func (c *CmdConfigProvider) GetConfigWithDefaultValue(configName string, defaultValue string) string {
	if v, ok := c.config[configName]; ok {
		if len(v) == 0 {
			return defaultValue
		}

		return v
	}

	return defaultValue
}

// OverrideWithValue allows to change programatically configuration by key
func (c *CmdConfigProvider) OverrideWithValue(configName string, valueToSet string) {
	c.config[configName] = valueToSet
}

func (c *CmdConfigProvider) initModule(cmd []string) {
	cmdSz := len(cmd)
	c.config = make(map[string]string)
	for i := 1; i < cmdSz; i++ {
		if strings.HasPrefix(cmd[i], "-") {
			key := cmd[i]
			// remove trailing negatives
			for strings.HasPrefix(key, "-") {
				key = key[1:]
			}

			// only negatives case - drop them
			if len(key) == 0 {
				continue
			}

			// try get value
			if cmdSz > i+1 {
				// analyze next entry in cmd
				potentialVal := cmd[i+1]

				// next parameter?
				if strings.HasPrefix(potentialVal, "-") {
					c.config[key] = ""
				} else {
					c.config[key] = potentialVal
				}
			} else {
				// support of empty key at the end
				c.config[key] = ""
			}
		}
	}
}

// New constructs new provider and initializes this module on the fly
func New() config.Provider {
	c := &CmdConfigProvider{}
	c.initModule(os.Args)
	return c
}

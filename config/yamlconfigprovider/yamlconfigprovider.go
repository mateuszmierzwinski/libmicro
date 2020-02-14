package yamlconfigprovider

import (
	"errors"
	"github.com/mateuszmierzwinski/libmicro/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// YamlConfigProvider is default config provider that is able to read configuration from yaml file
type YamlConfigProvider struct {
	config map[string]string
}

// GetConfigByName returns configuration value by provided key or empty string if does not exist
func (c *YamlConfigProvider) GetConfigByName(configName string) string {
	if v,ok := c.config[configName]; !ok {
		return ""
	} else {
		return v
	}
}

// GetConfigWithDefaultValue returns configuration value by provided key or default value if does not exist
func (c *YamlConfigProvider) GetConfigWithDefaultValue(configName string, defaultValue string) string {
	if v,ok := c.config[configName]; !ok {
		return defaultValue
	} else {
		if len(v) == 0 {
			return defaultValue
		} else {
			return v
		}
	}
}

// OverrideWithValue allows to change programatically configuration by key
func (c *YamlConfigProvider) OverrideWithValue(configName string, valueToSet string) {
	c.config[configName] = valueToSet
}

func (c *YamlConfigProvider) loadYamlFromPath() error {
	appDir,_ := filepath.Abs(os.Args[0])
	appName := strings.ToLower(filepath.Base(os.Args[0]))
	wd,_ := os.Getwd()
	ud,_ := os.UserConfigDir()
	pd,_ := os.UserHomeDir()

	paths := []string {
		filepath.Join(wd, appName + ".yaml"),
		filepath.Join(filepath.Dir(appDir), appName + ".yaml"),
		filepath.Join(ud, appName + ".yaml"),
		filepath.Join(pd, ".config", appName, appName + ".yaml"),
	}

	for _,path := range paths {
		if _,err := os.Lstat(path); err == nil {
			log.Println("-- Found configuration to load in:", path)
			return c.yamlLoad(path)
		} else {
			log.Println("-- Configuration not found in:", path)
		}
	}

	return errors.New("cannot find valid configuration in any of expected and listed above files")
}

func (c *YamlConfigProvider) yamlLoad(yamlPath string) error {
	pulledData,err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(pulledData, c.config)
	if err != nil {
		return err
	}

	return nil
}

func (c *YamlConfigProvider) initFromFile() error {
	c.config = make(map[string]string)

	err := c.loadYamlFromPath()
	if err != nil {
		return err
	}

	return nil
}

func New() config.ConfigProvider {
	c := &YamlConfigProvider{}

	err := c.initFromFile()
	if err != nil {
		log.Fatal(err)
	}

	return c
}
package yamlconfigprovider

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/mateuszmierzwinski/libmicro/config"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testCases = []struct {
	f        func(c *YamlConfigProvider) error
	expected error
}{
	{
		f: func(c *YamlConfigProvider) error {
			return c.initFromFile()
		},
		expected: errors.New("cannot find valid configuration in any of expected and listed above files"),
	}, {
		f: func(c *YamlConfigProvider) error {
			appDir, _ := filepath.Abs(os.Args[0])
			appName := strings.ToLower(filepath.Base(os.Args[0]))
			path := filepath.Join(filepath.Dir(appDir), appName+".yaml")

			bb := bytes.NewBuffer([]byte{})
			bb.WriteString("config_a: value\n")
			bb.WriteString("config_b: \"value\"\n")
			bb.WriteString("config_c: \"10\"\n")
			bb.WriteString("config_d: 10\n")

			fmt.Println(path)

			err := ioutil.WriteFile(path, bb.Bytes(), os.ModePerm)
			if err != nil {
				return err
			}
			defer os.Remove(path)

			return c.initFromFile()
		},
		expected: nil,
	}, {
		f: func(c *YamlConfigProvider) error {
			if c.GetConfigByName("config_a") != "value" {
				return errors.New("config_a != value")
			}

			return nil
		},
		expected: nil,
	}, {
		f: func(c *YamlConfigProvider) error {
			if c.GetConfigByName("config_b") != "value" {
				return errors.New("config_b != value")
			}

			return nil
		},
		expected: nil,
	}, {
		f: func(c *YamlConfigProvider) error {
			if c.GetConfigByName("config_c") != "10" {
				return errors.New("config_c != 10")
			}

			return nil
		},
		expected: nil,
	}, {
		f: func(c *YamlConfigProvider) error {
			if c.GetConfigByName("config_d") != "10" {
				return errors.New("config_d != 10")
			}
			return nil
		},
		expected: nil,
	}, {
		f: func(c *YamlConfigProvider) error {
			if c.GetConfigByName("config_e") != "" {
				return errors.New("config_d is not empty")
			}
			return nil
		},
		expected: nil,
	}, {
		f: func(c *YamlConfigProvider) error {
			c.OverrideWithValue("config_e", "alpha")
			if c.GetConfigByName("config_e") != "alpha" {
				return errors.New("config_d is not equal alpha")
			}
			return nil
		},
		expected: nil,
	}, {
		f: func(c *YamlConfigProvider) error {
			c.OverrideWithValue("config_b", "beta")
			if c.GetConfigByName("config_b") != "beta" {
				return errors.New("config_b is not equal beta")
			}
			return nil
		},
		expected: nil,
	},
}

func TestYamlConfigProviderInitializer(t *testing.T) {
	c := &YamlConfigProvider{}
	assert.Implements(t, (*config.Provider)(nil), c)
	assert.NotNil(t, c)

	for _, tc := range testCases {
		exp := tc.f(c)
		assert.Equal(t, tc.expected, exp)
	}
}

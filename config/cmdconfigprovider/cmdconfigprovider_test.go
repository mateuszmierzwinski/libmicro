package cmdconfigprovider

import (
	"github.com/mateuszmierzwinski/libmicro/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCmdConfigProviderInitializer(t *testing.T) {
	c := &CmdConfigProvider{}
	c.initModule([]string{"testRunnerApp", "-p", "test0", "--test", "test1", "-t", "-T", "test3", "-", "--"})

	assert.Equal(t, "test0", c.GetConfigByName("p"))
	assert.Equal(t, "test1", c.GetConfigByName("test"))
	assert.Equal(t, "", c.GetConfigByName("t"))
	assert.Equal(t, "test3", c.GetConfigByName("T"))
	assert.Equal(t, len(c.config), 4)
}

func TestCmdConfigProviderGetConfigWithDefaultValue(t *testing.T) {
	c := &CmdConfigProvider{}
	c.initModule([]string{"testRunnerApp", "-p", "test0", "--test", "test1", "-t", "-T", "test3", "-", "--", "-s"})

	assert.Equal(t, "", c.GetConfigByName("deep"))
	assert.Equal(t, "", c.GetConfigByName("s"))
	assert.Equal(t, "alpha", c.GetConfigWithDefaultValue("deep", "alpha"))
	assert.Equal(t, "test1", c.GetConfigWithDefaultValue("test", "beta"))
	assert.Equal(t, len(c.config), 5)

	c.OverrideWithValue("test", "beta")
	assert.Equal(t, "beta", c.GetConfigByName("test"))

	assert.Equal(t, "gamma", c.GetConfigWithDefaultValue("s", "gamma"))
}

func TestNew(t *testing.T) {
	c := New()
	assert.Implements(t, (*config.Provider)(nil), c)
	assert.NotNil(t, c)
}

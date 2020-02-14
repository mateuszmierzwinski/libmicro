package envconfigprovider

import (
	"github.com/mateuszmierzwinski/libmicro/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnvConfigProviderInitializer(t *testing.T) {
	c := &EnvConfigProvider{}
	c.initModule([]string{ "p=test0", "test=test1", "t=", "T=test3", "somestring", ""})

	assert.Equal(t, "test0", c.GetConfigByName("p"))
	assert.Equal(t, "test1", c.GetConfigByName("test"))
	assert.Equal(t, "", c.GetConfigByName("t"))
	assert.Equal(t, "test3", c.GetConfigByName("T"))
	assert.Equal(t, "", c.GetConfigByName("somestring"))
	assert.Equal(t, 5, len(c.config))
}

func TestEnvConfigProviderGetConfigWithDefaultValue(t *testing.T) {
	c := &EnvConfigProvider{}
	c.initModule([]string{ "p=test0", "test=test1", "t=", "T=test3", "somestring", ""})

	assert.Equal(t, "", c.GetConfigByName("deep"))
	assert.Equal(t, "", c.GetConfigByName("somestring"))
	assert.Equal(t, "alpha", c.GetConfigWithDefaultValue("deep", "alpha"))
	assert.Equal(t, "test1", c.GetConfigWithDefaultValue("test", "beta"))
	assert.Equal(t, 5, len(c.config))

	c.OverrideWithValue("test", "beta")
	assert.Equal(t, "beta", c.GetConfigByName("test"))

	assert.Equal(t, "gamma", c.GetConfigWithDefaultValue("somestring", "gamma"))
}

func TestNew(t *testing.T) {
	c := New()
	assert.Implements(t, (*config.ConfigProvider)(nil), c)
	assert.NotNil(t, c)
}
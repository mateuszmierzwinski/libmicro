package config

// ConfigProvider is interface for configuration loaders that can read config from different sources
type ConfigProvider interface {

	// GetConfigByName returns configuration value by provided key or empty string if does not exist
	GetConfigByName(string) string

	// GetConfigWithDefaultValue returns configuration value by provided key or default value if does not exist
	GetConfigWithDefaultValue(string, string) string

	// OverrideWithValue allows to change programatically configuration by key
	OverrideWithValue(string, string)
}

package cfg

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"sync"
)

var once sync.Once

type ConfigReader interface {
	GetStringWithDefault(key string, defaultValue string) string
	GetBoolWithDefault(key string, defaultValue bool) bool
	GetIntWithDefault(key string, defaultValue int) int
	GetInt64WithDefault(key string, defaultValue int64) int64
	MustGetString(key string) string
	MustGetBool(key string) bool
	MustGetInt(key string) int
	MustGetInt64(key string) int64
	MustGetStringSlice(key string) []string
}

type configReader struct {
	v *viper.Viper
}

var cfgReader *configReader

func Reader() ConfigReader {
	once.Do(func() {
		v := viper.New()
		v.AutomaticEnv()
		cfgReader = &configReader{
			v: v,
		}
	})

	return cfgReader
}
func (c *configReader) MustGetString(key string) string {
	if !c.v.IsSet(key) {
		panicEnvNotFound(key)
	}
	return c.v.GetString(key)
}

func (c *configReader) MustGetBool(key string) bool {
	if !c.v.IsSet(key) {
		panicEnvNotFound(key)
	}
	return c.v.GetBool(key)
}

func (c *configReader) MustGetInt(key string) int {
	if !c.v.IsSet(key) {
		panicEnvNotFound(key)
	}
	return c.v.GetInt(key)
}

func (c *configReader) MustGetInt64(key string) int64 {
	if !c.v.IsSet(key) {
		panicEnvNotFound(key)
	}
	return c.v.GetInt64(key)
}

func (c *configReader) MustGetStringSlice(key string) []string {
	if !c.v.IsSet(key) {
		panicEnvNotFound(key)
	}
	value := c.v.GetString(key)
	if value == "" {
		return []string{}
	}
	return strings.Split(value, ",")
}

func (c *configReader) GetStringWithDefault(key string, defaultValue string) string {
	if !c.v.IsSet(key) {
		printLogEnvUsingDefault(key, defaultValue)
		return defaultValue
	}
	return c.v.GetString(key)
}

func (c *configReader) GetBoolWithDefault(key string, defaultValue bool) bool {
	if !c.v.IsSet(key) {
		printLogEnvUsingDefault(key, defaultValue)
		return defaultValue
	}
	return c.v.GetBool(key)
}

func (c *configReader) GetIntWithDefault(key string, defaultValue int) int {
	if !c.v.IsSet(key) {
		printLogEnvUsingDefault(key, defaultValue)
		return defaultValue
	}
	return c.v.GetInt(key)
}

func (c *configReader) GetInt64WithDefault(key string, defaultValue int64) int64 {
	if !c.v.IsSet(key) {
		printLogEnvUsingDefault(key, defaultValue)
		return defaultValue
	}
	return c.v.GetInt64(key)
}

func panicEnvNotFound(key string) {
	panic(fmt.Sprintf("Environment variable `%s` not found in config.\n", key))
}

func printLogEnvUsingDefault(key string, defaultValue any) {
	fmt.Printf("Environment variable `%s` not found in config. Using default value: %v\n", key, defaultValue)
}

func printLogEnvNotFound(key string) {
	fmt.Printf("Environment variable `%s` not found in config.\n", key)
}

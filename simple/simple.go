package simple

import (
	"fmt"
	"sync"
)

type Configuration struct {
	config sync.Map
}

func New(kv map[string]interface{}) *Configuration {
	c := &Configuration{}
	for k, v := range kv {
		c.config.Store(k, v)
	}
	return c
}

func (c *Configuration) Get(key string) interface{} {
	v, ok := c.config.Load(key)
	if !ok {
		return nil
	}
	return v
}
func (c *Configuration) GetInt(key string) int {
	return 0
}
func (c *Configuration) GetBool(key string) bool {
	return false
}
func (c *Configuration) GetFloat64(key string) float64 {
	return 0.0
}
func (c *Configuration) GetString(key string) string {
	return fmt.Sprintf("%s", c.Get(key))
}
func (c *Configuration) GetIntSlice(key string) []int {
	return nil
}
func (c *Configuration) GetStringSlice(key string) []string {
	return nil
}
func (c *Configuration) GetStringMap(key string) map[string]interface{} {
	return nil
}

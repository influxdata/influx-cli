package mock

import "github.com/influxdata/influx-cli/v2/internal/config"

var _ config.Service = (*ConfigService)(nil)

type ConfigService struct {
	CreateConfigFn func(config.Config) (config.Config, error)
	DeleteConfigFn func(string) (config.Config, error)
	UpdateConfigFn func(config.Config) (config.Config, error)
	SwitchActiveFn func(string) (config.Config, error)
	ActiveFn       func() (config.Config, error)
	ListConfigsFn  func() (config.Configs, error)
}

func (c *ConfigService) CreateConfig(cfg config.Config) (config.Config, error) {
	return c.CreateConfigFn(cfg)
}
func (c *ConfigService) DeleteConfig(name string) (config.Config, error) {
	return c.DeleteConfigFn(name)
}
func (c *ConfigService) UpdateConfig(cfg config.Config) (config.Config, error) {
	return c.UpdateConfigFn(cfg)
}
func (c *ConfigService) SwitchActive(name string) (config.Config, error) {
	return c.SwitchActiveFn(name)
}
func (c *ConfigService) Active() (config.Config, error) {
	return c.ActiveFn()
}
func (c *ConfigService) ListConfigs() (config.Configs, error) {
	return c.ListConfigsFn()
}

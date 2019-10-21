package storm

import liquidweb "github.com/liquidweb/liquidweb-go"

// ConfigBackend is the interface for storm configs.
type ConfigBackend interface {
	Details(string) (*Config, error)
	List(ConfigListParams) (*ConfigList, error)
}

// ConfigClient is the API client for storm configs.
type ConfigClient struct {
	Backend liquidweb.Backend
}

// Details fetches the details for a storm config.
func (c *ConfigClient) Details(id string) (*Config, error) {
	var config Config
	params := ConfigParams{ID: id}

	err := c.Backend.Call("v1/Storm/Config/details", params, config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// List fetches a list of storm configs.
func (c *ConfigClient) List(params ConfigListParams) (*ConfigList, error) {
	configList := &ConfigList{}

	err := c.Backend.Call("v1/Storm/Config/list", params, configList)
	if err != nil {
		return nil, err
	}

	return configList, nil
}

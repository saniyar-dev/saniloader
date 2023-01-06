package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
)

func correctConfig(cfg ConfigType) (ConfigType, error) {
	if cfg.Proxy.Port == "" {
		cfg.Proxy.Port = _DEFAULTPORT
	}


	for index, backend := range cfg.Backends {
		if backend.Name == "" {
			cfg.Backends[index].Name = "container " + strconv.Itoa(index)
		}
		if backend.Id == "" {
			cfg.Backends[index].Id = cfg.Backends[index].Name
		}

		if backend.URL == "" {
			return ConfigType{}, errors.New("no url provided for some container. please correct the config file first")
		}
	}
	return cfg, nil
}

func readConfig(ConfigPath string) (ConfigType, error) {
	var cfg ConfigType
    data, err := ioutil.ReadFile(ConfigPath)
    if err != nil {
		return ConfigType{}, err
    }
    json.Unmarshal(data, &cfg)

	return correctConfig(cfg)
}
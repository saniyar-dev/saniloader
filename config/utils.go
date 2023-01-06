package config

func getCfgFile() (ConfigType, error) {
	if ConfigPath != "none" {
		return readConfig(ConfigPath)
	}
	return ConfigType{}, nil
}

func getCfgMade() (ConfigType, error) {
	if OnlyConfig {
		return ConfigType{}, nil
	} 
	return makeConfig()
}

func combineConfigs(cfg0 ConfigType, cfg1 ConfigType) ConfigType {
	var ansCfg ConfigType

	if cfg0.Proxy.Port != "" {
		ansCfg.Proxy.Port = cfg0.Proxy.Port
	} else {
		ansCfg.Proxy.Port = cfg1.Proxy.Port
	}

	ansCfg.Backends = append(cfg0.Backends, cfg1.Backends...)
	return ansCfg
}


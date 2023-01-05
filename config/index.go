package config

type ConfigType struct {
	Proxy ProxyType `json:"proxy"`
	Backends []BackendType `json:"backends"`
}

type ProxyType struct {
	Port string `json:"port"`
}

type BackendType struct {
	Name string
	Id string
	URL string `json:"url"`
}

var _DEFAULTPORT string = "3000"
var ConfigPath string = "none"
var DynamicMode bool = false
var OnlyConfig bool = false


func GetCfg() (ConfigType, error) {
	cfgFile, err := getCfgFile()
	if err != nil {
		return ConfigType{}, err
	}

	cfgMade, err := getCfgMade()
	if err != nil {
		return ConfigType{}, err
	}

	return combineConfigs(cfgFile, cfgMade), nil
}

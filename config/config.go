package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
)

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

func ReadConfig(ConfigPath string) (ConfigType, error) {
	var cfg ConfigType
    data, err := ioutil.ReadFile(ConfigPath)
    if err != nil {
		return ConfigType{}, err
    }
    json.Unmarshal(data, &cfg)

	return correctConfig(cfg)
}

func MakeConfig() (ConfigType, error){
	containersList, err := getDockerContainers()
	if err != nil {
		return ConfigType{}, err
	}

	var cfg ConfigType
	cfg.Proxy.Port = _DEFAULTPORT
	cfg.Backends = containersList
	return cfg, nil
}

func getContainerIp(containerName string) (string, error) {
	cmd := exec.Command("bash", "-c", "docker container inspect " + containerName + " --format '{{.NetworkSettings.IPAddress}}'")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return "http://" + strings.Trim(string(output), "\n") + ":" + _DEFAULTPORT, nil
}

func getDockerContainers() ([]BackendType, error) {
	cmd := exec.Command("bash", "-c", "docker ps --format '{{.ID}}\t{{.Names}}'")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return []BackendType{}, err
	}

	var ans []BackendType
	containersList := strings.Split(string(output), "\n")
	for _, containerLineStr := range containersList {
		containerDataArr := strings.Split(containerLineStr, "\t")
		if len(containerDataArr) < 2 {
			continue
		}

		containerIp, err := getContainerIp(containerDataArr[1])
		if err != nil {
			return []BackendType{}, err 
		}
		ans = append(ans, BackendType{Name: containerDataArr[1], Id: containerDataArr[0], URL: containerIp})
	}

	return ans, nil
}

func CombineConfigs (cfg0 ConfigType, cfg1 ConfigType) ConfigType {
	var ansCfg ConfigType

	if cfg0.Proxy.Port != "" {
		ansCfg.Proxy.Port = cfg0.Proxy.Port
	} else {
		ansCfg.Proxy.Port = cfg1.Proxy.Port
	}

	ansCfg.Backends = append(cfg0.Backends, cfg1.Backends...)
	return ansCfg
}
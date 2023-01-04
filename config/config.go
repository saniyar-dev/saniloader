package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

type ConfigType struct {
    Proxy    ProxyType     `json:"proxy"`
	Backends []BackendType `json:"backends"`
}

type ProxyType struct {
    Port string `json:"port"`
}

type BackendType struct {
	Id string `json:"id"`
	Name string `json:"name"`
	URL string `json:"url"`
	IsDead bool
}

var cfg ConfigType

func ReadConfig() (ConfigType, error) {
    data, err := ioutil.ReadFile("./config.json")
    if err != nil {
        log.Fatal(err.Error())
		return ConfigType{}, err
    }
    json.Unmarshal(data, &cfg)

	return CheckConfig(cfg)
}

func CheckConfig(cfg ConfigType) (ConfigType, error) {
	runningContainers, err := getDockerContainers()
	if err != nil {
		return ConfigType{}, err
	}

	var newCfg ConfigType
	newCfg.Proxy = cfg.Proxy
	for _, backend := range cfg.Backends {
		isUp := false
		for _, container := range runningContainers {
			if backend.Name == container.Name {
				isUp = true
			}
		}

		if isUp {
			newCfg.Backends = append(newCfg.Backends, BackendType{
				Name: backend.Name,
				Id: backend.Id,
				URL: backend.URL,
				IsDead: false,
			})
		}
	}

	return newCfg, nil
}

func getDockerContainers() ([]BackendType, error) {
	cmd := exec.Command("bash", "-c", `sudo docker ps --format '{{ .ID }}\t{{.Names}}'`)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return []BackendType{}, err
	}
	
	var ans []BackendType
	containersOutput := strings.Split(string(output), "\n")
	for _, containerDataLine := range containersOutput {
		containerDataArr := strings.Split(containerDataLine, "\t")
		if len(containerDataArr) < 2 {
			continue
		}
		ans = append(ans, BackendType{
			Id: containerDataArr[0], 
			Name: containerDataArr[1],
			IsDead: false,
		})
	}

	return ans, nil
}

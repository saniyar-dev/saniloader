package config

import (
	"os/exec"
	"strings"
)

func makeConfig() (ConfigType, error){
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

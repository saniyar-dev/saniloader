package config

import (
	"os/exec"
	"strings"
)

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"os/exec"
// 	"strings"
// )

// type ConfigType struct {
//     Proxy    ProxyType     `json:"proxy"`
// 	Backends []BackendType `json:"backends"`
// }

// type ProxyType struct {
//     Port string `json:"port"`
// }

// type BackendType struct {
// 	Id string
// 	Name string
// 	URL string `json:"url"`
// 	IsDead bool
// }

// var cfg ConfigType

// func ReadConfig() (ConfigType, error) {
//     data, err := ioutil.ReadFile("./config.json")
//     if err != nil {
//         log.Fatal(err.Error())
// 		return ConfigType{}, err
//     }
//     json.Unmarshal(data, &cfg)

// 	return CorrectConfig(cfg)
// }

// func CorrectConfig(cfg ConfigType) (ConfigType, error) {
// 	runningContainers, err := getDockerContainers()
// 	if err != nil {
// 		return ConfigType{}, err
// 	}

// 	var newCfg ConfigType
// 	newCfg.Proxy = cfg.Proxy
// 	for _, backend := range cfg.Backends {
// 		isUp := false
// 		for _, container := range runningContainers {
// 			if backend.Name == container.Name {
// 				isUp = true
// 			}
// 		}

// 		if isUp {
// 			newCfg.Backends = append(newCfg.Backends, BackendType{
// 				Name: backend.Name,
// 				Id: backend.Id,
// 				URL: backend.URL,
// 				IsDead: false,
// 			})
// 		}
// 	}

// 	return newCfg, nil
// }

// func getDockerContainers() ([]BackendType, error) {
// 	cmd := exec.Command("bash", "-c", `sudo docker ps --format '{{ .ID }}\t{{.Names}}'`)
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return []BackendType{}, err
// 	}

// 	var ans []BackendType
// 	containersOutput := strings.Split(string(output), "\n")
// 	for _, containerDataLine := range containersOutput {
// 		containerDataArr := strings.Split(containerDataLine, "\t")
// 		if len(containerDataArr) < 2 {
// 			continue
// 		}
// 		ans = append(ans, BackendType{
// 			Id: containerDataArr[0],
// 			Name: containerDataArr[1],
// 			IsDead: false,
// 		})
// 	}

// 	return ans, nil
// }

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

func ReadConfig() (ConfigType, error) {
	return ConfigType{}, nil
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
	cmd := exec.Command("bash", "-c", "sudo docker container inspect " + containerName + " --format '{{.NetworkSettings.IPAddress}}'")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return "http://" + strings.Trim(string(output), "\n") + ":" + _DEFAULTPORT, nil
}

func getDockerContainers() ([]BackendType, error) {
	cmd := exec.Command("bash", "-c", "sudo docker ps --format '{{.ID}}\t{{.Names}}'")
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

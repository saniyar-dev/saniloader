package config

import (
	"fmt"
	"os"
	"time"
)

func MakeConfigDynamic(configChannel chan ConfigType) {
	var cfg0, cfg1 ConfigType
	var err error
	for {
		cfg0, err = MakeConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cfg1, err = ReadConfig(ConfigPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		configChannel <- CombineConfigs(cfg0, cfg1)
		time.Sleep(5 * time.Second)
	}
}
package config

import (
	"fmt"
	"os"
	"time"
)

func MakeConfigDynamic(configChannel chan ConfigType) {
	for {
		cfg, err := GetCfg()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		configChannel <- cfg
		time.Sleep(5 * time.Second)
	}
}
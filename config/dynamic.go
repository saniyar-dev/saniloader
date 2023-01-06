package config

import (
	"fmt"
	"os"
	"time"
)

func MakeConfigDynamic() {
	for {
		cfg, err := GetCfg()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		ConfigChannel <- cfg
		time.Sleep(5 * time.Second)
	}
}
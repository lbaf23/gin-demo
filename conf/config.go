package conf

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

type config struct {
	RunMode    string
	HttpPort   string
	DBUser     string
	DBHost     string
	DBPassword string
	DBPort     string
	DBName     string
	TimeZone   string
}

var Config = &config{}

func Init() {
	if os.Getenv("RunMode") == "release" {
		fmt.Println("Load app.release.ini.")
		err := loadConfigFile("conf/app.release.ini")
		if err != nil {
			fmt.Println("Load app.release.ini failed.")
			fmt.Println("Load app.ini.")
			loadConfigFile("conf/app.ini")
		}
	} else {
		fmt.Println("Load app.debug.ini.")
		err := loadConfigFile("conf/app.debug.ini")
		if err != nil {
			fmt.Println("Load app.debug.ini failed.")
			fmt.Println("Load app.ini.")
			loadConfigFile("conf/app.ini")
		}
	}
}

func loadConfigFile(f string) error {
	var err error
	var cfg *ini.File
	cfg, err = ini.Load(f)
	if err != nil {
		return err
	}
	err = cfg.Section("config").MapTo(Config)
	if err != nil {
		return err
	}
	return nil
}

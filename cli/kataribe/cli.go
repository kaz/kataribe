package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/kaz/kataribe"
)

var configFile string
var modeGenerate bool

func init() {
	const (
		defaultConfigFile = "kataribe.toml"
		usage             = "configuration file"
	)
	flag.StringVar(&configFile, "conf", defaultConfigFile, usage)
	flag.StringVar(&configFile, "f", defaultConfigFile, usage+" (shorthand)")
	flag.BoolVar(&modeGenerate, "generate", false, "generate "+usage)
	flag.Parse()
}

func main() {
	if modeGenerate {
		f, err := os.Create(configFile)
		if err != nil {
			log.Fatal("Failed to generate "+configFile+":", err)
		}
		defer f.Close()
		_, err = f.Write([]byte(CONFIG_TOML))
		if err != nil {
			log.Fatal("Failed to write "+configFile+":", err)
		}
		os.Exit(0)
	}

	var config kataribe.Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		fmt.Println(err)
		flag.Usage()
		return
	}

	k := kataribe.New(os.Stdin, config)
	k.Run()
}

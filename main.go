package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chonla/cliargs"
	"github.com/chonla/umock/handlers/start"
	"github.com/chonla/umock/logger"
	"github.com/chonla/umock/models"
	"gopkg.in/yaml.v2"
)

var Name = "Micromock"
var Version = "Development"
var CommitID = ""

func main() {
	args, err := cliargs.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log := logger.New(logger.TRACE)
	if args.Options.Has("debug") {
		log.Level(logger.DEBUG)
		log.Debug("Turn on DEBUG mode ...")
	}

	conf := models.Config{}
	confFile := "./conf.yml"
	if args.Options.Has("conf") {
		confFile = args.Options.Get("conf")[0]
	}
	log.Debug("Use configuration file from %s\n", confFile)

	configContent, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Error("unable to read configuration from %s.\n", confFile)
		os.Exit(1)
	}

	err = yaml.Unmarshal([]byte(configContent), &conf)
	if err != nil {
		log.Error("%v\n", err)
		os.Exit(1)
	}

	switch args.Command {
	case "start":
		h, err := start.New(conf, log)
		if err != nil {
			log.Error("%v\n", err)
			os.Exit(1)
		}

		h.Start()
	case "version":
		log.Trace("%s %s(%s)", Name, Version, CommitID)
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chonla/cliargs"
	"github.com/chonla/umock/handlers/start"
	"github.com/chonla/umock/models"
	"gopkg.in/yaml.v2"
)

func main() {
	args, err := cliargs.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conf := models.Config{}
	confFile := "./conf.yml"
	if args.Options.Has("conf") {
		confFile = args.Options.Get("conf")[0]
	}

	configContent, err := ioutil.ReadFile(confFile)
	if err != nil {
		fmt.Printf("unable to read configuration from %s.\n", confFile)
		os.Exit(1)
	}

	err = yaml.Unmarshal([]byte(configContent), &conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch args.Command {
	case "start":
		h, err := start.New(conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		h.Start()
	}
}

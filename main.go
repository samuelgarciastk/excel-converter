package main

import (
	"flag"
	"fmt"
	"github.com/samuelgarciastk/excel-converter/converter"
	"github.com/samuelgarciastk/excel-converter/utils"
	"log"
	"os"
)

const configPath = "config/config.yml"

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage: see %s (cli|server) --help\n", os.Args[0])
		flag.PrintDefaults()
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "cli":
		cliParse(os.Args[2:])
	case "server":
		serverParse(os.Args[2:])
	default:
		flag.Usage()
	}
}

func cliParse(flags []string) {
	cliFlag := flag.NewFlagSet("cli", flag.ExitOnError)
	cliFlag.Usage = func() {
		_, _ = fmt.Fprintf(cliFlag.Output(), "Usage of cli: %s cli [OPTIONS]\n", os.Args[0])
		cliFlag.PrintDefaults()
	}

	cfg := cliFlag.String("c", configPath, "configuration file")
	src := cliFlag.String("s", "", "source directory")
	dst := cliFlag.String("d", "", "destination directory")
	tmpl := cliFlag.String("t", "", "template directory")

	err := cliFlag.Parse(flags)
	if err != nil {
		log.Fatal(err)
	}

	if cliFlag.Parsed() {
		config, err := utils.Load(*cfg)
		if err != nil {
			log.Fatal(err)
		}
		if *src != "" {
			config.Source = *src
		}
		if *dst != "" {
			config.Destination = *dst
		}
		if *tmpl != "" {
			config.Template = *tmpl
		}

		converter.BatchConvert(*config)
	}
}

func serverParse(flags []string) {
	serverFlag := flag.NewFlagSet("server", flag.ExitOnError)
	serverFlag.Usage = func() {
		_, _ = fmt.Fprintf(serverFlag.Output(), "Usage of server: %s server [OPTIONS] (start|stop|restart)\n", os.Args[0])
		serverFlag.PrintDefaults()
	}

	cfg := serverFlag.String("c", configPath, "configuration file")
	operation := serverFlag.Arg(0)

	err := serverFlag.Parse(flags)
	if err != nil {
		log.Fatal(err)
	}

	if serverFlag.Parsed() {
		if operation == "" {
			serverFlag.Usage()
			os.Exit(1)
		}
		config, err := utils.Load(*cfg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Server config: %v", config)
	}
}

package main

import (
	"flag"
	"fmt"
	"github.com/samuelgarciastk/excel-converter/converter"
	"github.com/samuelgarciastk/excel-converter/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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
		_, _ = fmt.Fprintf(serverFlag.Output(), "Usage of server: %s server [OPTIONS]\n", os.Args[0])
		serverFlag.PrintDefaults()
	}

	cfg := serverFlag.String("c", configPath, "configuration file")

	err := serverFlag.Parse(flags)
	if err != nil {
		log.Fatal(err)
	}

	if serverFlag.Parsed() {
		config, err := utils.Load(*cfg)
		if err != nil {
			log.Fatal(err)
		}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			convert(w, r, *config)
		})
		log.Printf("Starting server...")
		err = http.ListenAndServe(":"+strconv.Itoa(config.ServerPort), nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func convert(w http.ResponseWriter, r *http.Request, config utils.Config) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		_ = r.ParseMultipartForm(32 << 20)
		file, header, err := r.FormFile("file")
		if err != nil {
			_, _ = fmt.Fprintf(w, "Error retrieving the file.\n")
			return
		}
		defer func() {
			if fileErr := file.Close(); err == nil {
				err = fileErr
			}
		}()

		tempFile, err := ioutil.TempFile(os.TempDir(), header.Filename)
		if err != nil {
			_, _ = fmt.Fprintf(w, "%v\n", err)
			return
		}

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			_, _ = fmt.Fprintf(w, "%v\n", err)
			return
		}
		if _, err = tempFile.Write(fileBytes); err != nil {
			_, _ = fmt.Fprintf(w, "%v\n", err)
			return
		}
		if err = tempFile.Close(); err != nil {
			_, _ = fmt.Fprintf(w, "%v\n", err)
			return
		}
		log.Printf("Successfully uploaded file.")

		template, err := utils.ReadTemplate(config.Template)
		if err != nil {
			_, _ = fmt.Fprintf(w, "cannot read template file: %s, due to %v\n", config.Template, err)
			return
		}

		err = converter.ConvertFile(tempFile.Name(), filepath.Join(config.Destination, header.Filename), *template)
		if err != nil {
			_, _ = fmt.Fprintf(w, "cannot convert file: %s, due to %v", tempFile.Name(), err)
			return
		}
		_, _ = fmt.Fprintf(w, "Convert file [%s] successfully.", header.Filename)
	default:
		_, _ = fmt.Fprintf(w, "Sorry, only POST method is supported.\n")
	}
}

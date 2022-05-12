package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	yamlpatch "github.com/fox-md/yaml-patch"
	"github.com/rs/zerolog"
)

func main() {
	var filePath string
	var patchLine string
	var logLevel string
	var outputFile string

	flag.StringVar(&filePath, "file", "test.yaml", "file to patch")
	flag.StringVar(&outputFile, "outfile", "patched.yaml", "file to output file")
	flag.StringVar(&patchLine, "patch", "", "';' separated list of patches")
	flag.StringVar(&logLevel, "loglevel", "info", "logging level (info|debug)")
	flag.Parse()

	if len(patchLine) == 0 {
		log.Fatalf("patch not specified")
	}

	if strings.ToLower(logLevel) == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else if strings.ToLower(logLevel) == "info" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		log.Fatalf("log level must be set to 'info' or 'debug'")
	}

	yfile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	bs := yamlpatch.PatchFile(patchLine, yfile)

	fmt.Println(string(bs))

	err = ioutil.WriteFile(outputFile, bs, 0777)

	if err != nil {
		log.Fatal(err)
	}

}

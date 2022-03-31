package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"flag"
	yamlpatch "github.com/fox-md/yaml-patch"
)

func main() {
	var filePath string
	var patchLine string
	var logLevel string

	flag.StringVar(&filePath, "file", "test.yaml", "file to patch")
	flag.StringVar(&patchLine, "patch", "", "';' separated list of patches")
	flag.StringVar(&logLevel, "loglevel", "info", "logging level (info|debug)")
	flag.Parse()

	yfile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	
	bs := yamlpatch.PatchFile(patchLine, yfile)

	fmt.Println(string(bs))

	err = ioutil.WriteFile("patched.yaml", bs, 0777)

	if err != nil {
		log.Fatal(err)
	}

}

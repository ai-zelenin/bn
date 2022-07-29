package main

import (
	"fmt"
	"os"

	"github.com/xuri/xgen"
)

func main() {
	files, err := xgen.GetFileList("./spec/xsd")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		opts := &xgen.Options{
			FilePath:  file,
			InputDir:  "./spec/xsd",
			OutputDir: "",

			Lang:                "Go",
			Package:             "spec",
			IncludeMap:          make(map[string]bool),
			LocalNameNSMap:      make(map[string]string),
			NSSchemaLocationMap: make(map[string]string),
			ParseFileList:       make(map[string]bool),
			ParseFileMap:        make(map[string][]interface{}),
			ProtoTree:           make([]interface{}, 0),
			RemoteSchema:        make(map[string][]byte),
		}
		err = xgen.NewParser(opts).Parse()
		if err != nil {
			fmt.Printf("process error on %s: %s\r\n", file, err.Error())
			os.Exit(1)
		}
	}
	fmt.Println("done")
}

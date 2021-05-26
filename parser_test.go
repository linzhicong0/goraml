package goraml

import (
	"bytes"
	"fmt"
	"goraml/processor"
	"path/filepath"
	"testing"
)

func TestParseRamlFile(t *testing.T) {
	filePath := "raml_example/example.raml"

	api, err := ParseRamlFile(filePath)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v",api)


}
func TestPreProcess(t *testing.T) {

	filePath := "raml_example/example.raml"

	// Get the working directory
	workingDirectory, _ := filepath.Split(filePath)

	// Get the raw content of the raml file
	rawContent, err := processor.ReadFileContent(filePath)
	if err != nil {
		t.Fatal(err)
	}

	rawContentBuffer := bytes.NewBuffer(rawContent)

	// Do some pre-process jobs, like adding the include file's content into the raml file
	preProcessedContents, err := processor.PreProcess(rawContentBuffer, workingDirectory)
	fmt.Println(preProcessedContents)
}
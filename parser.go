package goraml

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"goraml/processor"
	"goraml/ramlTypes"
	"path/filepath"
)

// ParseRamlFile parse the raml file into API struct
func ParseRamlFile(ramlFile string) (*ramlTypes.API, error) {

	// Get the working directory
	workingDirectory, _ := filepath.Split(ramlFile)

	// Get the raw content of the raml file
	rawContent, err := processor.ReadFileContent(ramlFile)
	if err != nil {
		return nil, err
	}

	rawContentBuffer := bytes.NewBuffer(rawContent)

	// Do some pre-process jobs, like adding the include file's content into the raml file
	preProcessedContents, err := processor.PreProcess(rawContentBuffer, workingDirectory)

	fmt.Println(string(preProcessedContents))

	if err != nil {
		return nil, fmt.Errorf("error while processing the raml file, error: %s", err.Error())
	}

	return Unmarshal(preProcessedContents)
}

// Unmarshal parse the given content into API struct
func Unmarshal(rawContent []byte) (*ramlTypes.API, error) {

	result := &ramlTypes.API{}
	err := yaml.Unmarshal(rawContent, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

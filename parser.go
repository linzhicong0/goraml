package goraml

import (
	"bufio"
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"goraml/ramlTypes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const includeSymbol = "!include"

// ParseRamlFile parse the raml file into API struct
func ParseRamlFile(ramlFile string) (*ramlTypes.API, error) {

	// Get the working directory
	workingDirectory, _ := filepath.Split(ramlFile)

	// Get the raw content of the raml file
	rawContent, err := readFileContent(ramlFile)
	if err != nil {
		return nil, err
	}

	rawContentBuffer := bytes.NewBuffer(rawContent)

	// Do some pre-process jobs, like adding the include file's content into the raml file
	preProcessedContents, err := PreProcess(rawContentBuffer, workingDirectory)

	if err != nil {
		return nil, fmt.Errorf("error while processing the raml file, error: %s", err.Error())
	}

	return Unmarshal(preProcessedContents)
}

// read all the content in the given file
func readFileContent(filePath string) ([]byte, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found: %s", filePath)
	}
	return ioutil.ReadFile(filePath)
}

// PreProcess write the content into the raml file before parsing
// For now it just only write the included file into the
// raml file
func PreProcess(rawContent io.Reader, workingDirectory string) ([]byte, error) {

	var processedContentBuffer bytes.Buffer

	scanner := bufio.NewScanner(rawContent)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		// Find the !include
		if index := strings.Index(line, includeSymbol); index != -1 {
			symbleLength := len(includeSymbol)

			processedContentBuffer.WriteString(line[:index])
			// make it as: example: |
			processedContentBuffer.WriteString(" |\n")
			//Get the file path
			includedFilePath := strings.Trim(line[index+symbleLength:], " ")
			includedFileContent, err := readFileContent(filepath.Join(workingDirectory, includedFilePath))
			if err != nil {
				return nil, err
			}
			includedFileContentScanner := bufio.NewScanner(bytes.NewBuffer(includedFileContent))

			// Set the format with indent
			trimed := strings.Trim(line[:index], " ")
			startIndex := strings.Index(line, trimed) + 2
			indentString := strings.Repeat(" ", startIndex)
			// write the content
			for includedFileContentScanner.Scan() {
				contentLine := includedFileContentScanner.Text()
				// write the indent first
				processedContentBuffer.WriteString(indentString)
				// write the actual content
				processedContentBuffer.WriteString(contentLine)
				processedContentBuffer.WriteByte('\n')

			}

		} else {
			// Common line, just write it
			processedContentBuffer.WriteString(line)
			processedContentBuffer.WriteByte('\n')

		}
	}

	return processedContentBuffer.Bytes(), nil

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

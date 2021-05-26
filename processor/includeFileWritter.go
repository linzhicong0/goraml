package processor

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const includeSymbol = "!include"

// read all the content in the given file
func ReadFileContent(filePath string) ([]byte, error) {
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

			//processedContentBuffer.WriteString(line[:index])
			// make it as: example: |

			//Get the file path
			includedFilePath := strings.Trim(line[index+symbleLength:], " ")

			// TODO: according to different file type, using the different writer
			fileExt := path.Ext(includedFilePath)
			var includeFileWriter IncludeFileWriter

			switch fileExt {
			case ".json":
				includeFileWriter = new(JsonWriter)
			case ".raml":
				includeFileWriter = new(RamlWriter)
			}

			includeFileWriter.Write(line, index, &processedContentBuffer, filepath.Join(workingDirectory, includedFilePath))

		} else {
			// Common line, just write it
			processedContentBuffer.WriteString(line)
			processedContentBuffer.WriteByte('\n')

		}
	}

	return processedContentBuffer.Bytes(), nil

}

type IncludeFileWriter interface {
	Write(lineText string, index int, buffer *bytes.Buffer, includedFilePath string) error
}

// RamlWriter 目前先假设raml中的datatype和trait等是相同的写入方式
type RamlWriter struct{}
type JsonWriter struct{}

func (r *RamlWriter) Write(lineText string, index int, buffer *bytes.Buffer, includedFilePath string) error {
	// TODO: append the raml content into the main file


	trimedType := strings.Trim(lineText[:index], " ")
	startIndex := strings.Index(lineText, trimedType)

	// To see if the key is type
	// If not type, then treat it as normal
	//write the key name and change one line and write the indent
	if trimedType != "type:" {
		buffer.WriteString(lineText[:index])
		// Change line
		buffer.WriteString("\n")
		startIndex += 2
	}


	indentString := strings.Repeat(" ", startIndex)

	// TODO: if this include file also include other file, need to do the preprocess first

	// Get the raw content of the raml file
	rawContent, err := ReadFileContent(includedFilePath)
	if err != nil {
		return err
	}

	workingDirectory, _ := filepath.Split(includedFilePath)
	rawContentBuffer := bytes.NewBuffer(rawContent)
	includeFileContent, err := PreProcess(rawContentBuffer, workingDirectory)
	if err != nil {
		return err
	}
	includedFileContentScanner := bufio.NewScanner(bytes.NewBuffer(includeFileContent))

	var ramlFileType string
	// Get the first line content
	// Check if the file start with #%RAML and get the raml file type
	if includedFileContentScanner.Scan() {
		firstLineContent := includedFileContentScanner.Text()
		if firstLineContent[:6] != "#%RAML" {
			return fmt.Errorf("error file type: %s", firstLineContent[:5])
		}

		ramlFileType = firstLineContent[11:19]

	}

	switch ramlFileType {
	case "DataType":
		// write the content
		for includedFileContentScanner.Scan() {

			contentLine := includedFileContentScanner.Text()
			// write the indent first
			buffer.WriteString(indentString)
			// write the actual content
			buffer.WriteString(contentLine)
			buffer.WriteByte('\n')

		}

	}

	return nil
}

func (j *JsonWriter) Write(lineText string, index int, buffer *bytes.Buffer, includedFilePath string) error {

	buffer.WriteString(lineText[:index])
	buffer.WriteString(" |\n")
	// Set the format with indent
	trimed := strings.Trim(lineText[:index], " ")
	startIndex := strings.Index(lineText, trimed) + 2
	indentString := strings.Repeat(" ", startIndex)

	// Read the include file content
	includeFileContent, err := ReadFileContent(includedFilePath)
	if err != nil {
		return err
	}
	includedFileContentScanner := bufio.NewScanner(bytes.NewBuffer(includeFileContent))
	// write the content
	for includedFileContentScanner.Scan() {
		contentLine := includedFileContentScanner.Text()
		// write the indent first
		buffer.WriteString(indentString)
		// write the actual content
		buffer.WriteString(contentLine)
		buffer.WriteByte('\n')

	}
	return nil
}

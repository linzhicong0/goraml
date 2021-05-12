package goraml

import (
	"fmt"
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
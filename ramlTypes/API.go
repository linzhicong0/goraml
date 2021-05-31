package ramlTypes

import "fmt"

// API is the main struct for describing the API
type API struct {
	Title         string               `yaml:"title"`
	RAMLVersion   string               `yaml:"raml_version"`
	Version       string               `yaml:"version"`
	Resources     map[string]Resource  `yaml:",inline"`
	MediaType     string               `yaml:"mediaType"`
	BaseUri       string               `yaml:"baseUri"`
	Types         map[string]NameTyped `yaml:"types"`
	Documentation Documentation        `yaml:"documentation"`
}

// GetType Get the named typed from the given type name
func (api *API) GetType(typeName string) (*NameTyped, error) {
	if value, ok := api.Types[typeName]; ok {
		return &value, nil
	} else {
		return nil, fmt.Errorf("can not find type: %s", typeName)
	}
}

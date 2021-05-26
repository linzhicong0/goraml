package ramlTypes

// API is the main struct for describing the API
type API struct {
	Title         string              `yaml:"title"`
	RAMLVersion   string              `yaml:"raml_version"`
	Version       string              `yaml:"version"`
	Resources     map[string]Resource `yaml:",inline"`
	MediaType     string              `yaml:"mediaType"`
	BaseUri       string              `yaml:"baseUri"`
	Types         map[string]NameTyped   `yaml:"types"`
	Documentation Documentation       `yaml:"documentation"`
}

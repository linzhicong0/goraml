package ramlTypes

type NamedParameters struct {
	Required    bool                       `yaml:"required"`
	Description string                     `yaml:"description"`
	Type        string                     `yaml:"type"`
	DisplayName string                     `yaml:"displayName"`
	Properties  map[string]NamedParameters `yaml:"properties"`
	Example     interface{}                `yaml:"example"`
}

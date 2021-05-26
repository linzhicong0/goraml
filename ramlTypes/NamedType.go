package ramlTypes

type NameTyped struct {
	Type       string                     `yaml:"type"`
	Properties map[string]NamedParameters `yaml:"properties"`
}

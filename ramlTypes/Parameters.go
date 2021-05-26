package ramlTypes

// Parameters describe the parameters for request body, query parameters and headers
type Parameters struct {
	DisplayName string                `yaml:"displayName"`
	Type        string                `yaml:"type"`
	Description string                `yaml:"description"`
	Example     interface{}           `yaml:"example"`
	Enum        []interface{}         `yaml:"enum,flow"`
	Required    bool                  `yaml:"required"`
	Default     interface{}           `yaml:"default"`
	// The parameter may be an object, so it would have properties
	Properties  map[string]Parameters `yaml:"properties"`
}

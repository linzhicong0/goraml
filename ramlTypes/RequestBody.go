package ramlTypes
// RequestBody is only for post, put, patch, which is describing a request body for a request
type RequestBody struct {
	Type       string                `yaml:"type"`
	Properties map[string]Parameters `yaml:"properties"`
	Example    string                `yaml:"example"`
}

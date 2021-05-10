package ramlTypes

type ContentType string
type HTTPStatusCode int

// Response describe the response from a resource
type Response struct {
	Description string                     `yaml:"description"`
	Headers     map[string]Parameters      `yaml:"headers"`
	Body        map[ContentType]Parameters `yaml:"body"`
}

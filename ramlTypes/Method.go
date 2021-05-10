package ramlTypes
// Method is for describing the method inside a resource
type Method struct {
	DisplayName     string                      `yaml:"displayName"`
	Description     string                      `yaml:"description"`
	QueryParameters map[string]Parameters       `yaml:"queryParameters"`
	Headers         map[string]Parameters       `yaml:"headers"`
	Responses       map[HTTPStatusCode]Response `yaml:"responses"`
	RequestBody     map[ContentType]RequestBody `yaml:"body"`
}

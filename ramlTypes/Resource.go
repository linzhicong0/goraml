package ramlTypes
// Resource is for describing a resource in the API
type Resource struct {
	Description string              `yaml:"description"`
	DisplayName string              `yaml:"displayName"`
	Get         *Method             `yaml:"get"`
	Put         *Method             `yaml:"put"`
	Post        *Method             `yaml:"post"`
	Patch       *Method             `yaml:"patch"`
	Delete      *Method             `yaml:"delete"`
	SubResource map[string]Resource `yaml:",inline"`
}

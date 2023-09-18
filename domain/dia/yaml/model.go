package yaml

type Model struct {
	Name        string                   `yaml:"name"`
	Fields      map[string]Field         `yaml:"fields"`
	Identifiers map[string]Identifier    `yaml:"identifiers"`
	Related     map[string]ModelRelation `yaml:"related"`
}

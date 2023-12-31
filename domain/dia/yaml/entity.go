package yaml

type Entity struct {
	Name    string                    `yaml:"name"`
	Fields  map[string]Field          `yaml:"fields"`
	Related map[string]EntityRelation `yaml:"related"`
}

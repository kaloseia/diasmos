package yaml

import "log"

type Model struct {
	Name        string                   `yaml:"name"`
	Fields      map[string]Field         `yaml:"fields"`
	Identifiers map[string]Identifier    `yaml:"identifiers"`
	Related     map[string]ModelRelation `yaml:"related"`
}

func (m *Model) GetIdentifierFields() []Field {
	var fields []Field
	for _, identifier := range m.Identifiers {
		for _, fieldName := range identifier.Fields {
			idField, fieldExists := m.Fields[fieldName]
			if !fieldExists {
				log.Printf("identifier field '%s' does not exist in model '%s'", fieldName, m.Name)
				continue
			}
			fields = append(fields, idField)
		}
	}
	return fields
}

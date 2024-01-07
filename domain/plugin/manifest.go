package plugin

import "encoding/json"

type Manifest struct {
	ID          ID     `json:"id"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type manifestInternal struct {
	ID

	Author      string `json:"author"`
	Description string `json:"description"`
}

func (m *Manifest) UnmarshalJSON(data []byte) error {
	var internal manifestInternal
	if internalErr := json.Unmarshal(data, &internal); internalErr != nil {
		return internalErr
	}

	m.ID = internal.ID
	m.Author = internal.Author
	m.Description = internal.Description
	return nil
}

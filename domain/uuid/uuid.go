package uuid

import (
	"github.com/google/uuid"
)

// Mocks allows for uuid initialization test overrides as a queue
var Mocks = []string{}

func ResetMocks() {
	Mocks = []string{}
}

func AddMock(value string) {
	Mocks = append(Mocks, value)
}

func New() string {
	if len(Mocks) == 0 {
		return uuid.NewString()
	}
	nextUUID, remainingUUIDs := Mocks[0], Mocks[1:]
	Mocks = remainingUUIDs
	return nextUUID
}

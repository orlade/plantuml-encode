package plantuml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	source := `@startuml
Bob -> Alice : hello
@enduml`

	expected := "SYWkIImgAStDuNBAJrBGjLDmpCbCJbMmKiX8pSd9vt98pKifpSq11000__y0"

	encoded, err := DeflateAndEncode(source)

	assert.NoError(t, err)
	assert.Equal(t, expected, encoded)
}

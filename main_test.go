package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	source := `@startuml
Bob -> Alice : hello
@enduml`

	expected := "UDfoA2v9B2efpStXSifFKj2rKt3CoKnELR1Io4ZDoSddSaZDIodDpG44003___W93C00"

	encoded, err := deflateAndEncode([]byte(source))

	assert.NoError(t, err)
	assert.Equal(t, expected, encoded)
}

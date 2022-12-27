package puml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlantuml_StartUML(t *testing.T) {
	assert.NotNil(t, NewPlantUML().StartUML())
}

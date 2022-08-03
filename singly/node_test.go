package singly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeString(t *testing.T) {
	n1 := Node[string]{Data: "sample data"}
	n2 := Node[float64]{Data: 3.14}

	assert.Equal(t, n1.String(), "sample data")
	assert.Equal(t, n2.String(), "3.14")
}

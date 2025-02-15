package strain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	sample := `2
Gallium Arsenide diamond structure
Ga -0.03337275954690459 0.01235894777016857 0.06825325671325802
As 0.2779642058165548 0.2947427293064877 0.3003355704697987
---
X refaxis: [0, 0.7071067811865475, 1]
Y ref axis [0, -0.7071067811865475, 1]
Z refaxis [0.4082482904638631, 0.8164965809277261, 1.2247448713915892]
`

	_, err := Decode(sample)

	assert.NoError(t, err)
}

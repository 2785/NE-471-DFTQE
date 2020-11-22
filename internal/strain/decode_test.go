package strain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	sample := `2
Gallium Arsenide diamond structure
Ga 0.0 0.0 0.0
As 0.24999999999459144 0.25000000005710166 0.2500000001185888
---
X refaxis: [0, 0.7071067811865475, 1]	Y ref axis [0, -0.7071067811865475, 1]	Z refaxis [0.4082482904638631, 0.8164965809277261, 1.2247448713915892]
`

	_, err := Decode(sample)

	assert.NoError(t, err)
}

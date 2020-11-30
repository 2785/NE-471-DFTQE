package nanohub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const test_dos string = `Energy (eV) (#1), Density of States (#1)
                  -6,                 2.12
                -5.9,                 1.91
                -5.8,               0.1779
                -5.7,                    0
                -5.6,                    0
                -5.5,                    0
                -5.4,                    0
                -5.3,                    0
                -5.2,                    0
                -5.1,                    0
                  -5,                    0
                -4.9,                    0
                -4.8,                    0
                -4.7,                    0
                -4.6,                    0
                -4.5,                    0
                -4.4,                    0
                -4.3,                    0
                -4.2,                    0
                -4.1,                    0
                  -4,                    0
                -3.9,                    0
                -3.8,                    0
                -3.7,                    0
                -3.6,                    0
                -3.5,                    0
                -3.4,                    0
                -3.3,                    0
                -3.2,                    0
                -3.1,                    0
                  -3,                    0
                -2.9,                    0
                -2.8,             0.001808
                -2.7,             0.009704
                -2.6,              0.03898
                -2.5,               0.1133
                -2.4,               0.2617
                -2.3,               0.5332
                -2.2,               0.9111
                -2.1,                1.465
                  -2,                1.454
                -1.9,                1.261
                -1.8,                1.128
                -1.7,                1.018
                -1.6,               0.8324
                -1.5,               0.7333
                -1.4,               0.6565
                -1.3,               0.5988
                -1.2,               0.5511
                -1.1,               0.5086
                  -1,               0.4705
                -0.9,               0.4371
                -0.8,                0.411
                -0.7,               0.3913
                -0.6,                0.376
                -0.5,               0.3634
                -0.4,                0.352
                -0.3,               0.3416
                -0.2,               0.3331
                -0.1,               0.3274
                   0,               0.3277
                 0.1,               0.3321
                 0.2,               0.3446
                 0.3,               0.3625
                 0.4,               0.3679
                 0.5,               0.4097
                 0.6,               0.5245
                 0.7,                0.669
                 0.8,               0.7932
                 0.9,               0.8783
                   1,               0.9654
                 1.1,                1.075
                 1.2,                1.197
                 1.3,                1.286
                 1.4,                1.385
                 1.5,                 1.36
                 1.6,                1.259
                 1.7,                1.199
                 1.8,                1.152
                 1.9,                1.145
                   2,                1.147
                 2.1,                1.165
                 2.2,                1.136
                 2.3,                1.115
                 2.4,                1.105
                 2.5,                1.099
                 2.6,                1.083
                 2.7,                1.061
                 2.8,                1.044
                 2.9,                1.015
                   3,                1.062
                 3.1,               0.9813
                 3.2,               0.9492
                 3.3,               0.9103
                 3.4,               0.8752
                 3.5,               0.8242
                 3.6,               0.7803
                 3.7,               0.7235
                 3.8,               0.6576
                 3.9,               0.6139
                   4,               0.5907
                 4.1,               0.5808
                 4.2,                0.576
                 4.3,                0.571
                 4.4,                0.565
                 4.5,               0.5727
                 4.6,                0.593
                 4.7,               0.6078
                 4.8,               0.6018
                 4.9,               0.5819
                   5,               0.5542
                 5.1,                 0.49
                 5.2,               0.4262
                 5.3,               0.3278
                 5.4,                0.203
                 5.5,               0.1069
                 5.6,               0.1655
                 5.7,               0.2311
                 5.8,               0.3038
                 5.9,               0.3855
                   6,                0.458
                 6.1,               0.5662
                 6.2,               0.6861
                 6.3,               0.8071
                 6.4,               0.9264
                 6.5,                1.079
                 6.6,                1.198
                 6.7,                1.411
                 6.8,                1.662
                 6.9,                1.913
                   7,                2.026
                 7.1,                1.914
                 7.2,                1.659
                 7.3,                1.321
                 7.4,                1.017
                 7.5,               0.7139
                 7.6,               0.4696
                 7.7,               0.4062
                 7.8,               0.3778
                 7.9,               0.4106
                   8,               0.5591
                 8.1,               0.7432
                 8.2,                1.043
                 8.3,                1.274
                 8.4,                1.474
                 8.5,                1.949
                 8.6,                2.065
                 8.7,                1.822
                 8.8,                1.575
                 8.9,                1.306
                   9,                1.063
                 9.1,               0.8479
                 9.2,               0.7193
                 9.3,               0.6474
                 9.4,               0.6013
                 9.5,               0.5943
                 9.6,               0.5082
                 9.7,               0.5163
                 9.8,               0.5104
                 9.9,               0.4864
                  10,                0.525

------------------------------------------------------------
 Fermi level
------------------------------------------------------------
Energy (eV) (#1), Density of States (#1)
              5.4013,                    0
              5.4013,                 2.12
`

const test_band string = `               0.866,               -6.164
              1.0392,               -6.456
              1.2124,               -7.005
              1.3856,               -7.484
              1.5588,               -7.794
              1.7321,                 -7.9
              1.9321,               -7.768
              2.1321,               -7.388
              2.3321,               -6.796
              2.5321,               -6.136
              2.7321,               -5.787

------------------------------------------------------------
 Band 2
------------------------------------------------------------
K-Point (#1), Energy (eV) (#1)
              2.7321,               -2.326
              2.5321,               -1.812
              2.3321,               -0.665
              2.1321,                0.747
              1.9321,                2.259
              1.7321,                3.266
              1.5588,                2.178
              1.3856,                0.489
              1.2124,               -1.106
              1.0392,               -2.359
               0.866,               -2.892

------------------------------------------------------------
 Band 3
------------------------------------------------------------
K-Point (#1), Energy (eV) (#1)
               0.866,                2.795
              1.0392,                2.884
              1.2124,                3.139
              1.3856,                3.514
              1.5588,                3.898
              1.7321,                4.122
              1.9321,                 3.73
              2.1321,                3.027
              2.3321,                2.381
              2.5321,                1.952
              2.7321,                1.801

------------------------------------------------------------
 Band 4
------------------------------------------------------------
K-Point (#1), Energy (eV) (#1)
              2.7321,                3.287
              2.5321,                3.414
              2.3321,                3.792
              2.1321,                4.377
              1.9321,                4.942
              1.7321,                4.874
              1.5588,                5.202
              1.3856,                5.225
              1.2124,                5.173
              1.0392,                5.142
               0.866,                5.134

------------------------------------------------------------
 Band 5
------------------------------------------------------------
K-Point (#1), Energy (eV) (#1)
               0.866,                 6.79
              1.0392,                 6.86
              1.2124,                7.072
              1.3856,                7.233
              1.5588,                6.588
              1.7321,                5.992
              1.9321,                 6.66
              2.1321,                7.111
              2.3321,                6.326
              2.5321,                 5.43
              2.7321,                5.076

------------------------------------------------------------
 Band 6
------------------------------------------------------------
K-Point (#1), Energy (eV) (#1)
              2.7321,                 6.32
              2.5321,                 6.51
              2.3321,                7.149
              2.1321,                8.349
              1.9321,                9.023
              1.7321,                8.854
              1.5588,                9.054
              1.3856,                8.927
              1.2124,                 8.62
              1.0392,                8.154
               0.866,                7.959

------------------------------------------------------------
 Band 7
------------------------------------------------------------
K-Point (#1), Energy (eV) (#1)
               0.866,               11.388
              1.0392,               11.545
              1.2124,               11.396
              1.3856,               10.273
              1.5588,                9.433
              1.7321,                9.766
              1.9321,                9.495
              2.1321,               10.093
              2.3321,               11.884
              2.5321,               13.878
              2.7321,               15.096

------------------------------------------------------------
 Band 8
------------------------------------------------------------
K-Point (#1), Energy (eV) (#1)
              2.7321,               15.196
              2.5321,               14.123
              2.3321,               12.407
              2.1321,               10.817
              1.9321,               10.332
              1.7321,               10.259
              1.5588,               10.562
              1.3856,                10.72
              1.2124,               11.788
              1.0392,               11.824
               0.866,               11.758

------------------------------------------------------------
 High Symmetry Point 0.5 0.5 0.5
------------------------------------------------------------
K-Point (#1), Energy (eV) (#1)
               0.866,                 -7.9
               0.866,               15.196

------------------------------------------------------------
 High Symmetry Point 0.0 0.0 0.0
------------------------------------------------------------
K-Point (#1), Energy (eV) (#1)
              1.7321,                 -7.9
              1.7321,               15.196

------------------------------------------------------------
 High Symmetry Point 1.0 0.0 0.0
------------------------------------------------------------
K-Point (#1), Energy (eV) (#1)
              2.7321,                 -7.9
              2.7321,               15.196

------------------------------------------------------------
 Fermi level: 5.4013
------------------------------------------------------------
K-Point, Energy (eV)
               0.866,               5.4013
              2.7321,               5.4013

`

func TestDecodeDoS(t *testing.T) {
	_, err := DecodeDoS(test_dos)
	assert.NoError(t, err)
}

func TestDecodeBands(t *testing.T) {
	_, err := DecodeBands(test_band)

	assert.NoError(t, err)

	t.Fail()
}

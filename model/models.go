package model

// Input .
type Input struct {
	Ga, As, XRef, YRef, ZRef Coord
}

// Coord .
type Coord struct {
	X, Y, Z float64
}

// DoSEntry .
type DoSEntry struct {
	Energy  float32
	Density float32
}

// DoSInfo .
type DoSInfo struct {
	DoS        []DoSEntry
	FermiLevel []DoSEntry
}

// BandEntry .
type BandEntry struct {
	K      float32
	Energy float32
}

// BandSymmetry .
type BandSymmetry struct {
	S0, S0_5, S1 []BandEntry
}

// BandInfo .
type BandInfo struct {
	Bands      [][]BandEntry
	Symmetry   BandSymmetry
	FermiLevel []BandEntry
}

// Simulation .
type Simulation struct {
	Input Input
	DoS   DoSInfo
	Bands BandInfo
}

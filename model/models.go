package model

import "sort"

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
	Energy  float64
	Density float64
}

// DoSInfo .
type DoSInfo struct {
	DoS        []DoSEntry
	FermiLevel []DoSEntry
}

// BandEntry .
type BandEntry struct {
	K      float64
	Energy float64
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

// BandGap ...
func (sim *Simulation) BandGap() float64 {
	hi := []float64{}
	lo := []float64{}
	for _, v := range sim.Bands.Bands {
		for _, vi := range v {
			if vi.Energy > sim.DoS.FermiLevel[0].Energy {
				hi = append(hi, vi.Energy)
			} else {
				lo = append(lo, vi.Energy)
			}
		}
	}

	sort.Float64s(hi)
	sort.Float64s(lo)

	return hi[0] - lo[len(lo)-1]
}

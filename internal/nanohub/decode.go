package nanohub

import (
	"encoding/csv"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/2785/n471-proj-carrot/model"
)

// DecodeDoS .
func DecodeDoS(s string) (*model.DoSInfo, error) {

	parts := strings.Split(s, `------------------------------------------------------------
 Fermi level
------------------------------------------------------------`)

	if len(parts) != 2 {
		return nil, errors.New("unexpected number of parts")
	}

	dosPart := csv.NewReader(strings.NewReader(parts[0]))

	dosRecords, err := dosPart.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("error reading part: %w", err)
	}

	fermiPart := csv.NewReader(strings.NewReader(parts[1]))

	fermiRecords, err := fermiPart.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("error reading part: %w", err)
	}

	dosRecords = dosRecords[1:]

	fermiRecords = fermiRecords[1:]

	out := &model.DoSInfo{
		DoS:        make([]model.DoSEntry, len(dosRecords)),
		FermiLevel: make([]model.DoSEntry, len(fermiRecords)),
	}

	for i, v := range dosRecords {
		e, err := strconv.ParseFloat(strings.TrimSpace(v[0]), 32)
		if err != nil {
			return nil, fmt.Errorf("could not convert to float: %w", err)
		}
		d, err := strconv.ParseFloat(strings.TrimSpace(v[1]), 32)
		if err != nil {
			return nil, fmt.Errorf("could not convert to float: %w", err)
		}

		out.DoS[i] = model.DoSEntry{Energy: float32(e), Density: float32(d)}
	}

	for i, v := range fermiRecords {
		e, err := strconv.ParseFloat(strings.TrimSpace(v[0]), 32)
		if err != nil {
			return nil, fmt.Errorf("could not convert to float: %w", err)
		}
		d, err := strconv.ParseFloat(strings.TrimSpace(v[1]), 32)
		if err != nil {
			return nil, fmt.Errorf("could not convert to float: %w", err)
		}

		out.FermiLevel[i] = model.DoSEntry{Energy: float32(e), Density: float32(d)}
	}

	return out, nil
}

var symmSplitter = regexp.MustCompile(`--+\n High Symmetry Point 0\.5.*\n--+\nK-P.+\n`)

var fermiSplitter = regexp.MustCompile(`--+\n Fermi.*\n--+\nK-P.+\n`)

var bandSplitter = regexp.MustCompile(`--+\n Band.*\n--+\nK-P.+\n`)

var symmRegSplitter = regexp.MustCompile(`--+\n High Symmetry.*\n--+\nK-P.+\n`)

// DecodeBands .
func DecodeBands(s string) (*model.BandInfo, error) {
	firstSplit := symmSplitter.Split(s, -1)
	if len(firstSplit) != 2 {
		return nil, errors.New("unexpected number of segments")
	}

	secondSplit := fermiSplitter.Split(firstSplit[1], -1)

	if len(secondSplit) != 2 {
		return nil, errors.New("unexpected number of segments")
	}

	bands := bandSplitter.Split(firstSplit[0], -1)
	symms := symmRegSplitter.Split(secondSplit[0], -1)
	fermi := secondSplit[1]

	if len(bands) == 1 || len(symms) == 1 || len(symms) != 3 {
		return nil, errors.New("no match found")
	}

	out := &model.BandInfo{
		Bands: make([][]model.BandEntry, len(bands)),
	}

	for i, band := range bands {
		c, err := csv.NewReader(strings.NewReader(band)).ReadAll()
		if err != nil {
			return nil, fmt.Errorf("error parsing csv: %w", err)
		}

		out.Bands[i] = make([]model.BandEntry, len(c))

		for j, entry := range c {
			k, err := strconv.ParseFloat(strings.TrimSpace(entry[0]), 32)
			if err != nil {
				return nil, fmt.Errorf("could not convert to float: %w", err)
			}
			e, err := strconv.ParseFloat(strings.TrimSpace(entry[1]), 32)
			if err != nil {
				return nil, fmt.Errorf("could not convert to float: %w", err)
			}

			out.Bands[i][j] = model.BandEntry{
				Energy: float32(e),
				K:      float32(k),
			}
		}
	}

	symmEntry := make([][]model.BandEntry, 3)

	for i, v := range symms {
		c, err := csv.NewReader(strings.NewReader(v)).ReadAll()
		if err != nil {
			return nil, fmt.Errorf("error parsing csv: %w", err)
		}
		symmEntry[i] = make([]model.BandEntry, len(c))
		for j, e := range c {
			k, err := strconv.ParseFloat(strings.TrimSpace(e[0]), 32)
			if err != nil {
				return nil, fmt.Errorf("could not convert to float: %w", err)
			}
			e, err := strconv.ParseFloat(strings.TrimSpace(e[1]), 32)
			if err != nil {
				return nil, fmt.Errorf("could not convert to float: %w", err)
			}

			symmEntry[i][j] = model.BandEntry{
				Energy: float32(e),
				K:      float32(k),
			}
		}

	}

	out.Symmetry.S0_5 = symmEntry[0]
	out.Symmetry.S0 = symmEntry[1]
	out.Symmetry.S1 = symmEntry[2]

	c, err := csv.NewReader(strings.NewReader(fermi)).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error parsing csv: %w", err)
	}

	out.FermiLevel = make([]model.BandEntry, len(c))

	for i, v := range c {
		k, err := strconv.ParseFloat(strings.TrimSpace(v[0]), 32)
		if err != nil {
			return nil, fmt.Errorf("could not convert to float: %w", err)
		}
		e, err := strconv.ParseFloat(strings.TrimSpace(v[1]), 32)
		if err != nil {
			return nil, fmt.Errorf("could not convert to float: %w", err)
		}

		out.FermiLevel[i] = model.BandEntry{
			K:      float32(k),
			Energy: float32(e),
		}

	}

	return out, nil
}
